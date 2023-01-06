package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type User struct {
	ID      int    `json:"user_id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Mobile  string `json:"mobile_number"`
	Address string `json:"address"`
	Age     string `json:"age"`
}

func main() {
	fmt.Print("Hello World")

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	users := []User{}

	app.Get("/api/allUsers", func(c *fiber.Ctx) error {
		return c.JSON(users)
	})

	app.Post("/api/users", func(c *fiber.Ctx) error {
		user := &User{}

		if err := c.BodyParser(user); err != nil {
			return err
		}

		user.ID = len(users) + 1

		users = append(users, *user)

		return c.JSON(users)
	})

	app.Patch("api/users/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")

		if err != nil {
			return c.Status(401).SendString("Invalid ID")
		}

		user := &User{}

		if err := c.BodyParser(user); err != nil {
			return c.Status(401).SendString("Invalid request body")
		}

		for i, u := range users {
			if u.ID == id {
				users[i].Name = user.Name
				users[i].Email = user.Email
				users[i].Mobile = user.Mobile
				users[i].Address = user.Address
				users[i].Age = user.Age
				break
			}
		}

		return c.JSON(users)

	})

	app.Delete("api/users/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")

		if err != nil {
			return c.Status(401).SendString("Invalid ID")
		}

		for i, u := range users {
			if u.ID == id {
				users = append(users[:i], users[i+1:]...)
				break
			}
		}

		return c.JSON(users)

	})

	log.Fatal(app.Listen(":4000"))
}
