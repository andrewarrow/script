package browser

import "github.com/andrewarrow/feedback/wasm"

var base = "https://script.fly.dev"

//var base = "http://localhost:3001"

func RunScriptForFly() {
	//guid := Global.Start
	href := Global.Location.Href
	m := map[string]any{}
	m["href"] = href
	go wasm.DoPost(base+"/core/install", m)

	parent := Document.Id("fly2024").JValue.Get("parentElement")
	theme := "widget_us_dark"
	if Global.Start == "light" {
		theme = "widget_us_light"
	}
	parent.Set("innerHTML", Document.Render(theme, m))
}
