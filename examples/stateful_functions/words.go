package main

import "math/rand"

var Scale int = 1

var DefaultCountByWord = map[string]int{
	"there":   3 * Scale,
	"is":      4 * Scale,
	"a":       5 * Scale,
	"silence": 6 * Scale,
	"where":   7 * Scale,
	"hath":    8 * Scale,
	"been":    9 * Scale,
	"no":      2 * Scale,
	"sound":   2 * Scale,
}

func GenerateDefaultWordSet() []string {
	return generateWordSet(DefaultCountByWord)
}

func generateWordSet(countByWords map[string]int) []string {
	var ws []string
	for word, x := range countByWords {
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
