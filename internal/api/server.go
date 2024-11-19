package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func StartServer() {

	fmt.Println("I am main function...")

	app := fiber.New()

	// routes

	app.Listen("localhost:9000")
}
