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
		c.Locals("userID", claims["id"])
		c.Locals("username", claims["username"])
		return c.Next()
	} else {
		return c.Status(http.StatusUnauthorized).SendString("Invalid token")
	}
}

func AuthStatusMiddleware(c *fiber.Ctx) error {
	isLoggedIn := false
	var username string
	var userID string
	tokenString := c.Cookies("rbbt_token")

	if tokenString != "" {
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err == nil && token.Valid {

			isLoggedIn = true

			if claims, ok := token.Claims.(jwt.MapClaims); ok {

				if usernameClaim, ok := claims["username"].(string); ok {
					username = usernameClaim
				}
				if idClaim, ok := claims["id"].(string); ok {
					userID = idClaim
				}
			}
		} else {
			fmt.Println("Token validation error:", err)
		}
	}

	c.Locals("IsLoggedIn", isLoggedIn)
	c.Locals("Username", username)
	c.Locals("ID", userID)

	return c.Next()
}
