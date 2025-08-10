package httpServer

import (
	"github.com/aquaswim/moyoki/internal/core/domain"
	"github.com/aquaswim/moyoki/internal/core/port"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
)

type PanelHandler struct {
	routeService port.RouteService
}

func NewPanelHandler(
	routeService port.RouteService,
) *PanelHandler {
	return &PanelHandler{
		routeService: routeService,
	}
}

func (h PanelHandler) RegisterHandler(app *fiber.App) {
	api := app.Group("/api")

	// manage routes
	api.Get("/routes", h.getRoutes)
	api.Post("/routes", h.createRoutes)
	api.Get("/routes/count", h.countRoutes)
	api.Get("/routes/:id", h.getRouteByID)
	api.Put("/routes/:id", h.updateRouteByID)
	api.Delete("/routes/:id", h.DeleteRouteByID)
}

func (h PanelHandler) getRoutes(ctx *fiber.Ctx) error {
	data, err := h.routeService.Find(ctx.Context())
	if err != nil {
		return NewErrorCodeRes(http.StatusBadRequest, err).Send(ctx)
	}

	return NewSuccessRes(data).Send(ctx)
}

func (h PanelHandler) createRoutes(ctx *fiber.Ctx) error {
	var data domain.RouteItemRequestData
	if err := ctx.BodyParser(&data); err != nil {
		return NewErrorCodeRes(http.StatusBadRequest, err).Send(ctx)
	}

	if err := h.routeService.Save(ctx.Context(), &data); err != nil {
		return NewErrorRes(err).Send(ctx)
	}

	return NewCreatedRes(data).Send(ctx)
}

func (h PanelHandler) getRouteByID(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return NewErrorCodeRes(http.StatusBadRequest, err).Send(ctx)
	}

	route, err := h.routeService.FindByID(ctx.Context(), id)
	if err != nil {
		return NewErrorCodeRes(http.StatusBadRequest, err).Send(ctx)
	}

	return NewSuccessRes(route).Send(ctx)
}

func (h PanelHandler) updateRouteByID(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return NewErrorCodeRes(http.StatusBadRequest, err).Send(ctx)
	}

	var data domain.RouteItemRequestData
	if err := ctx.BodyParser(&data); err != nil {
		return NewErrorCodeRes(http.StatusBadRequest, err).Send(ctx)
	}

	if err := h.routeService.Update(ctx.Context(), id, &data); err != nil {
		return NewErrorRes(err).Send(ctx)
	}

	return NewSuccessRes(data).Send(ctx)
}

func (h PanelHandler) DeleteRouteByID(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return NewErrorCodeRes(http.StatusBadRequest, err).Send(ctx)
	}

	if err := h.routeService.Delete(ctx.Context(), id); err != nil {
		return NewErrorRes(err).Send(ctx)
	}
	return NewSuccessRes(nil).Send(ctx)
}

func (h PanelHandler) countRoutes(ctx *fiber.Ctx) error {
	count, err := h.routeService.Count(ctx.Context())
	if err != nil {
		return NewErrorCodeRes(http.StatusBadRequest, err).Send(ctx)
	}
	return NewSuccessRes(count).Send(ctx)
}
