package main

import "sync"

func NewInMemorySixelStore() *InMemorySixelStore {
	return &InMemorySixelStore{
		map[string]string{},
		sync.RWMutex{},
	}
}

type InMemorySixelStore struct {
	store map[string]string
	lock  sync.RWMutex
}

func (i *InMemorySixelStore) GetSixelImage(id string) string {
	i.lock.Lock()
	defer i.lock.Unlock()
	return i.store[id]
}

func (i *InMemorySixelStore) StoreSixelImage(id, image string) {
	i.lock.Lock()
	defer i.lock.Unlock()
	i.store[id] = image
}
