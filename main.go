package main

import (
	"flag"
	"log"
	"time"

	utils "github.com/rajapremsai/text-search-engine-go/utils"
)

func main() {
	var dumpPath, query string
	flag.StringVar(&dumpPath, "p", "enwiki-latest-abstract1.xml.gz", "wiki dump data")
	// flag.StringVar(&dumpPath, "p", "P:\\Prem\\text_search_engine_using_go\\enwiki-latest-abstract1.gz.xml", "wiki dump data")
	flag.StringVar(&query, "q", "India, officially the Republic of India", "Search query")
	flag.Parse()
	log.Println("Full text search is in progress")

	start := time.Now()
	docs, err := utils.LoadDocuments(dumpPath)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Loaded %d documents in %v", len(docs), time.Since(start))

	start = time.Now()
	idx := make(utils.Index)
	idx.Add(docs)
	log.Printf("Indexed %d docs in %v", len(docs), time.Since(start))
	start = time.Now()
	matchedIDS := idx.Search(query)
	log.Printf("Search found %d in documents in %v", len(matchedIDS), time.Since(start))
	for _, id := range matchedIDS {
		doc := docs[id]
		log.Printf("%d\t%s\n", id, doc.Text)

	}
}
