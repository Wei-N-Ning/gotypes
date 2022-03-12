package main

import "math/rand"

var DefaultCountByWord = map[string]int{
	"there":   3,
	"is":      4,
	"a":       5,
	"silence": 6,
	"where":   7,
	"hath":    8,
	"been":    9,
	"no":      2,
	"sound":   2,
}

func GenerateDefaultWordSet(scale int) []string {
	return generateWordSet(DefaultCountByWord, scale)
}

func generateWordSet(countByWords map[string]int, scale int) []string {
	var ws []string
	for word, x := range countByWords {
		x *= scale
		xs := make([]string, x)
		for i := 0; i < x; i++ {
			xs[i] = word
		}
		ws = append(ws, xs...)
	}
	rand.Shuffle(len(ws), func(i int, j int) {
		ws[i], ws[j] = ws[j], ws[i]
	})
	return ws
}
