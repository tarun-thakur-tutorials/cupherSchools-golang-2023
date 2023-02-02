package main

import "github.com/spf13/cast"

const (
	LRU = 1
	MRU = 2
)

type Cache struct {
	topIndex int
	mark     int
	capacity int
	array    []interface{}
}

func NewCache(capacity int, strategyId int) *Cache {
	c := &Cache{
		topIndex: -1,
		capacity: capacity,
		array:    make([]interface{}, capacity),
	}

	switch strategyId {
	case LRU:
		c.mark = 0
	case MRU:
		c.mark = capacity - 1 // last index
	}
	return c
}

func (c Cache) IsCacheFull() bool {
	return c.topIndex == c.capacity-1
}

func (c *Cache) Insert(strategyId int, data interface{}) string {
	switch strategyId {
	case LRU:
		if c.IsCacheFull() {
			// replace the element at the mark
			c.array[c.mark] = data
			// update the mark
			if c.mark < c.capacity-1 {
				c.mark++
			} else {
				c.mark = 0
			}
		} else {
			c.topIndex++
			c.array[c.topIndex] = data
		}
	case MRU:
		if c.IsCacheFull() {
			c.array[c.mark] = data
			if c.mark == 0 {
				c.mark = c.capacity - 1
			} else {
				c.mark--
			}
		} else {
			c.topIndex++
			c.array[c.topIndex] = data
		}
	}

	output := ""
	for _, elem := range c.array {
		// output = fmt.Sprintf("%v %v", output, elem)
		output += " " + cast.ToString(elem)
	}
	return output
}
