package browser

func RunScriptForFly() {
	guid := Global.Start
	m := map[string]any{}
	Document.RenderToId("a"+guid, "widget_us_dark", m)
	// <script src="https://cdn.tailwindcss.com" class="">
	// document.createElement('script'); script.src =
	//document.head.appendChild(script)
}
