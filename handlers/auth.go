package handlers

import (
	"os"
	"time"

	"github.com/Nedinator/ribbit/data"
	"github.com/Nedinator/ribbit/util"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Signup(c *fiber.Ctx) error {
	user := new(data.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse body"})
	}
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

	return c.Redirect("/login")
}

func Login(c *fiber.Ctx) error {
	user := new(data.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Cannot parse body"})
	}

	dbUser := new(data.User)
	filter := bson.M{"username": user.Username}

	err := data.Db.Collection("users").FindOne(c.Context(), filter).Decode(&dbUser)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid credentials"})
		}
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Internal error"})
	}

	if !util.CheckPasswordHash(user.Password, dbUser.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid credentials", "success": false})
	}

	token, err := GenerateJWT(dbUser.ID, dbUser.Username)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot generate token", "success": false})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "rbbt_token",
		Value:    token,
		Expires:  time.Now().Add(24 * time.Hour * 30),
		HTTPOnly: true,
	})

	return c.JSON(fiber.Map{"message": "Login successful", "success": true})

}

func Logout(c *fiber.Ctx) error {
	expired := time.Now().Add(-time.Hour * 24)
	c.Cookie(&fiber.Cookie{
		Name:    "rbbt_token",
		Value:   "",
		Expires: expired,
	})
	return c.Redirect("/")
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

func GenerateJWT(id string, username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 24 * 30).Unix()
	jwtSecret := os.Getenv("JWT_SECRET")
	tokenString, err := token.SignedString([]byte(jwtSecret))
	return tokenString, err
}
