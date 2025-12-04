package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/abh1SHAKE/seekr/utils"
)

func main() {
	var inputFile string
	var query string

	flag.StringVar(&inputFile, "file", "", "Path to the text file you want to search")
	flag.StringVar(&query, "query", "", "Search query")
	flag.Parse()

	if inputFile == "" {
		log.Fatal("Error: --file is required (example: --file='./notes.text')")
	}

	if query == "" {
		log.Fatal("Error: --query is required (example: --query='wild cats')")
	}

	log.Printf("Searching for : %q", query)
	log.Printf("Reading file: %s", inputFile)

	start := time.Now()
	docs, err := utils.LoadDocs(inputFile)
	if err != nil {
		log.Fatalf("Failed to load documents: %v", err)
	}

	log.Printf("Loaded %d documents in %v", len(docs), time.Since(start))

	start = time.Now()
	index := make(utils.Index)
	index.Add(docs)

	log.Printf("Indexed %d documents in %v", len(docs), time.Since(start))

	start = time.Now()
	matchedIDs := index.Search(query)

	log.Printf("Found %d matching documents in %v", len(matchedIDs), time.Since(start))

	if len(matchedIDs) == 0 {
		fmt.Println("No results found.")
		return
	}

	fmt.Println("-- Search Results --")
	for _, id := range matchedIDs {
		doc := docs[id]
		fmt.Printf("[%d] %s\n", id, doc.Text)
	}
}