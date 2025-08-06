package httpServer

import "github.com/gofiber/fiber/v2"

type StdResponse struct {
	Code    int         `json:"-"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func (res StdResponse) Send(ctx *fiber.Ctx) error {
	return ctx.Status(res.Code).JSON(res)
}

func NewSuccessRes(data interface{}) *StdResponse {
	return &StdResponse{
		Code:    200,
		Message: "OK",
		Data:    data,
	}
}

func NewCreatedRes(data interface{}) *StdResponse {
	return &StdResponse{
		Code:    201,
		Message: "Created",
		Data:    data,
	}
}

func NewErrorRes(err error) *StdResponse {
	return &StdResponse{
		Code:    500,
		Message: err.Error(),
	}
}

func NewErrorCodeRes(code int, err error) *StdResponse {
	return &StdResponse{
		Code:    code,
		Message: err.Error(),
	}
}
