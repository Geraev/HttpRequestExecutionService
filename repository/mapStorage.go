package repository

import (
	. "HttpRequestExecutionService/models"
	"sync"
)

type MapRequestStorage struct {
	mx *sync.RWMutex
	m map[string]DaoRequest
}

func NewMapRequest() *MapRequestStorage {
	return &MapRequestStorage{&sync.RWMutex{}, make(map[string]DaoRequest)}
}

func (r *MapRequestStorage) Put(key string, value DaoRequest)  {
	r.mx.Lock()
	r.m[key] = value
	r.mx.Unlock()
}

func (r *MapRequestStorage) Get(key string) (val DaoRequest, ok bool)  {
	r.mx.RLock()
	val, ok = r.m[key]
	r.mx.RUnlock()
	return
}

func (r *MapRequestStorage) Delete(key string) {
	r.mx.Lock()
	delete(r.m, key)
	r.mx.Unlock()
}