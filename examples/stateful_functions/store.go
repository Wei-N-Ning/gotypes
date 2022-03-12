package main

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"os"
	"sync"
	"time"
)

type state map[string]int

type Store struct {
	sync.RWMutex

	filename string
	latencyF func()
}

func (store *Store) getState() state {
	store.RLock()
	defer store.RUnlock()
	store.latencyF()

	bs, err := ioutil.ReadFile(store.filename)
	if err != nil {
		panic(err)
	}
	st := state{}
	if err := json.Unmarshal(bs, &st); err != nil {
		panic(err)
	}
	return st
}

func (store *Store) saveState(st state) {
	store.Lock()
	defer store.Unlock()
	store.latencyF()

	bs, err := json.Marshal(st)
	if err != nil {
		panic(err)
	}
	if err := ioutil.WriteFile(store.filename, bs, 0700); err != nil {
		panic(err)
	}
}

func (store *Store) Upsert(word string) {
	st := store.getState()
	if x, ok := st[word]; ok {
		st[word] = x + 1
	} else {
		st[word] = 1
	}
	store.saveState(st)
}

func (store *Store) Read(word string) int {
	return store.getState()[word]
}

func Spinup(minLat int, maxLat int) *Store {
	filename := "/tmp/examples_stateful_functions"
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defaultSt := state{}
	bs, err := json.Marshal(defaultSt)
	if err != nil {
		panic(err)
	}
	_, err = f.Write(bs)
	if err != nil {
		panic(err)
	}
	if err := f.Close(); err != nil {
		panic(err)
	}
	store := Store{
		filename: filename,
		latencyF: func() {
			time.Sleep(time.Duration(rand.Intn(maxLat-minLat)+minLat) * time.Millisecond)
		},
	}
	return &store
}

func (store *Store) Teardown() {
	if _, err := os.Stat(store.filename); err != nil {
		return
	}
	if err := os.Remove(store.filename); err != nil {
		panic(err)
	}
}
