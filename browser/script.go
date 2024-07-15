package browser

import "github.com/andrewarrow/feedback/wasm"

func RunScriptForFly() {
	guid := Global.Start
	href := Global.Location.Href
	m := map[string]any{}
	m["href"] = href
	go wasm.DoPost("https://fly.script.dev/core/install", m)
	Document.RenderToId("a"+guid, "widget_us_dark", m)
	// <script src="https://cdn.tailwindcss.com" class="">
	// document.createElement('script'); script.src =
	//document.head.appendChild(script)
}
