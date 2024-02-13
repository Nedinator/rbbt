package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func JwtMiddleware(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(http.StatusUnauthorized).SendString("Authorization header is required")
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenString == authHeader {
		return c.Status(http.StatusUnauthorized).SendString("Invalid or malformed token")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return c.Status(http.StatusUnauthorized).SendString("Invalid or expired token")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		c.Locals("userID", claims["userID"])
		c.Locals("username", claims["username"])
		return c.Next()
	} else {
		return c.Status(http.StatusUnauthorized).SendString("Invalid token")
	}
}

func AuthStatusMiddleware(c *fiber.Ctx) error {
	isLoggedIn := false
	tokenString := c.Cookies("rbbt_token")
	fmt.Println("Token from cookie:", tokenString) // Add logging

	if tokenString != "" {
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		fmt.Println("Token validation error:", err) // Add logging
		if err == nil && token.Valid {
			fmt.Println("Token is valid") // Add logging
			isLoggedIn = true
		}
	}

	c.Locals("IsLoggedIn", isLoggedIn)
	return c.Next()
}
