package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"

	"hadirin-back/utils"
)

func Protected(secret string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			return unauthorized(c, "Token tidak ditemukan")
		}
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			// Tolak token yang algoritmanya bukan HMAC (mencegah
			// serangan pemalsuan dengan algoritma "none" / RS256)
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(secret), nil
		})
		if err != nil || !token.Valid {
			return unauthorized(c, "Token tidak valid atau kedaluwarsa")
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return unauthorized(c, "Token tidak valid")
		}
		userID, ok := claims["user_id"].(float64)
		if !ok {
			return unauthorized(c, "Token tidak valid")
		}

		// Simpan ID user agar bisa diambil handler via c.Locals("user_id")
		c.Locals("user_id", uint(userID))

		return c.Next()
	}
}

func unauthorized(c *fiber.Ctx, message string) error {
	return utils.Error(c, fiber.StatusUnauthorized, message)
}
