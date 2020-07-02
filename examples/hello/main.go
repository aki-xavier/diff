package main

import (
	"encoding/json"
	"log"

	diff "github.com/aki-xavier/diff/src"
)

func main() {
	// changes := diff.CompareLines("hello world\nhello aki\n", "hello world\nhello aki\n")
	// for _, change := range changes {
	// 	data, err := json.Marshal(change)
	// 	if err == nil {
	// 		log.Println(string(data))
	// 	}
	// }

	changes := diff.CompareArray([]string{"hello world\n", "hello aki"}, []string{"hello world\n", "hello aki"})
	for _, change := range changes {
		data, err := json.Marshal(change)
		if err == nil {
			log.Println(string(data))
		}
	}
}
