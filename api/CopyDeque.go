package api

import "errors"

type CopyDeque struct {
	data [5][]byte
}

func (deque *CopyDeque) PushFrontLeakBack(element []byte) {
	length := len(deque.data)

	i := 1
	prev := deque.data[0]
	deque.data[0] = element
	for i < length {
		curr := deque.data[i]
		deque.data[i] = prev
		prev = curr
		i++
	}

}

func (deque CopyDeque) GetFront() ([]byte, error) {
	if deque.data[0] == nil {
		return []byte{}, errors.New("nothing in deque")
	} else {
		return deque.data[0], nil
	}
}
