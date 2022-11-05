package http_router

import (
	"cloud-configurations/internal/database"
	"cloud-configurations/internal/utils"
	"github.com/gofiber/fiber/v2"
)

func Create(c *fiber.Ctx) error {
	if len(c.FormValue("username")) == 0 || len(c.FormValue("name")) == 0 || len(c.FormValue("data")) == 0 {
		return c.JSON(fiber.Map{"Status": "Failure"})
	}

	err := database.DB.Create(&database.Configuration{
		UID:   utils.RandStringRunes(24),
		Name:  c.FormValue("name"),
		Owner: c.FormValue("username"),
		Data:  c.FormValue("data"),
	})

	if err != nil {
		return c.JSON(fiber.Map{"Status": "Failure"})
	}

	return c.JSON(fiber.Map{"Status": "Success"})
}

func Get(c *fiber.Ctx) error {
	if len(c.FormValue("uid")) == 0 || len(c.FormValue("username")) == 0 {
		return c.JSON(fiber.Map{"Status": "Failure"})
	}

	var userConfig database.Configuration
	err := database.DB.Find(&userConfig, "`uid` = ?", c.FormValue("uid"))
	if err != nil {
		return c.JSON(fiber.Map{"Status": "Failure"})
	}

	if userConfig.Owner != c.FormValue("username") {
		return c.JSON(fiber.Map{"Status": "Failure"})
	}

	return c.JSON(fiber.Map{"Status": "Success", "Configuration": userConfig})
}

func GetAll(c *fiber.Ctx) error {
	if len(c.FormValue("username")) == 0 {
		return c.JSON(fiber.Map{"Status": "Failure"})
	}

	var userConfigs []database.Configuration
	database.DB.Find(&userConfigs, "`owner` = ?", c.FormValue("username"))

	if len(userConfigs) == 0 {
		return c.JSON(fiber.Map{"Status": "Empty"})
	}

	for a := 0; a < len(userConfigs); a++ {
		userConfigs[a].Data = "";
	}

	return c.JSON(fiber.Map{"Status": "Success", "Configs": userConfigs})
}

func Save(c *fiber.Ctx) error {
	if len(c.FormValue("uid")) == 0 || len(c.FormValue("username")) == 0 || len(c.FormValue("data")) == 0 {
		return c.JSON(fiber.Map{"Status": "Failure"})
	}

	var userConfig database.Configuration
	database.DB.Find(&userConfig, "`uid` = ?", c.FormValue("uid"))

	if userConfig.Owner != c.FormValue("username") {
		return c.JSON(fiber.Map{"Status": "Failure"})
	}

	err := database.DB.Model(&userConfig).Updates(map[string]interface{}{
		"`data`": c.FormValue("data"),
	})
	if err != nil {
		return c.JSON(fiber.Map{"Status": "Failure"})
	}

	return c.JSON(fiber.Map{"Status": "Success"})
}

func Delete(c *fiber.Ctx) error {
	if len(c.FormValue("uid")) == 0 || len(c.FormValue("username")) == 0 || len(c.FormValue("data")) == 0 {
		return c.JSON(fiber.Map{"Status": "Failure"})
	}

	var userConfig database.Configuration
	database.DB.Find(&userConfig, "`uid` = ?", c.FormValue("uid"))

	if userConfig.Owner != c.FormValue("username") {
		return c.JSON(fiber.Map{"Status": "Failure"})
	}

	err := database.DB.Delete(&userConfig)
	if err != nil {
		return c.JSON(fiber.Map{"Status": "Failure"})
	}

	return c.JSON(fiber.Map{"Status": "Success"})
}
