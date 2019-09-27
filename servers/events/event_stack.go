package events

import (
	"container/list"
	"github.com/9299381/wego/contracts"
	"sync"
)

type eventStack struct {
	stacks *list.List
	lock   *sync.RWMutex
}

func newEventStack() *eventStack {

	return &eventStack{
		stacks: list.New(),
		lock:   &sync.RWMutex{},
	}
}

func (it *eventStack) Push(elem *contracts.Payload) {
	it.lock.Lock()
	defer it.lock.Unlock()
	it.stacks.PushBack(elem)
}

func (it *eventStack) Pop() *contracts.Payload {
	it.lock.Lock()
	defer it.lock.Unlock()
	elem := it.stacks.Front()
	if elem != nil {
		it.stacks.Remove(elem)
		return elem.Value.(*contracts.Payload)
	}
	return nil
}

func (it *eventStack) First() *contracts.Payload {
	elem := it.stacks.Back()
	if elem != nil {
		return elem.Value.(*contracts.Payload)
	}
	return nil
}

func (it *eventStack) Len() int {
	return it.stacks.Len()
}

func (it *eventStack) Empty() bool {
	return it.stacks.Len() == 0
}
