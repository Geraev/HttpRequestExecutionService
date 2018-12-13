package repository

import (
	. "HttpRequestExecutionService/models"
	"sync"
)

type MapResponseStorage struct {
	mx *sync.RWMutex
	m  map[string]DaoResponse
}

func NewMapResponse() *MapResponseStorage {
	return &MapResponseStorage{&sync.RWMutex{}, make(map[string]DaoResponse)}
}

func (r *MapResponseStorage) Put(key string, value DaoResponse) {
	r.mx.Lock()
	r.m[key] = value
	r.mx.Unlock()
}

func (r *MapResponseStorage) Get(key string) (val DaoResponse, ok bool) {
	r.mx.RLock()
	val, ok = r.m[key]
	r.mx.RUnlock()
	return
}

func (r *MapResponseStorage) Delete(key string) {
	r.mx.Lock()
	delete(r.m, key)
	r.mx.Unlock()
}
