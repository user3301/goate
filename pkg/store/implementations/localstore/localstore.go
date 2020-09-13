package localstore

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/user3301/grpclab/pkg/types"
)

// LocalUserStore storage for user credentials
type LocalUserStore struct {
	storage map[string]string
	lock    *sync.RWMutex
}

// NewLocalUserStore LocalUserStore constructor
func NewLocalUserStore() *LocalUserStore {
	return &LocalUserStore{
		storage: map[string]string{},
		lock:    &sync.RWMutex{},
	}
}

// SaveUser save use credentials
func (l *LocalUserStore) CreateUser(_ context.Context, userDetails types.UserDetails) error {
	log.Printf("user to be registered %v", userDetails)
	defer l.lock.Unlock()
	l.lock.Lock()
	if _, exist := l.storage[userDetails.Username]; exist {
		return fmt.Errorf("username %s has been taken", userDetails.Username)
	}
	l.storage[userDetails.Username] = userDetails.Password
	return nil
}

// GetUser Get user credentials by user name
func (l *LocalUserStore) Verify(_ context.Context, userDetails types.UserDetails) (verified bool, reason string) {
	log.Print("user store entered")
	defer l.lock.Unlock()
	l.lock.Lock()
	password, exist := l.storage[userDetails.Username]
	if !exist {
		return false, fmt.Sprintf("username %s not exist", userDetails.Username)
	}
	if password != userDetails.Password {
		return false, "password incorrect"
	}
	return true, "ok"
}

// UpdateUser update user record
func (l LocalUserStore) UpdateUser(_ context.Context, details types.UserDetails) (updateSuccess bool, err error) {
	defer l.lock.Unlock()
	l.lock.Lock()
	_, exist := l.storage[details.Username]
	if !exist {
		return false, fmt.Errorf("record not exist")
	}
	l.storage[details.Username] = details.Password
	return true, nil
}
