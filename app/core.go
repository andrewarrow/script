package app

import (
	"fmt"

	"github.com/andrewarrow/feedback/router"
)

func Core(c *router.Context, second, third string) {
	if second == "start" && third == "" && c.Method == "GET" {
		handleStart(c)
		return
	}
	if second == "stripe" && third != "" && c.Method == "GET" {
		handleStripeShow(c, third)
		return
	}
	if second == "stripe" && third == "" && c.Method == "POST" {
		handleStripePost(c)
		return
	}
	if second == "about-us" && third == "" && c.Method == "GET" {
		handleAboutUs(c)
		return
	}
	if second == "privacy" && third == "" && c.Method == "GET" {
		handlePrivacy(c)
		return
	}
	if second == "terms" && third == "" && c.Method == "GET" {
		handleTerms(c)
		return
	}
	if second == "register" && third == "" && c.Method == "GET" {
		handleRegister(c)
		return
	}
	if second == "login" && third == "" && c.Method == "GET" {
		handleLogin(c)
		return
	}
	if second == "register" && third == "" && c.Method == "POST" {
		router.HandleCreateUserAutoForm(c, "")
		return
	}
	if second == "login" && third == "" && c.Method == "POST" {
		router.HandleCreateSessionAutoForm(c)
		return
	}
	if second == "logout" && third == "" && c.Method == "DELETE" {
		router.DestroySession(c)
		return
	}
	c.NotFound = true
}

func handleIndex(c *router.Context) {
	send := map[string]any{}
	c.SendContentInLayout("welcome.html", send, 200)
}

func handleRegister(c *router.Context) {
	send := map[string]any{}
	c.SendContentInLayout("register.html", send, 200)
}
func handleLogin(c *router.Context) {
	send := map[string]any{}
	c.SendContentInLayout("login.html", send, 200)
}

func handlePrivacy(c *router.Context) {
	send := map[string]any{}
	c.SendContentInLayout("privacy.html", send, 200)
}
func handleTerms(c *router.Context) {
	send := map[string]any{}
	c.SendContentInLayout("terms.html", send, 200)
}
func handleAboutUs(c *router.Context) {
	send := map[string]any{}
	c.SendContentInLayout("about_us.html", send, 200)
}
func handleStart(c *router.Context) {
	send := map[string]any{}
	items := c.All("stripe", "where user_id=$1", "", c.User["id"])
	send["items"] = items
	c.SendContentInLayout("start.html", send, 200)
}
func handleStripePost(c *router.Context) {
	c.ReadJsonBodyIntoParams()
	c.Params["user_id"] = c.User["id"]
	c.ValidateAndInsert("stripe")
	send := map[string]any{}
	c.SendContentAsJson(send, 200)
}
func handleStripeShow(c *router.Context, guid string) {
	send := map[string]any{}
	item := c.One("stripe", "where guid=$1", guid)
	send["item"] = item
	script := fmt.Sprintf(scriptTemplate, guid)
	send["script"] = fmt.Sprintf(`<script type="text/javascript">%s</script>`, script)
	c.SendContentInLayout("stripe.html", send, 200)
}

var scriptTemplate = `(function(f,l,y,d,e,v,x){ f[y]=f[y]||function(){(f[y].q=f[y].q||[]).push(arguments)}; v=l.createElement(e);t.async=1;v.src="https://script.fly.dev/core/tag/"+i; x=l.getElementsByTagName(d)[0];x.parentNode.insertBefore(v,x); })(window, document, "flydev", "script", "%s");`
