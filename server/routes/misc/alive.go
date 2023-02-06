package misc

import "github.com/gofiber/fiber/v2"

func Alive(ctx *fiber.Ctx) error {
	return ctx.SendString("iTorrent API v2 // Alive")
}
