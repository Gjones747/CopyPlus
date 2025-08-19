package api

import "errors"

type ItemType int

const (
	text ItemType = iota
	Png
	Jpeg
)

type CopyItem struct {
	ItemData []byte
	ItemType ItemType
	init     bool
}

type CopyDeque struct {
	data [5]CopyItem
}

func (deque *CopyDeque) PushFrontLeakBack(element CopyItem) {
	for i := len(deque.data) - 1; i >= 1; i-- {
		deque.data[i] = deque.data[i-1]
	}

	deque.data[0] = element
}

func (deque *CopyDeque) GetFront() (CopyItem, error) {
	if !deque.data[0].init {
		return CopyItem{}, errors.New("nothing in deque")
	} else {
		return deque.data[0], nil
	}
}
