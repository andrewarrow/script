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
	if second == "tag" && third != "" && c.Method == "GET" {
		handleTag(c, third)
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

func handleTag(c *router.Context, guid string) {
	//Access-Control-Allow-Origin: *
	//Access-Control-Allow-Methods: GET
	//Access-Control-Allow-Headers: Content-Type
	c.Writer.Header().Set("Content-Length", fmt.Sprintf("%d", len(actionScript)))
	c.Writer.Header().Set("Content-Type", "application/javascript")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.WriteHeader(200)
	c.Writer.Write([]byte(actionScript))
}

var scriptTemplate = `function fetchAndExecuteScript(url) { const script = document.createElement('script'); script.src = url; script.onload = () => { console.log("Script loaded successfully"); }; script.onerror = (error) => { console.error("Error loading script from", error); }; document.head.appendChild(script); } fetchAndExecuteScript('https://script.fly.dev/core/tag/%s');`

var actionScript = `(function() { const newDiv = document.createElement('div'); newDiv.innerHTML = 'testing'; document.currentScript.parentNode.insertBefore(newDiv, document.currentScript.nextSibling); })();`
