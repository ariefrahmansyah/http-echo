package main

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()
	app.Get("/echo", echo)
	app.Post("/echo", echo)
	app.Listen(8080)
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
