package helper

import (
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

func JSON(c *fiber.Ctx, response interface{}) error {
	raw, err := json.MarshalWithOption(response, json.DisableHTMLEscape())
	if err != nil {
		return err
	}
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)
	err = c.Send(raw)
	if err != nil {
		return err
	}
	return nil
}
