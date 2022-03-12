package main

import "math/rand"

type Request struct {
	action  string
	payload interface{}
}

type WriteAction struct {
	word string
	x    int
}

func GenerateReadRequest(word string) Request {
	return Request{
		action:  "READ",
		payload: word,
	}
}

func GenerateWriteRequest(word string, x int) Request {
	return Request{
		action:  "WRITE",
		payload: WriteAction{word, x},
	}
}

// RnadomRequests creates N write-request (N = len(wordSet)), and 3x read-request.
// Their order is randomized
func RnadomRequests(wordSet []string) []Request {
	var rs []Request
	for _, word := range wordSet {
		rs = append(
			rs,
			GenerateWriteRequest(word, 1),
			GenerateReadRequest(word),
			GenerateReadRequest(word),
			GenerateReadRequest(word),
		)
	}
	rand.Shuffle(len(rs), func(i int, j int) {
		rs[i], rs[j] = rs[j], rs[i]
	})
	return rs
}
