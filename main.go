package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Muhammadsoomro88/go-jira/setup"
	"github.com/Muhammadsoomro88/go-jira/store"
	"github.com/Muhammadsoomro88/go-jira/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func getStoryData(c *fiber.Ctx) error {
	story := c.Params("storyNum")

	vi := setup.GetViper()
	resp, err := http.Get("https://" + vi.GetString("Username") + ":" + vi.GetString("Password") + vi.GetString("Url") + story)
	if err != nil {
		c.Status(fiber.StatusBadRequest).JSON(err.Error())
		return err
	}

	// decode data into struct
	var d utils.Jira
	if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(err.Error())
		return err
	}

	store.StoreJiraData(story, d)

	return c.Status(fiber.StatusOK).JSON(d)
}

func main() {
	fmt.Println("JIRA Rest Api")

	// fiber initialization
	app := fiber.New()
	app.Use(logger.New())
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to JIRA Rest API")
	})

	app.Get("/story/:storyNum", getStoryData)

	log.Fatal(app.Listen(":8080"))
}
