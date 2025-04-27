package main

import (
	"log"
	"math/rand/v2"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func main() {
	quotes := []string{
		"No debi haber hecho eso",
		"Necesitaba aprender esta leccion para avanzar",
	}
	app := fiber.New()
	api := app.Group("/api/v1")

	api.Get("/quotes", func(ctx *fiber.Ctx) error {
		quoteIdx := rand.IntN(len(quotes))

		return ctx.JSON(fiber.Map{
			"error": false,
			"data": fiber.Map{
				"quote": quotes[quoteIdx],
			},
		})
	})

	api.Post("/quotes", func(ctx *fiber.Ctx) error {
		type Payload struct {
			Quote string `json:"quote"`
		}

		var payload Payload

		if err := ctx.BodyParser(&payload); err != nil {
			return ctx.SendStatus(http.StatusInternalServerError)
		}

		quotes = append(quotes, payload.Quote)

		return ctx.SendStatus(http.StatusCreated)
	})

	if err := app.Listen(":9999"); err != nil {
		log.Fatal(err)
	}
}
