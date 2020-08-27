package server

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io"
	"math/rand"

	"github.com/gobuffalo/buffalo"
)

// SetCookie sets cookie value for whole domain.
func SetCookie(c buffalo.Context, key, value string) {
	c.Cookies().SetWithPath(key, value, "/")
}

func selectNewKey(m map[string]bool) string {
	h := sha1.New()
	for {
		io.WriteString(h, fmt.Sprintf("%d", rand.Intn(10000)))
		res := base64.URLEncoding.EncodeToString(h.Sum(nil))

		if _, ok := m[res]; !ok {
			return res
		}
	}
}
