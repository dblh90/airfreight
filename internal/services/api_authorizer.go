package services

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
)

func Unauthorizer(c *fiber.Ctx) error {
	response := c.Response()
	marshal, err := json.Marshal(map[string]string{
		"error": "Unauthorized",
	})
	if err != nil {
		return err
	}
	response.Header.SetContentType(fiber.MIMEApplicationJSONCharsetUTF8)
	response.SetBodyRaw(marshal)
	response.SetStatusCode(fiber.StatusUnauthorized)
	return nil
}

func Authorizer(user, pass string) bool {
	if user == "hamza" && pass == "123456" {
		return true
	}
	if user == "admin" && pass == "aabbcc" {
		return true
	}
	return false
}
