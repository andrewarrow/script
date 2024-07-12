package browser

import (
	"fmt"
)

func RunScriptForFly() {
	fmt.Println("11111")
	Global.SetWindowFunc("ScriptFlyDevStripe", ScriptFlyDevStripe)
}

func ScriptFlyDevStripe(guid string) {
	fmt.Println("21111")
}
