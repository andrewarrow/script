package app

import (
	"github.com/andrewarrow/feedback/router"
)

func Welcome(c *router.Context, second, third string) {
	if second == "" && third == "" && c.Method == "GET" {
		handleWelcomeIndex(c)
		return
	}
	c.NotFound = true
}

func handleWelcomeIndex(c *router.Context) {
	if len(c.User) > 0 {
		router.Redirect(c, "/core/start")
		return
	}

	send := map[string]any{}
	items := c.All("domain", "order by created_at desc", "")
	send["items"] = items
	c.SendContentInLayout("welcome.html", send, 200)
}
