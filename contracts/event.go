package contracts

import (
	"container/list"
	"sync"
)

type EventStack struct {
	stacks *list.List
	lock   *sync.RWMutex
}

func NewEvent() *EventStack {

	return &EventStack{
		stacks: list.New(),
		lock:   &sync.RWMutex{},
	}
}

func (it *EventStack) Push(elem *Payload) {
	it.lock.Lock()
	defer it.lock.Unlock()
	it.stacks.PushBack(elem)
}

func (it *EventStack) Pop() *Payload {
	it.lock.Lock()
	defer it.lock.Unlock()
	elem := it.stacks.Front()
	if elem != nil {
		it.stacks.Remove(elem)
		return elem.Value.(*Payload)
	}
	return nil
}

func (it *EventStack) First() *Payload {
	elem := it.stacks.Back()
	if elem != nil {
		return elem.Value.(*Payload)
	}
	return nil
}

func (it *EventStack) Len() int {
	return it.stacks.Len()
}

func (it *EventStack) Empty() bool {
	return it.stacks.Len() == 0
}
