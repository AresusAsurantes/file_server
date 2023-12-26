package cache

import "sync"

var CommonCache sync.Map

func Put(key, value any) {
	CommonCache.Store(key, value)
}

func Get(key any) (value any, ok bool) {
	return CommonCache.Load(key)
}
