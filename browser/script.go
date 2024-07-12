package browser

import (
	"fmt"

	"github.com/andrewarrow/feedback/wasm"
)

var Global *wasm.Global
var Document *wasm.Document

func RunScriptForFly() {
	fmt.Println("11111")
	Global.SetWindowFunc("ScriptFlyDevStripe", ScriptFlyDevStripe)
}

func ScriptFlyDevStripe(guid string) {
	fmt.Println("21111")
}
