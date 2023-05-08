package Singleton

import (
	"sync"
	"testing"
)

func TestGetInstance(t *testing.T) {
	// 通过GetInstance获取单例实例
	instance1 := GetInstance()
	instance2 := GetInstance()

	if instance1 == nil {
		t.Error("Expected non-nil instance, but got nil")
	}

	if instance2 == nil {
		t.Error("Expected non-nil instance, but got nil")
	}

	if instance1 != instance2 {
		t.Error("Expected same instance, but got different instances")
	}
}

func TestGetInstanceConcurrent(t *testing.T) {
	// 创建WaitGroup和Mutex来确保并发安全性
	var wg sync.WaitGroup
	var mu sync.Mutex

	// 并发获取单例实例
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			instance := GetInstance()

			// 使用锁来确保单例实例的并发安全性
			mu.Lock()
			if instance == nil {
				t.Error("Expected non-nil instance, but got nil")
			}
			mu.Unlock()
		}()
	}

	wg.Wait()
}
