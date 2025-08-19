package main

import (
	"fmt"

	"golang.design/x/clipboard"

	"github.com/Gjones747/goCopyPaste/api"
	"github.com/Gjones747/goCopyPaste/internal"
	hook "github.com/robotn/gohook"
)

func main() {

	stack := new(api.CopyDeque)

	hook.Register(hook.KeyDown, []string{"c", "cmd"}, func(e hook.Event) {
		fmt.Println("sixe sevenn")

	})

	if err := clipboard.Init(); err != nil {
		fmt.Println("failed to setup clipboard")
	}

	// individual threads
	// so both threads can work together
	go internal.ImageWatcher(stack)
	go internal.TextWatcher(stack)
	// starter := hook.Start()
	// prolly wont need the hook
	// <-hook.Process(starter)
	select{}
}
