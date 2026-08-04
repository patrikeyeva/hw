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
	front *ListItem
	back  *ListItem
	len   int
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *ListItem {
	return l.front
}

func (l *list) Back() *ListItem {
	return l.back
}

func (l *list) PushFront(v interface{}) *ListItem {
	// create item
	item := &ListItem{
		Value: v,
		Next:  l.front,
		Prev:  nil,
	}

	if l.front != nil {
		l.front.Prev = item
	}
	l.front = item
	if l.back == nil {
		l.back = item
	}
	l.len++

	return item
}

func (l *list) PushBack(v interface{}) *ListItem {
	// create item
	item := &ListItem{
		Value: v,
		Next:  nil,
		Prev:  l.back,
	}
	if l.back != nil {
		l.back.Next = item
	}
	l.back = item

	if l.front == nil {
		l.front = item
	}
	l.len++

	return item
}

func (l *list) Remove(i *ListItem) {
	switch {
	case l.front == l.back:
		l.front = nil
		l.back = nil

	case i == l.front:
		i.Next.Prev = nil
		l.front = i.Next

	case i == l.back:
		i.Prev.Next = nil
		l.back = i.Prev

	default:
		i.Prev.Next = i.Next
		i.Next.Prev = i.Prev
	}
	l.len--
}

func (l *list) MoveToFront(i *ListItem) {
	switch i {
	case l.front:
		return
	case l.back:
		i.Prev.Next = nil
		l.back = i.Prev

		i.Prev = nil
		i.Next = l.front

		l.front.Prev = i
		l.front = i

	default:
		i.Prev.Next = i.Next
		i.Next.Prev = i.Prev

		i.Prev = nil
		i.Next = l.front

		l.front.Prev = i
		l.front = i
	}
}

func NewList() List {
	return new(list{
		front: nil,
		back:  nil,
		len:   0,
	})
}
