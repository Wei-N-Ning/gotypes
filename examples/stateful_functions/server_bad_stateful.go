package main

import "github.com/Wei-N-Ning/gotypes/pkg/iterator"

type BadStatefulServer struct{}

func (_ BadStatefulServer) serve(store *Store, requests []Request) {
	iter := iterator.FromSlice(requests)
	workers := iterator.ChunkSlice(iter, NumWorkers)

	iterator.ParMapUnord(workers, func(reqs []Request) error {
		localState := map[string]int{}
		for _, req := range reqs {
			if req.action == "WRITE" {
				w := req.payload.(WriteAction)
				localState[w.word] += 1
			} else if req.action == "READ" {
				localVersion := localState[req.payload.(string)]
				publishedVersion := store.Read(req.payload.(string))
				_ = localVersion + publishedVersion
			}
		}
		return nil
	}).Count()
}
