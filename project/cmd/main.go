package main

import (
	"first-app/config"
	"first-app/db"
	"first-app/domain/events"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	db.InitDB()

	app := fiber.New()
	app.Use(cors.New())

	eventsRepository := events.NewRepository(db.DB.Db)
	eventsService := events.NewService(events.ServiceParams{
		Repo: eventsRepository,
	})

	events.NewHttpHandler(app, eventsService)

	app.Use(func(c *fiber.Ctx) error {
		return fiber.ErrNotFound
	})

	log.Fatal(app.Listen(config.GetEnv("APP_PORT")))

}
