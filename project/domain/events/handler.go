package events

import "github.com/gofiber/fiber/v2"

var BaseRoute = "/events"

type httpHandler struct {
	service Servicer
}

func NewHttpHandler(app *fiber.App, service Servicer) {
	handler := &httpHandler{
		service: service,
	}

	app.Get(BaseRoute, handler.getEvents)
}

func (h *httpHandler) getEvents(c *fiber.Ctx) error {
	events := h.service.GetAllEvents()
	return c.JSON(events)
}
