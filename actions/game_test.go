package actions

import "net/http"

func (as *ActionSuite) Test_Game_GetWorld() {
	as.Request("/game/get_world", map[string]interface{}{"game_id": "abc"}, http.StatusNotFound)
	as.Request("/game/get_world", map[string]interface{}{"game_id": as.u.GameID}, http.StatusForbidden)

	as.g.Start(as.c, as.u.UserID)

	as.Request("/game/get_world", map[string]interface{}{"game_id": as.u.GameID}, http.StatusOK)
}
