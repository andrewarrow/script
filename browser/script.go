package browser

import "github.com/andrewarrow/feedback/wasm"

// var base = "https://fly.script.dev"
var base = "http://localhost:3000"

func RunScriptForFly() {
	guid := Global.Start
	href := Global.Location.Href
	m := map[string]any{}
	m["href"] = href
	go wasm.DoPost(base+"/core/install", m)
	Document.RenderToId("a"+guid, "widget_us_dark", m)
	// <script src="https://cdn.tailwindcss.com" class="">
	// document.createElement('script'); script.src =
	//document.head.appendChild(script)
}
