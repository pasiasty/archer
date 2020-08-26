package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/pasiasty/archer/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
