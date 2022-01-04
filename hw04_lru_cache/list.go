package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	buffer []*ListItem
}

func (l *list) Len() int {
	return len(l.buffer)
}

func (l *list) Front() *ListItem {
	if l.Len() == 0 {
		return nil
	}
	return l.buffer[0]
}

func (l *list) Back() *ListItem {
	if l.Len() == 0 {
		return nil
	}
	return l.buffer[len(l.buffer)-1]
}

func (l *list) PushFront(v interface{}) *ListItem {
	li := ListItem{v, nil, nil}
	l.buffer = append([]*ListItem{&li}, l.buffer...)

	l.setNext(0)
	l.setPrev(1)
	return &li
}

func (l *list) PushBack(v interface{}) *ListItem {
	li := ListItem{v, nil, nil}
	l.buffer = append(l.buffer, &li)

	l.setPrev(l.Len() - 1)
	l.setNext(l.Len() - 2)
	return &li
}

func (l *list) Remove(removeLi *ListItem) {
	removeIndex := 0
	for i, li := range l.buffer {
		if li == removeLi {
			removeIndex = i
		}
	}
	l.buffer = append(l.buffer[:removeIndex], l.buffer[removeIndex+1:]...)
	l.setNext(removeIndex - 1)
	l.setPrev(removeIndex)
}

func (l *list) MoveToFront(toFrontLi *ListItem) {
	l.Remove(toFrontLi)
	l.PushFront(toFrontLi.Value)
}

func (l *list) setPrev(currIndex int) {
	switch {
	case l.Len() == 0 || currIndex < 0 || currIndex > l.Len()-1:
		return
	case currIndex == 0:
		l.buffer[currIndex].Prev = nil
	case currIndex > 0 && l.Len() > currIndex:
		l.buffer[currIndex].Prev = l.buffer[currIndex-1]
	}
}
func (l *list) setNext(currIndex int) {
	switch nextIndex := currIndex + 1; {
	case l.Len() == 0 || currIndex < 0 || currIndex > l.Len()-1:
		return
	case nextIndex == l.Len():
		l.buffer[currIndex].Next = nil
	case currIndex >= 0 && l.Len() > currIndex:
		l.buffer[currIndex].Next = l.buffer[nextIndex]
	}
}

func NewList() List {
	return new(list)
}
