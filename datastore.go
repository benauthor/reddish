// A very dumb in-memory k/v store

package main

import "sync"

type Datastore struct {
	sync.RWMutex
	data map[string][]byte
}

func NewDatastore() *Datastore {
	return &Datastore{
		data: make(map[string][]byte),
	}
}

func (d *Datastore) Set(key string, value []byte) {
	d.Lock()
	d.data[key] = value
	d.Unlock()
}

func (d *Datastore) Get(key string) []byte {
	d.RLock()
	ret := d.data[key]
	d.RUnlock()
	return ret
}
