package Singleton

import "sync"

type singleton struct {
	count int
}

var (
	instance *singleton
	once     sync.Once
)

func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}

func (s *singleton) AddOne() {
	s.count++
}

func (s *singleton) GetCount() int {
	return s.count
}
