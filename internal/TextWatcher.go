package internal

import (
	"context"

	"github.com/Gjones747/goCopyPaste/api"
	"golang.design/x/clipboard"
)

func TextWatcher(stack *api.CopyDeque) {
	imageWatcher := clipboard.Watch(context.Background(), clipboard.FmtText)

	for data := range imageWatcher {
		copyItem := api.CopyItem{
			ItemData: data,
			ItemType: api.Png,
		}
		stack.PushFrontLeakBack(copyItem)
	}
}
