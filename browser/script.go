package browser

import "github.com/andrewarrow/feedback/wasm"

var base = "https://fly.script.dev"

//var base = "http://localhost:3001"

func RunScriptForFly() {
	//guid := Global.Start
	href := Global.Location.Href
	m := map[string]any{}
	m["href"] = href
	go wasm.DoPost(base+"/core/install", m)

	parent := Document.Id("fly2024").JValue.Get("parentElement")
	parent.Set("innerHTML", Document.Render("widget_us_dark", m))
}
