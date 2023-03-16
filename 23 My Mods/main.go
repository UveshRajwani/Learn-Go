package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// fmt.Println("Hello Mod in golang")
	// greeter()
	app := fiber.New(
		fiber.Config{
			Prefork:     true,
			JSONEncoder: json.Marshal,
			JSONDecoder: json.Unmarshal,
		},
	)

	app.Get("/", func(c *fiber.Ctx) error {
		var u = map[string]string{"timing": time.Now().String()}
		// return c.SendString(string(u))
		fmt.Println("api used")
		return c.JSON(u)
	})

	app.Listen(":3000")
}

func greeter() {
	fmt.Println("Hey there mod users")
}
