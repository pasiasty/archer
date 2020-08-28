package actions

import (
	"testing"

	"github.com/gobuffalo/envy"
	"github.com/gobuffalo/httptest"
	"github.com/gobuffalo/packr/v2"
	"github.com/gobuffalo/suite"
	"github.com/pasiasty/archer/server"
)

type ActionSuite struct {
	*suite.Action
}

func (as *ActionSuite) BeforeTest(suiteName, testName string) {
	envy.Set("GO_ENV", "test")
	gm = server.CreateGameManager()
}

func (as *ActionSuite) ProperRequest(path string) *httptest.Request {
	req := as.HTML(path)
	req.Headers["User-Agent"] = "Mozilla"
	req.Headers["X-CSRF-Token"] = "test"
	return req
}

func Test_ActionSuite(t *testing.T) {
	action, err := suite.NewActionWithFixtures(App(), packr.New("Test_ActionSuite", "../fixtures"))
	if err != nil {
		t.Fatal(err)
	}

	as := &ActionSuite{
		Action: action,
	}
	suite.Run(t, as)
}
