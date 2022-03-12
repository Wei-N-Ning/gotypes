package main

type BaselineServer struct{}

func (_ BaselineServer) serve(store *Store, requests []Request) {
	for _, req := range requests {
		if req.action == "WRITE" {
			w := req.payload.(WriteAction)
			store.Upsert(w.word)
		} else if req.action == "READ" {
			store.Read(req.payload.(string))
		}
	}
}
