package db

import (
	"sync"

	"github.com/TOMO-CAT/UserManagementSystem/pkg/util"
)

var (
	PtrUserRegisterMemoryLock = NewMemoryLock() // 用户注册内存锁
)

type MemoryLock struct {
	mtx       *sync.Mutex
	lockItems *util.Set
}

func NewMemoryLock() *MemoryLock {
	return &MemoryLock{
		mtx:       &sync.Mutex{},
		lockItems: util.NewSet(),
	}
}

// TryLock 尝试加锁, 加锁成功返回 true
func (m *MemoryLock) TryLock(item interface{}) bool {
	m.mtx.Lock()
	defer m.mtx.Unlock()

	if m.lockItems.Contains(item) {
		return false
	}
	m.lockItems.Add(item)
	return true
}

// Unlock 解锁
func (m *MemoryLock) Unlock(item interface{}) {
	m.mtx.Lock()
	m.lockItems.Remove(item)
	m.mtx.Unlock()
}
