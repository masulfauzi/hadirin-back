package utils

import "github.com/gofiber/fiber/v2"

// Success mengirim response sukses dengan format standar.
// code diisi HTTP status code (fiber.StatusOK, fiber.StatusCreated, dst.)
func Success(c *fiber.Ctx, code int, message string, data interface{}) error {
	return c.Status(code).JSON(fiber.Map{
		"success": true,
		"code":    code,
		"message": message,
		"data":    data,
	})
}

// Error mengirim response gagal dengan format standar, data selalu null
func Error(c *fiber.Ctx, code int, message string) error {
	return c.Status(code).JSON(fiber.Map{
		"success": false,
		"code":    code,
		"message": message,
		"data":    nil,
	})
}
