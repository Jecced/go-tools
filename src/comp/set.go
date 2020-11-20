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

type Set interface {
	Add(item key) bool
	Remove(item key) bool
	Has(item key) bool
	Size() int
	Clear()
	IsEmpty() bool
	List() []key
}

// 非线程安全的set
func NewSet(items ...key) Set {
	return commConstructor(false, items...)
}

// 线程安全的set
func NewCSet(items ...key) Set {
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

// 返回是否 add 成功, 如果原本这个值已经存在返回false
func (s *set) Add(item key) bool {
	if s.sync {
		s.Lock()
		defer s.Unlock()
	}
	if s.Has(item) {
		return false
	}
	s.m[item] = true
	return true
}

func (s *set) Remove(item key) bool {
	if s.sync {
		s.Lock()
		defer s.Unlock()
	}
	has := s.Has(item)
	delete(s.m, item)
	return has
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
