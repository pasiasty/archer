package actions

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/pasiasty/archer/server"
)

// GameGetWorld default implementation.
func GameGetWorld(c buffalo.Context) error {
	players := []string{"", ""}
	w := server.CreateWorld(players)
	return c.Render(http.StatusOK, r.JSON(w.GetPublicWorld()))
}
