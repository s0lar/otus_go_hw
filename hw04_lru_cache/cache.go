package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

//	Для чего нужна структура cacheItem?
//	В интерфейсах она не учавствует. Get мог бы возвращать cacheItem, но по сигнатуре и тестам - не подходит
//	.. можно пояснить или пример увидеть использования?
//	type cacheItem struct {
//		key   Key
//		value interface{}
//	}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	//	Key already in cache
	if val, ok := c.items[key]; ok {
		val.Value = value
		c.queue.MoveToFront(val)
		return true
	}

	//	Remove last if overflow queue
	if c.queue.Len() == c.capacity {
		last := c.queue.Back()
		for k, v := range c.items {
			if v == last {
				delete(c.items, k)
				c.queue.Remove(last)
			}
		}
	}

	//	New key in cache
	li := c.queue.PushFront(value)
	c.items[key] = li

	return false
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	if val, ok := c.items[key]; ok {
		c.queue.MoveToFront(val)
		//	!!! Внимание
		//	Из-за того, что c.queue.MoveToFront - удаляет переменную и создает новую в начале слайса
		//	меняется ее адрес в c.queue.buffer поэтому нужно обновить ее в map
		//	возмжно решение не правильное и cache не должен знать об особенностях реализации list!!!
		c.items[key] = c.queue.Front()
		return val.Value, true
	}
	return nil, false
}

func (c *lruCache) Clear() {
	c.queue = NewList()
	c.items = make(map[Key]*ListItem, c.capacity)
}
