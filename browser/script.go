package browser

import (
	"fmt"
)

func RunScriptForFly() {
	guid := Global.Start
	fmt.Println("11111", guid)
	//Global.SetWindowFunc("ScriptFlyDevStripe", ScriptFlyDevStripe)
	div := Document.Id("a" + guid)
	div.Set("innerHTML", "working today")
}
