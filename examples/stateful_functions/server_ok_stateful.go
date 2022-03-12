package main

import "github.com/Wei-N-Ning/gotypes/pkg/iterator"

type OkStatefulServer struct{}

func (_ OkStatefulServer) serve(store *Store, requests []Request) {
	iter := iterator.FromSlice(requests)
	workers := iterator.ChunkSlice(iter, NumWorkers)

	iterator.ParMapUnord(workers, func(reqs []Request) error {
		localState := map[string]int{}
		for idx, req := range reqs {
			if idx > 0 && idx%30 == 0 {
				store.BatchUpsert(localState)
			}
			if req.action == "WRITE" {
				w := req.payload.(WriteAction)
				localState[w.word] += 1
			} else if req.action == "READ" {
				localVersion := localState[req.payload.(string)]
				publishedVersion := store.Read(req.payload.(string))
				_ = localVersion + publishedVersion
			}
		}
		store.BatchUpsert(localState)
		return nil
	}).Count()
}
