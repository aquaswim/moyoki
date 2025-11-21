package mockServer

import (
	"time"

	"github.com/aquaswim/moyoki/internal/core/domain"
	"github.com/gofiber/fiber/v2"
)

func createAccessLog(ctx *fiber.Ctx) *domain.AccessLog {
	return &domain.AccessLog{
		Path:       ctx.Path(),
		Method:     ctx.Method(),
		RemoteAddr: ctx.IP(),
		ReqQuery:   string(ctx.Request().URI().QueryString()),
		ReqBody:    string(ctx.BodyRaw()),
		ReqHeaders: ctx.Request().Header.String(),
		CreatedAt:  time.Now(),
	}
}
