package internal

import (
	"context"

	"github.com/Gjones747/goCopyPaste/api"
	"golang.design/x/clipboard"
)

func ImageWatcher(stack *api.CopyDeque) {
	imageWatcher := clipboard.Watch(context.Background(), clipboard.FmtImage)

	for data := range imageWatcher {
		copyItem := api.CopyItem{
			ItemData: data,
			ItemType: api.Png,
		}
		stack.PushFrontLeakBack(copyItem)
	}
}
