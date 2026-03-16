package middleware

import (
	"go-ecommerce-api/pkg/logger"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func RequestLogger() fiber.Handler {

	return func(c *fiber.Ctx) error {

		start := time.Now()
		err := c.Next()
		duration := time.Since(start)

		logger.Log.Info("HTTP Request",
			zap.String("method", c.Method()),
			zap.String("path", c.Path()),
			zap.Int("status", c.Response().StatusCode()),
			zap.Duration("latency", duration),
		)

		return err
	}
}
