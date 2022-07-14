package main

import (
	"evesan/orders/cli/internal/retriever"
	"flag"
	"log"
	"os"
	"strings"
)

func main() {
	path := flag.String("file", "./sources/pendingmin.txt", "Input file")
	nRoutines := flag.Int("routines", 10, "Number of routines to use")
	flag.Parse()

	data, err := os.ReadFile(*path)
	if err != nil {
		log.Fatal(err)
	}

	pending := strings.Split(string(data), "\n")
	results := retriever.GetOrders(pending, *nRoutines)

	data = []byte(strings.Join(results, "\n"))

	err = os.WriteFile("res.csv", data, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

}
