package userstore

import (
	"fmt"
	"sync"
)

type LocalUserStore struct {
	storage map[string]string
	lock    *sync.RWMutex
}

func NewLocalUserStore() *LocalUserStore {
	return &LocalUserStore{
		storage: map[string]string{},
		lock:    &sync.RWMutex{},
	}
}

func (l *LocalUserStore) SaveUser(userName, password string) error {
	defer l.lock.Unlock()
	l.lock.Lock()
	if _, exist := l.storage[userName]; exist {
		return fmt.Errorf("username %s has been taken", userName)
	}
	l.storage[userName] = password
	return nil
}

func (l *LocalUserStore) GetUser(userName string) (bool, string) {
	defer l.lock.Lock()
	l.lock.Lock()
	password, exist := l.storage[userName]
	if !exist {
		return false, ""
	}
	return true, password
}
