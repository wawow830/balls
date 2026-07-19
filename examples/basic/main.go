package main

import (
	"encoding/json"
	"fmt"

	"github.com/wawow830/balls/sequence"
)

func main() {
	generator := sequence.New(42)
	for range 3 {
		encoded, err := json.Marshal(generator.Next())
		if err != nil {
			panic(err)
		}
		fmt.Println(string(encoded))
	}
}
