package actions

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
)

// PreparationCreateGame default implementation.
func PreparationCreateGame(c buffalo.Context) error {
	gameID, userID := gm.createGame()
	return c.Render(http.StatusOK, r.JSON(GameResponse{GameID: gameID, UserID: userID}))
}

// PreparationJoinGame default implementation.
func PreparationJoinGame(c buffalo.Context) error {
	gameID := c.Param("game_id")
	userID, err := gm.joinGame(gameID)
	if err != nil {
		return err
	}
	return c.Render(http.StatusOK, r.JSON(GameResponse{GameID: gameID, UserID: userID}))
}

// PreparationUserReady default implementation.
func PreparationUserReady(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.JSON(GameResponse{}))
}

// PreparationPollGame default implementation.
func PreparationPollGame(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.JSON(GameResponse{}))
}
