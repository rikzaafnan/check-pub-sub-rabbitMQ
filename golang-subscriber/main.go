package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/streadway/amqp"
)

func main() {
	app := fiber.New()
	port := ":3002"

	conn, err := amqp.Dial("amqp://localhost:5672")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer ch.Close()

	exchange := "logs"
	q, err := ch.QueueDeclare("", false, false, true, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	err = ch.QueueBind(q.Name, "", exchange, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	app.Post("/subscribe", func(c *fiber.Ctx) error {
		msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
		if err != nil {
			log.Fatal(err)
		}

		go func() {
			for msg := range msgs {
				log.Printf("Received a message: %s", msg.Body)
			}
		}()

		return c.SendStatus(fiber.StatusOK)
	})

	log.Fatal(app.Listen(port))
}
