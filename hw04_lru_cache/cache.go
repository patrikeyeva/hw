package hw04lrucache

import "sync"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type cacheItem struct {
	Value    interface{}
	ValueKey Key
}
type lruCache struct {
	capacity int               // ёмкость (количество сохраняемых в кэше элементов)
	queue    List              // очередь [последних используемых элементов] на основе двусвязного списка
	items    map[Key]*ListItem // словарь, отображающий ключ (строка) на элемент очереди
	mu       sync.Mutex
}

func (c *lruCache) Set(key Key, value interface{}) bool { // Добавить значение в кэш по ключу.
	c.mu.Lock()
	defer c.mu.Unlock()

	if item, exist := c.items[key]; exist {
		item.Value = newCacheItem(value, key)
		c.queue.MoveToFront(item)

		return true
	}
	c.queue.PushBack(newCacheItem(value, key))
	elem := c.queue.Back()
	c.queue.MoveToFront(elem)
	c.items[key] = elem

	if c.queue.Len() > c.capacity {
		back := c.queue.Back()

		c.queue.Remove(back)
		if v, ok := back.Value.(cacheItem); ok {
			delete(c.items, v.ValueKey)
		}
	}
	return false
}

func (c *lruCache) Get(key Key) (interface{}, bool) { // Получить значение из кэша по ключу.
	c.mu.Lock()
	defer c.mu.Unlock()

	if item, exist := c.items[key]; exist {
		c.queue.MoveToFront(item)
		if v, ok := item.Value.(cacheItem); ok {
			return v.Value, true
		}

		return nil, false
	}
	return nil, false
}

func (c *lruCache) Clear() { // Очистить кэш.
	c.mu.Lock()
	defer c.mu.Unlock()

	c.queue = NewList()
	c.items = make(map[Key]*ListItem, c.capacity)
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

func newCacheItem(value interface{}, key Key) cacheItem {
	return cacheItem{
		Value:    value,
		ValueKey: key,
	}
}
