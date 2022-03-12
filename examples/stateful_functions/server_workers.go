package main

import "github.com/Wei-N-Ning/gotypes/pkg/iterator"

type WorkersServer struct{}

const NumWorkers = 100

func (_ WorkersServer) serve(store *Store, requests []Request) {
	iter := iterator.FromSlice(requests)
	workers := iterator.ChunkSlice(iter, NumWorkers)

	iterator.ParMapUnord(workers, func(reqs []Request) error {
		for _, req := range reqs {
			if req.action == "WRITE" {
				w := req.payload.(WriteAction)
				store.Upsert(w.word)
			} else if req.action == "READ" {
				store.Read(req.payload.(string))
			}
		}
		return nil
	}).Count()
}
