package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/tegaraditya/mmj-whatsapp-client/internal/api/routes"
	"github.com/tegaraditya/mmj-whatsapp-client/pkg/whatsapp"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Failed to load .env file: %v. ", err)
		fmt.Println("Using default environment variables")
	}

	app := fiber.New(fiber.Config{
		JSONEncoder:  sonic.Marshal,
		JSONDecoder:  sonic.Unmarshal,
		ServerHeader: "Fiber",
	})

	app.Use(cors.New())
	app.Use(helmet.New())
	app.Use(healthcheck.New())
	app.Use(logger.New())

	client, err := whatsapp.NewClient()
	if err != nil {
		panic(fmt.Sprintf("Failed to create WhatsApp client: %v", err))
	}

	err = client.Start()
	if err != nil {
		panic(fmt.Sprintf("Failed to start WhatsApp client: %v", err))
	}

	routes.SetupRoutes(app, client)

	log.Fatal(app.Listen(fmt.Sprintf(":%s", "3000")))

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	client.Stop()
}
