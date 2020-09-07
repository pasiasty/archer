package actions

import (
	"fmt"
	"net/http"
	"strconv"

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

// PreparationGameStatus tells the status of the game.
func PreparationGameStatus(c buffalo.Context) error {
	gameID := c.Param("game_id")

	status, err := gm.GameStatus(c, gameID)
	if err != nil {
		return err
	}
	return c.Render(http.StatusOK, r.JSON(status))
}

// PreparationGameSettings defines settings of the game.
func PreparationGameSettings(c buffalo.Context) error {
	gameID := c.Param("game_id")
	userID := c.Param("user_id")
	shootTimeoutStr := c.Param("shoot_timeout")
	loopedWorldStr := c.Param("looped_world")

	shootTimeout, err := strconv.ParseInt(shootTimeoutStr, 10, 32)
	if err != nil {
		return c.Error(http.StatusBadRequest, fmt.Errorf("wrong timeout value: %v", err))
	}

	loopedWorld, err := strconv.ParseBool(loopedWorldStr)
	if err != nil {
		return c.Error(http.StatusBadRequest, fmt.Errorf("wrong looped world value: %v", err))
	}

	game, err := gm.GetGame(c, gameID)
	if err != nil {
		return err
	}

	gs := server.CreateGameSettings(
		server.WithShootTimeout(int32(shootTimeout)),
		server.WithLoopedWorld(loopedWorld))
	if err := game.ApplyGameSettings(c, userID, gs); err != nil {
		return err
	}
	return c.Render(http.StatusOK, r.JSON(server.Empty{}))
}
