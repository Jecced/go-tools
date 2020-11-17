package comp

import (
	"sync"
)

type key interface{}

type set struct {
	m    map[key]bool
	sync bool
	sync.RWMutex
}

// 非线程安全的set
func Set(items ...key) *set {
	return commConstructor(false, items...)
}

// 线程安全的set
func CSet(items ...key) *set {
	return commConstructor(true, items...)
}

// 通用初始化
func commConstructor(sync bool, items ...key) *set {
	cache := map[key]bool{}
	for _, item := range items {
		cache[item] = true
	}
	return &set{m: cache, sync: sync}
}

func (s *set) Add(item key) {
	if s.sync {
		s.Lock()
		defer s.Unlock()
	}
	s.m[item] = true
}

func (s *set) Remove(item key) {
	if s.sync {
		s.Lock()
		defer s.Unlock()
	}
	delete(s.m, item)
}

func (s *set) Has(item key) bool {
	if s.sync {
		s.RLock()
		defer s.RUnlock()
	}
	_, ok := s.m[item]
	return ok
}

func (s *set) Size() int {
	return len(s.m)
}

func (s *set) Clear() {
	s.RLock()
	defer s.RUnlock()
	s.m = map[key]bool{}
}

func (s *set) IsEmpty() bool {
	if s.Size() == 0 {
		return true
	}
	return false
}

func (s *set) List() []key {
	if s.sync {
		s.RLock()
		defer s.RUnlock()
	}
	var list []key
	for item := range s.m {
		list = append(list, item)
	}
	return list
}
