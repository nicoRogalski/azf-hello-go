package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rogalni/azf-hello-go/internal/adapter/rest/dto"
	"github.com/rogalni/azf-hello-go/internal/service"
)

func HelloHandler(c *fiber.Ctx) error {

	service := service.NewHelloService()
	m := service.GetHelloMessage()
	c.JSON(&dto.Message{
		Id:   m.Id,
		Code: m.Code,
		Text: m.Text,
	})
	return nil
}
