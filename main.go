package main

import (
	"net/http"
	"os"
	"time"

	"github.com/gofiber/fiber"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	app := fiber.New()

	app.Get("/*", echo)
	app.Post("/*", echo)

	app.Listen(port)
}

func echo(c *fiber.Ctx) {
	delay := c.Query("delay")
	if delay == "" {
		delay = "0ms"
	}

	delayDuration, err := time.ParseDuration(delay)
	if err != nil {
		c.Status(http.StatusBadRequest).Send(err)
		return
	}

	time.Sleep(delayDuration)

	body := c.Query("message")
	if c.Method() == "POST" {
		body = c.Body()
	}

	c.Set("Content-Type", c.Get("Content-Type"))
	c.Send(body)
}
