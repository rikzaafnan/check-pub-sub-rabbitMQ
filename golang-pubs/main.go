package main

import (
	"log"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	// Buat koneksi ke server RabbitMQ
	conn, err := amqp.Dial("amqp://user:password@localhost:35672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// Buka channel
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// Deklarasikan queue
	q, err := ch.QueueDeclare(
		"hello", // Nama queue
		false,   // Durable
		false,   // Delete when unused
		false,   // Exclusive
		false,   // No-wait
		nil,     // Arguments
	)
	failOnError(err, "Failed to declare a queue")

	// Pesan yang akan dikirim ke queue
	body := "cobaa ciobacoabcaos"

	// Publish pesan ke queue
	err = ch.Publish(
		"",     // Exchange
		q.Name, // Routing key
		false,  // Mandatory
		false,  // Immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")

	log.Printf(" [x] Sent %s", body)
}
