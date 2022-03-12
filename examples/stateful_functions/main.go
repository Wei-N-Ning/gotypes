package main

func main() {
	ws := GenerateDefaultWordSet(4)
	requests := RnadomRequests(ws)

	store := Spinup(5, 20)
	baseline := timeThis("baseline", 0, func() {
		BaselineServer{}.serve(store, requests)
	})

	store = Spinup(5, 20)
	timeThis("multi-workers", baseline, func() {
		WorkersServer{}.serve(store, requests)
	})

	store = Spinup(5, 20)
	timeThis("bad-stateful-function", baseline, func() {
		BadStatefulServer{}.serve(store, requests)
	})

	store.Teardown()
}
