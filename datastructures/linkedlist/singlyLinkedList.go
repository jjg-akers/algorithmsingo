package datastructures

import "errors"

type el struct {
	okChan chan bool
	next   *el
}

type list struct {
	length int
	head   *el
	tail   *el
}

// a simple FIFO linked list
func newList() *list {

	return &list{
		length: 0,
		head:   nil,
		tail:   nil,
	}
}

//Add will place new elements at the head
func (l *list) add(ch chan bool) {

	newEl := &el{
		okChan: ch,
	}

	if l.head == nil {
		// this is first element
		l.head = newEl
		l.tail = newEl
		l.length++

	} else {
		// assigne new el as tails next
		l.tail.next = newEl

		//reassign tail
		l.tail = newEl

	}

}

//Pop will pull elements from the tail
func (l *list) pop() (*el, error) {
	if l.head == nil {
		return nil, errors.New("List is empty")
	}

	newHead := l.head.next

	toReturn := l.head

	//reassign head
	l.head = newHead

	l.length--

	return toReturn, nil
}
