package actions

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/pasiasty/archer/server"
)

// PreparationCreateGame default implementation.
func PreparationCreateGame(c buffalo.Context) error {
	user := gm.CreateGame()
	user.StoreToCookie(c)
	return c.Render(http.StatusOK, r.JSON(server.EmptyResponse{}))
}

// PreparationJoinGame default implementation.
func PreparationJoinGame(c buffalo.Context) error {
	gameID := c.Param("game_id")
	user, err := gm.JoinGame(gameID)
	if err != nil {
		return err
	}
	user.StoreToCookie(c)
	return c.Render(http.StatusOK, r.JSON(server.EmptyResponse{}))
}

// PreparationUserReady default implementation.
func PreparationUserReady(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.JSON(server.EmptyResponse{}))
}

// PreparationPollGame default implementation.
func PreparationPollGame(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.JSON(server.EmptyResponse{}))
}
