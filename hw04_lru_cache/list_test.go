package hw04lrucache

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestList(t *testing.T) {
	t.Run("empty list", func(t *testing.T) {
		l := NewList()

		require.Equal(t, 0, l.Len())
		require.Nil(t, l.Front())
		require.Nil(t, l.Back())
	})

	t.Run("push front", func(t *testing.T) {
		l := NewList()

		l.PushFront(10) // [10]
		require.Nil(t, l.Front().Prev)
		require.Nil(t, l.Front().Next)

		l.PushFront(20) // [20, 10]
		require.Nil(t, l.Front().Prev)

		l.PushFront(30) // [30, 20, 10]
		require.Nil(t, l.Front().Prev)

		require.Equal(t, 3, l.Len())
		require.Equal(t, 20, l.Front().Next.Value)
		require.Equal(t, 10, l.Front().Next.Next.Value)
		require.Equal(t, 20, l.Back().Prev.Value)
		require.Equal(t, 30, l.Back().Prev.Prev.Value)
		require.Nil(t, l.Back().Prev.Prev.Prev)
		require.Nil(t, l.Back().Next)
	})

	t.Run("push back", func(t *testing.T) {
		l := NewList()

		l.PushBack(10) // [10]
		require.Nil(t, l.Back().Prev)
		require.Nil(t, l.Back().Next)

		l.PushBack(20) // [10, 20]
		require.Nil(t, l.Back().Next)

		l.PushBack(30) // [10, 20, 30]
		require.Nil(t, l.Back().Next)

		require.Equal(t, 3, l.Len())
		require.Equal(t, 20, l.Back().Prev.Value)
		require.Equal(t, 10, l.Back().Prev.Prev.Value)
	})

	t.Run("remove", func(t *testing.T) {
		l := NewList()

		l.PushBack(10) // [10]
		l.PushBack(20) // [10, 20]
		l.PushBack(30) // [10, 20, 30]
		require.Equal(t, 3, l.Len())

		middle := l.Front().Next
		l.Remove(middle)
		require.Equal(t, 2, l.Len())
		require.Equal(t, 30, l.Front().Next.Value) // 10 -> 30
		require.Equal(t, 10, l.Back().Prev.Value)  // 10 <- 30

		back := l.Back()
		l.Remove(back)
		require.Equal(t, 1, l.Len())
		require.Nil(t, l.Back().Prev)
		require.Nil(t, l.Back().Next)

		l.Remove(l.Front())
		require.Equal(t, 0, l.Len())
		require.Nil(t, l.Front())
		require.Nil(t, l.Back())
	})

	t.Run("complex", func(t *testing.T) {
		l := NewList()

		l.PushFront(10) // [10]
		l.PushBack(20)  // [10, 20]
		l.PushBack(30)  // [10, 20, 30]
		require.Equal(t, 3, l.Len())

		middle := l.Front().Next // 20
		l.Remove(middle)         // [10, 30]
		require.Equal(t, 2, l.Len())

		for i, v := range [...]int{40, 50, 60, 70, 80} {
			if i%2 == 0 {
				l.PushFront(v)
			} else {
				l.PushBack(v)
			}
		} // [80, 60, 40, 10, 30, 50, 70]

		require.Equal(t, 7, l.Len())
		require.Equal(t, 80, l.Front().Value)
		require.Equal(t, 70, l.Back().Value)

		l.MoveToFront(l.Front()) // [80, 60, 40, 10, 30, 50, 70]
		l.MoveToFront(l.Back())  // [70, 80, 60, 40, 10, 30, 50]

		elems := make([]int, 0, l.Len())
		for i := l.Front(); i != nil; i = i.Next {
			elems = append(elems, i.Value.(int))
		}
		require.Equal(t, []int{70, 80, 60, 40, 10, 30, 50}, elems)
	})
}
