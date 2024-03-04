package dashboard

import (
	"github.com/Nedinator/ribbit/data"
	"github.com/gofiber/fiber/v2"
)

func GetLinks(c *fiber.Ctx) ([]data.Url, error) {
	var links []data.Url

	owner := c.Locals("username").(string)
	if owner == "" {
		return nil, fiber.NewError(fiber.StatusBadRequest, "User not recognized")
	}

	err := data.DB().Where("owner = ?", owner).Order("created_at DESC").Find(&links).Error
	if err != nil {
		return nil, err
	}

	return links, nil
}
