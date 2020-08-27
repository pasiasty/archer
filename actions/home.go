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
	userID, err := gm.joinGame(gameID)
	if err != nil {
		return err
	}
	setCookie(c, "game_id", gameID)
	setCookie(c, "user_id", userID)
	return c.Redirect(307, "/")
}
