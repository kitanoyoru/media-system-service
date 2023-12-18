/*package v0

import "github.com/gofiber/fiber/v2"

const (
	AuthTokenCookie = "auth_token"
)

func Authorize(c *fiber.Ctx) error {
	cookie := c.Cookies(AuthTokenCookie)

	if err := r.authService.VerifyJWTToken(cookie); err != nil {
		return fiber.ErrUnauthorized
	}

	return c.Next()
}
*/
