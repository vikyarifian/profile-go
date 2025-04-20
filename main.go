package main

import (
	"log"
	"net"
	"os"
	"path/filepath"
	"profile-go/routes"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func LoadEnv() {
	appPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	appPath = strings.ReplaceAll(appPath, "\\bin", "")
	appPath = strings.ReplaceAll(appPath, "\\main", "")
	appPath = strings.ReplaceAll(appPath, "/bin", "")
	appPath = strings.ReplaceAll(appPath, "/main", "")
	appPath = strings.ReplaceAll(appPath, "\\", "/")

	os.Setenv("APP_PATH", appPath+"/")

	err = godotenv.Load(appPath + ".env")
	if err != nil {
		err = godotenv.Load(appPath + "/.env")
		if err != nil {
			err = godotenv.Load(appPath + "\\.env")
			if err != nil {
				log.Fatal("Error loading .env file path: ", appPath)
			}
		}
	}
}

func main() {

	LoadEnv()

	app := fiber.New(fiber.Config{
		Network: fiber.NetworkTCP,
	})

	app.Use(logger.New())

	app.Static("assets", os.Getenv("APP_PATH")+"assets")

	app.Use(func(c *fiber.Ctx) error {
		c.Set("Cache-Control", "no-store, no-cache, must-revalidate, proxy-revalidate")
		c.Set("Pragma", "no-cache")
		c.Set("Expires", "0")
		c.Set("Surrogate-Control", "no-store")
		return c.Next()
	})

	routes.WebRoutes(app)
	listen, _ := net.Listen("tcp", ":6861")

	log.Fatal(app.Listener(listen))

}
