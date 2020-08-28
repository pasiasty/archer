package actions

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
)

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("index.html"))
}

// JoinHomeHandler lets to join the game.
func JoinHomeHandler(c buffalo.Context) error {
	gameID := c.Param("game_id")
	user, err := gm.JoinGame(gameID)
	if err != nil {
		return c.Error(http.StatusNotFound, err)
	}
	user.StoreToCookie(c)
	return c.Redirect(http.StatusTemporaryRedirect, "/")
}
