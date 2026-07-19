package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/wawow830/balls/sequence"
)

func main() {
	count := flag.Int("count", 10, "number of JSON records to emit")
	seed := flag.Uint64("seed", 1, "deterministic generator seed")
	flag.Parse()

	if *count < 0 {
		fmt.Fprintln(os.Stderr, "count must be non-negative")
		os.Exit(2)
	}

	output := bufio.NewWriter(os.Stdout)
	encoder := json.NewEncoder(output)
	generator := sequence.New(*seed)
	for range *count {
		if err := encoder.Encode(generator.Next()); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
	if err := output.Flush(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
