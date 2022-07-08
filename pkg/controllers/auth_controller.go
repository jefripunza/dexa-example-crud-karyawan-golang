package controllers

import (
	utils "dexa-training-crud-karyawan/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

// untuk memperbarui token jika token yang lama masih aktif
func NewToken(c *fiber.Ctx) error {
	// Generate a new Access token.
	token, err := utils.GenerateNewAccessToken()
	if err != nil {
		// Return status 500 and token generation error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"access_token": token,
	})
}
