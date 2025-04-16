package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tegaraditya/mmj-whatsapp-client/internal/api/handlers"
	"github.com/tegaraditya/mmj-whatsapp-client/internal/api/middlewares"
	"github.com/tegaraditya/mmj-whatsapp-client/pkg/whatsapp"
)

func SetupRoutes(app *fiber.App, wac *whatsapp.WhatsAppClient) {
	h := handlers.CreateHandler(wac)

	app.Get("/", h.GetAppInfo)
	app.Post("/send", middlewares.AuthorizeIP(), h.SendMessage)
}
