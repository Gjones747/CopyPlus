package main

import (
	"fmt"

	"golang.design/x/clipboard"

	"github.com/Gjones747/goCopyPaste/api"

	hook "github.com/robotn/gohook"
)

func main() {

	hook.Register(hook.KeyDown, []string{"c", "cmd"}, func(e hook.Event) {
		fmt.Println("sixe sevenn")

		hook.End()
	})

	stack := new(api.CopyDeque)
	if err := clipboard.Init(); err != nil {
		fmt.Println("failed to setup clipboard")
	}
	stack.PushFrontLeakBack(clipboard.Read(clipboard.FmtText))

	if front, err := stack.GetFront(); err == nil {
		fmt.Println(string(front))
	} else {
		fmt.Println(err.Error())
	}
	starter := hook.Start()
	<-hook.Process(starter)
}
