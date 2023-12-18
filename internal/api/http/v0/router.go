package v0

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/kitanoyoru/media-system-service/internal/api/http/v0/controllers"
	"github.com/kitanoyoru/media-system-service/internal/services/auth"
	"gorm.io/gorm"
)

func NewRouter(db *gorm.DB, authService *auth.AuthService) (*fiber.App, error) {
	app := fiber.New()

	app.Use(recover.New())
	app.Use(session.New())
	app.Use(logger.New())
	app.Use(cors.New())
	app.Use(limiter.New(limiter.Config{
		Max:        20,
		Expiration: time.Second * 60,
	}))

	controllers.NewAuthController(db, authService).Route(app)

	return app, nil
}
