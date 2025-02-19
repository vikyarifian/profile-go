package main

import (
	"log"
	"net"
	"net/http"
	"profile-go/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// func main() {

// 	app := fiber.New(fiber.Config{
// 		Network: fiber.NetworkTCP,
// 	})

// 	app.Use(logger.New())

// 	app.Static("/assets", "./assets")

// 	app.Use(func(c *fiber.Ctx) error {
// 		c.Set("Cache-Control", "no-store, no-cache, must-revalidate, proxy-revalidate")
// 		c.Set("Pragma", "no-cache")
// 		c.Set("Expires", "0")
// 		c.Set("Surrogate-Control", "no-store")
// 		return c.Next()
// 	})

// 	routes.WebRoutes(app)
// 	listen, _ := net.Listen("tcp", ":6861")
// 	// BulkUser()
// 	// BulProduct()
// 	log.Fatal(app.Listener(listen))

// }

func handler() http.HandlerFunc {

	app := fiber.New(fiber.Config{
		Network: fiber.NetworkTCP,
	})

	app.Use(logger.New())

	app.Static("/assets", "./assets")

	app.Use(func(c *fiber.Ctx) error {
		c.Set("Cache-Control", "no-store, no-cache, must-revalidate, proxy-revalidate")
		c.Set("Pragma", "no-cache")
		c.Set("Expires", "0")
		c.Set("Surrogate-Control", "no-store")
		return c.Next()
	})

	routes.WebRoutes(app)
	listen, _ := net.Listen("tcp", ":6861")
	// BulkUser()
	// BulProduct()
	log.Fatal(app.Listener(listen))

	return adaptor.FiberApp(app)
}

// Handler is the main entry point of the application. Think of it like the main() method
func Handler(w http.ResponseWriter, r *http.Request) {
	// This is needed to set the proper request path in `*fiber.Ctx`
	r.RequestURI = r.URL.String()

	handler().ServeHTTP(w, r)
}
