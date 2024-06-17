package middlewares

// func LoggingMiddleware(logger *zap.Logger) fiber.Handler {
// 	return func(c *fiber.Ctx) error {
// 		start := time.Now()

// 		// Proses permintaan
// 		err := c.Next()

// 		// Ambil data permintaan dan respons
// 		request := string(c.Request().Body())
// 		response := string(c.Response().Body())
// 		duration := time.Since(start)

// 		// Catat log
// 		logger.Info("Request Log",
// 			zap.String("method", c.Method()),
// 			zap.String("url", c.OriginalURL()),
// 			zap.String("request", request),
// 			zap.String("response", response),
// 			zap.Error(err),
// 			zap.Duration("duration", duration),
// 		)

// 		return err
// 	}
// }
