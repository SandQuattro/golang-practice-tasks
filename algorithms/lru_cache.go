package main

import (
	"container/list"
	"errors"
	"log"
)

type dataEntry struct {
	key   any
	value any
}

type Cache struct {
	capacity int
	cache    map[any]*list.Element
	eviction *list.List
}

func NewCache(capacity int) *Cache {
	return &Cache{
		capacity: capacity,
		cache:    make(map[any]*list.Element, capacity),
		eviction: list.New(),
	}
}

func (c *Cache) Set(key, value any) {
	// проверяем, есть ли в кеше элемент?
	if element, exists := c.cache[key]; exists {
		// если да, перемещаем его вперед в списке
		c.eviction.MoveToFront(element)
		// обновляем значение у элемента и выходим
		element.Value.(*dataEntry).value = value
		return
	}

	// если элемента еще нет в кеше
	// создаем элемент из entry
	element := c.eviction.PushFront(&dataEntry{key, value})
	// кладем его в кеш
	c.cache[key] = element

	// выполняем вытеснение least recently used LRU элемента
	if c.eviction.Len() > c.capacity {
		lruElement := c.eviction.Back()
		if lruElement != nil {
			delete(c.cache, lruElement.Value.(*dataEntry).key)
			c.eviction.Remove(lruElement)
		}
	}
}

func (c *Cache) Get(key any) (any, error) {
	if val, exists := c.cache[key]; exists {
		// Если ключ есть в кеше, двигаем элемент в начало списка
		c.eviction.MoveToFront(val)
		// и возвращаем искомое значение
		return val.Value.(*dataEntry).value, nil
	}

	return nil, errors.New("key not found")
}

func printEvictionElements(list *list.List) {
	for e := list.Front(); e != nil; e = e.Next() {
		log.Println(e.Value)
	}
}

func main() {
	cache := NewCache(2)
	log.Printf("our empty cache: %v", cache)

	cache.Set("key1", "value1")
	cache.Set("key2", "value2")
	cache.Set("key3", "value3")
	cache.Set("key4", "value4")
	log.Printf("our filled cache: %v", cache)

	val, err := cache.Get("key4")
	if err != nil {
		log.Println(err)
	}
	log.Println(val)

	if val, err = cache.Get("key3"); err != nil {
		log.Println(err)
	}
	log.Println(val)

	if val, err = cache.Get("key2"); err != nil {
		log.Println(err)
	}
	log.Println(val)

	if val, err = cache.Get("key1"); err != nil {
		log.Println(err)
	}
	log.Println(val)

	log.Println(cache)

	printEvictionElements(cache.eviction)
}
