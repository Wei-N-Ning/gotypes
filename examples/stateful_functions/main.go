package main

func main() {
	ws := GenerateDefaultWordSet()
	requests := RnadomRequests(ws)
	Scale = 1

	store := Spinup(5, 20)
	baseline := timeThis("baseline", 0, func() {
		BaselineServer{}.serve(store, requests)
	})

	store = Spinup(5, 20)
	timeThis("multi-workers", baseline, func() {
		WorkersServer{}.serve(store, requests)
	})

	store.Teardown()
}
