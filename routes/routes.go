package routes

import (
	"profile-go/components"
	"profile-go/utils"

	"github.com/gofiber/fiber/v2"
)

func routes() {

}

func WebRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return utils.Render(c, components.Layout())
	})
}
