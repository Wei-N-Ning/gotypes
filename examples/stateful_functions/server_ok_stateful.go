package main

import "github.com/Wei-N-Ning/gotypes/pkg/iterator"

type OkStatefulServer struct{}

const MergeTicks = 4

func (_ OkStatefulServer) serve(store *Store, requests []Request) {
	iter := iterator.FromSlice(requests)
	requestsPerWorker := iterator.ChunkSlice(iter, NumWorkers)

	iterator.ParMapUnord(requestsPerWorker, func(reqs []Request) error {
		localState := map[string]int{}
		for idx, req := range reqs {
			if idx > 0 && idx%MergeTicks == 0 {
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
