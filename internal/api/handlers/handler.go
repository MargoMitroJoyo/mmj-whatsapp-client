package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/tegaraditya/mmj-whatsapp-client/pkg/whatsapp"
)

type Handler struct {
	Client *whatsapp.WhatsAppClient
}

func CreateHandler(client *whatsapp.WhatsAppClient) *Handler {
	return &Handler{Client: client}
}

func (h *Handler) GetAppInfo(c *fiber.Ctx) error {
	return c.SendString("WhatsApp Client API")
}

func (h *Handler) SendMessage(c *fiber.Ctx) error {
	var req struct {
		To      string `json:"to"`
		Message string `json:"message"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request format",
		})
	}

	if req.To == "" || req.Message == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Missing required fields",
		})
	}

	err := h.Client.SendMessage(req.To, req.Message)
	if err != nil {
		panic(fmt.Sprintf("Failed to send message: %v", err))
	}

	return c.JSON(fiber.Map{
		"success": true,
		"to":      req.To,
		"message": req.Message,
	})
}
