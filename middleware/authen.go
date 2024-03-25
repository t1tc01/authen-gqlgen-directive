package middleware

import (
	"context"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		header := c.Get("Authorization")

		if header == "" {
			return c.Next()
		}

		//
		tokenStr := strings.TrimPrefix(header, "Bearer")
		c.Context().SetUserValue("have_token", true)
		c.Context().SetUserValue("token", tokenStr)

		fmt.Println("PASS: Authen middleware")
		return c.Next()

	}
}

func HaveToken(ctx context.Context) bool {
	token, _ := ctx.Value("have_token").(bool)
	return token
}

func Token(ctx context.Context) string {
	token, _ := ctx.Value("token").(string) // Safe to ignore 'ok' due to preceding check
	return token
}
