// A very dumb in-memory k/v store

package main

type Datastore struct {
	// not threadsafe!
	data map[string][]byte
}

func NewDatastore() *Datastore {
	return &Datastore{
		data: make(map[string][]byte),
	}
}

func (d *Datastore) Set(key string, value []byte) {
	d.data[key] = value
}

func (d *Datastore) Get(key string) []byte {
	return d.data[key]
}
