package server

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gobuffalo/buffalo"
)

// IgnoreBots is a middleware used for ignoring requests coming from bots.
func IgnoreBots(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		userAgent := strings.ToUpper(c.Request().UserAgent())
		if !strings.Contains(userAgent, "MOZILLA") || strings.Contains(userAgent, "BOT") {
			return c.Error(http.StatusForbidden, errors.New("bots not permitted"))
		}
		err := next(c)
		return err
	}
}
