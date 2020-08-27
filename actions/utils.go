package actions

import (
	"github.com/gobuffalo/buffalo"
)

func setCookie(c buffalo.Context, key, value string) {
	c.Cookies().SetWithPath(key, value, "/")
}
