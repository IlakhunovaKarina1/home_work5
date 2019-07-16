package main


import "container/list"

type LRUCache struct {
	queue *list.List
	cache map[string]*list.Element
	cSize int
}

func (l *LRUCache) Check(s string) bool {
	if _, ok := l.cache[s]; ok {
		return true
	} else {
		return false
	}
}
