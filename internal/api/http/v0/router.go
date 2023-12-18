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
	"github.com/kitanoyoru/media-system-service/internal/services/recommendation"
	"github.com/kitanoyoru/media-system-service/internal/services/tendencies"
	"gorm.io/gorm"
)

const (
	AuthTokenCookie = "auth_token"
)

func NewRouter(db *gorm.DB) (*fiber.App, error) {
	app := fiber.New()

	app.Use(recover.New())
	app.Use(session.New())
	app.Use(logger.New())
	app.Use(cors.New())
	app.Use(limiter.New(limiter.Config{
		Max:        20,
		Expiration: time.Second * 60,
	}))

	authService := auth.NewAuthService(db)
	controllers.NewAuthController(db, authService).Route(app)

	app.Use("/api/v0", func(c *fiber.Ctx) error {
		cookie := c.Cookies(AuthTokenCookie)

		if err := authService.VerifyJWTToken(cookie); err != nil {
			return fiber.ErrUnauthorized
		}

		return c.Next()
	})

	tendencyService := tendencies.NewTendencyService(db)
	controllers.NewTendencyController(db, tendencyService).Route(app)

	recommendationService := recommendation.NewRecommendationService(db)
	controllers.NewRecommendationController(db, recommendationService).Route(app)

	return app, nil
}
