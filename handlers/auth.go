package handlers

import (
	"fmt"

	"github.com/Nedinator/ribbit/data"
	"github.com/Nedinator/ribbit/util"
	"github.com/gofiber/fiber/v2"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Signup(c *fiber.Ctx) error {
	user := new(data.User)

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse body"})
	}
	fmt.Println(user)
	if checkUsername(user.Username, c) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Username already exists"})
	}

	hashedPassword, err := util.HashPassword(user.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot hash password"})
	}
	user.Password = hashedPassword

	id, err := gonanoid.New(15)
	if err != nil {
		return c.Status(500).SendString("Internal Server Error. If you see this you should prolly dial 911...")
	}

	user.ID = id

	_, err = data.Db.Collection("users").InsertOne(c.Context(), user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot create user"})
	}

	return c.JSON("User created")
}

func Login(c *fiber.Ctx) error {
	user := new(data.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse body"})
	}

	dbUser := new(data.User)
	filter := bson.M{"username": user.Username}

	err := data.Db.Collection("users").FindOne(c.Context(), filter).Decode(&dbUser)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
		}
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Internal error"})
	}

	if !util.CheckPasswordHash(user.Password, dbUser.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	// TODO: Generate JWT

	return c.JSON("user logged in")
}

func checkUsername(username string, c *fiber.Ctx) bool {
	filter := bson.M{"username": username}
	user := new(data.User)
	if err := data.Db.Collection("users").FindOne(c.Context(), filter).Decode(&user); err != nil {
		if err == mongo.ErrNoDocuments {
			return false
		}
	}
	return true
}
