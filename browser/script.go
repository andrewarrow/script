package browser

func RunScriptForFly() {
	guid := Global.Start
	m := map[string]any{}
	Document.RenderToId("a"+guid, "widget_stripe", m)
}
