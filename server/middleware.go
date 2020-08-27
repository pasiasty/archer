package server

import (
	"errors"
	"strings"

	"github.com/gobuffalo/buffalo"
)

// IgnoreBots is a middleware used for ignoring requests coming from bots.
func IgnoreBots(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		if strings.Contains(c.Request().UserAgent(), "Slackbot") {
			return errors.New("bots not permitted")
		}
		err := next(c)
		return err
	}
}
