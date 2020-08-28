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
	return c.Render(http.StatusOK, r.JSON(server.Empty{}))
}

// PreparationJoinGame default implementation.
func PreparationJoinGame(c buffalo.Context) error {
	gameID := c.Param("game_id")
	user, err := gm.JoinGame(c, gameID)
	if err != nil {
		return err
	}
	user.StoreToCookie(c)
	return c.Render(http.StatusOK, r.JSON(server.Empty{}))
}

// PreparationUserReady default implementation.
func PreparationUserReady(c buffalo.Context) error {
	gameID := c.Param("game_id")
	userID := c.Param("user_id")

	if err := gm.MarkUserReady(c, gameID, userID); err != nil {
		return err
	}
	return getUsersList(c, gameID)
}

// PreparationListUsers default implementation.
func PreparationListUsers(c buffalo.Context) error {
	gameID := c.Param("game_id")
	return getUsersList(c, gameID)
}

// PreparationAddPlayer default implementation.
func PreparationAddPlayer(c buffalo.Context) error {
	gameID := c.Param("game_id")
	userID := c.Param("user_id")

	if err := gm.AddPlayer(c, gameID, userID); err != nil {
		return err
	}

	return getUsersList(c, gameID)
}

// PreparationRemovePlayer default implementation.
func PreparationRemovePlayer(c buffalo.Context) error {
	gameID := c.Param("game_id")
	userID := c.Param("user_id")

	if err := gm.RemovePlayer(c, gameID, userID); err != nil {
		return err
	}

	return getUsersList(c, gameID)
}

func getUsersList(c buffalo.Context, gameID string) error {
	game, err := gm.GetGame(c, gameID)
	if err != nil {
		return err
	}

	return c.Render(http.StatusOK, r.JSON(server.UsersList{Users: game.GetUsersList()}))
}

// PreparationStartGame default implementation.
func PreparationStartGame(c buffalo.Context) error {
	gameID := c.Param("game_id")
	userID := c.Param("user_id")

	if err := gm.StartGame(c, gameID, userID); err != nil {
		return err
	}
	return c.Render(http.StatusOK, r.JSON(server.Empty{}))
}

// PreparationGameHasStarted default implementation.
func PreparationGameHasStarted(c buffalo.Context) error {
	gameID := c.Param("game_id")

	started, err := gm.GameHasStarted(c, gameID)
	if err != nil {
		return err
	}
	return c.Render(http.StatusOK, r.JSON(started))
}
