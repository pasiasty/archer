package actions

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pasiasty/archer/server"
)

var (
	base64Regexp = `[0-9a-zA-Z-_]{3,}=`
)

func (as *ActionSuite) Test_Preparation_CreateGame() {
	req := as.ProperRequest("/preparation/create_game")
	res := req.Post(new(interface{}))

	as.Equal(http.StatusOK, res.Code)
	setCookies := res.Header().Values("Set-Cookie")
	as.Equal(4, len(setCookies))

	gameIDRes := setCookies[0]
	userIDRes := setCookies[1]
	usernameRes := setCookies[2]
	isHostRes := setCookies[3]

	as.Regexp(fmt.Sprintf(`game_id=%s;`, base64Regexp), gameIDRes)
	as.Regexp(fmt.Sprintf(`user_id=%s;`, base64Regexp), userIDRes)
	as.Regexp(`username="[a-zA-Z ]+";`, usernameRes)
	as.Regexp(`is_host=true;`, isHostRes)
}

func (as *ActionSuite) Test_Preparation_JoinGame() {
	u := gm.CreateGame()
	gameID := u.GameID

	req := as.ProperRequest("/preparation/join_game")
	res := req.Post(map[string]interface{}{"game_id": gameID})
	as.Equal(http.StatusOK, res.Code)

	setCookies := res.Header().Values("Set-Cookie")
	as.Equal(4, len(setCookies))

	gameIDRes := setCookies[0]
	userIDRes := setCookies[1]
	usernameRes := setCookies[2]
	isHostRes := setCookies[3]

	as.Regexp(fmt.Sprintf(`game_id=%s;`, gameID), gameIDRes)
	as.Regexp(fmt.Sprintf(`user_id=%s;`, base64Regexp), userIDRes)
	as.Regexp(`username="[a-zA-Z ]+";`, usernameRes)
	as.Regexp(`is_host=false;`, isHostRes)

	res = req.Post(map[string]interface{}{"game_id": "abc"})
	as.Equal(http.StatusNotFound, res.Code)
}

func (as *ActionSuite) Test_Preparation_UserReady() {
	u := gm.CreateGame()
	gameID := u.GameID
	userID := u.UserID

	req := as.ProperRequest("/preparation/user_ready")
	res := req.Post(map[string]interface{}{"game_id": "abc"})
	as.Equal(http.StatusNotFound, res.Code)

	req = as.ProperRequest("/preparation/user_ready")
	res = req.Post(map[string]interface{}{"game_id": gameID, "user_id": "abc"})
	as.Equal(http.StatusNotFound, res.Code)

	req = as.ProperRequest("/preparation/user_ready")
	res = req.Post(map[string]interface{}{"game_id": gameID, "user_id": userID})
	as.Equal(http.StatusOK, res.Code)
}

func (as *ActionSuite) Test_Preparation_ListUsers() {
	u := gm.CreateGame()
	gm.JoinGame(as.c, u.GameID)
	gm.JoinGame(as.c, u.GameID)

	ul := fetchUsersList(as, u.GameID)
	as.Equal(3, len(ul.Users))

	req := as.ProperRequest("/preparation/list_users")
	res := req.Post(map[string]interface{}{"game_id": "abc"})
	as.Equal(http.StatusNotFound, res.Code)
}

func (as *ActionSuite) Test_Preparation_AddPlayer() {
	u := gm.CreateGame()

	req := as.ProperRequest("/preparation/add_player")
	res := req.Post(map[string]interface{}{"game_id": "abc"})
	as.Equal(http.StatusNotFound, res.Code)

	req = as.ProperRequest("/preparation/add_player")
	res = req.Post(map[string]interface{}{"game_id": u.GameID, "user_id": "abc"})
	as.Equal(http.StatusNotFound, res.Code)

	ul := fetchUsersList(as, u.GameID)
	as.Equal(1, len(ul.Users[0].Players))

	req = as.ProperRequest("/preparation/add_player")
	res = req.Post(map[string]interface{}{"game_id": u.GameID, "user_id": u.UserID})
	as.Equal(http.StatusOK, res.Code)

	ul = fetchUsersList(as, u.GameID)
	as.Equal(2, len(ul.Users[0].Players))

	req = as.ProperRequest("/preparation/add_player")
	res = req.Post(map[string]interface{}{"game_id": u.GameID, "user_id": u.UserID})
	as.Equal(http.StatusOK, res.Code)

	ul = fetchUsersList(as, u.GameID)
	as.Equal(3, len(ul.Users[0].Players))

	u.MarkReady()

	req = as.ProperRequest("/preparation/add_player")
	res = req.Post(map[string]interface{}{"game_id": u.GameID, "user_id": u.UserID})
	as.Equal(http.StatusForbidden, res.Code)
}

func (as *ActionSuite) Test_Preparation_RemovePlayer() {
	u := gm.CreateGame()

	req := as.ProperRequest("/preparation/remove_player")
	res := req.Post(map[string]interface{}{"game_id": "abc"})
	as.Equal(http.StatusNotFound, res.Code)

	req = as.ProperRequest("/preparation/remove_player")
	res = req.Post(map[string]interface{}{"game_id": u.GameID, "user_id": "abc"})
	as.Equal(http.StatusNotFound, res.Code)

	req = as.ProperRequest("/preparation/remove_player")
	res = req.Post(map[string]interface{}{"game_id": u.GameID, "user_id": u.UserID})
	as.Equal(http.StatusForbidden, res.Code)

	req = as.ProperRequest("/preparation/add_player")
	res = req.Post(map[string]interface{}{"game_id": u.GameID, "user_id": u.UserID})
	as.Equal(http.StatusOK, res.Code)

	req = as.ProperRequest("/preparation/add_player")
	res = req.Post(map[string]interface{}{"game_id": u.GameID, "user_id": u.UserID})
	as.Equal(http.StatusOK, res.Code)

	req = as.ProperRequest("/preparation/remove_player")
	res = req.Post(map[string]interface{}{"game_id": u.GameID, "user_id": u.UserID})
	as.Equal(http.StatusOK, res.Code)

	ul := fetchUsersList(as, u.GameID)
	as.Equal(2, len(ul.Users[0].Players))

	u.MarkReady()

	req = as.ProperRequest("/preparation/remove_player")
	res = req.Post(map[string]interface{}{"game_id": u.GameID, "user_id": u.UserID})
	as.Equal(http.StatusForbidden, res.Code)
}

func (as *ActionSuite) Test_Preparation_StartGame() {
}

func (as *ActionSuite) Test_Preparation_GameHasStarted() {
}

func fetchUsersList(as *ActionSuite, gameID string) *server.UsersList {
	req := as.ProperRequest("/preparation/list_users")
	res := req.Post(map[string]interface{}{"game_id": gameID})
	as.Equal(http.StatusOK, res.Code)

	usersList := &server.UsersList{}
	err := json.Unmarshal(res.ResponseRecorder.Body.Bytes(), usersList)
	as.Equal(err, nil)
	return usersList
}
