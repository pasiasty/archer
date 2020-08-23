package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/pasiasty/archer2/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
