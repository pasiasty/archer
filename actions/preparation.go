package actions

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
)

// PreparationCreateGame default implementation.
func PreparationCreateGame(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.JSON("OK"))
}
