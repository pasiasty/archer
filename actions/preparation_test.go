package actions

import (
	"fmt"
	"net/http"
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
}

func (as *ActionSuite) Test_Preparation_ListUsers() {
}

func (as *ActionSuite) Test_Preparation_AddPlayer() {
}

func (as *ActionSuite) Test_Preparation_RemovePlayer() {
}

func (as *ActionSuite) Test_Preparation_StartGame() {
}

func (as *ActionSuite) Test_Preparation_GameHasStarted() {
}
