package main

import (
	"bufio"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
)

const inputFilename string = `shakespeare-only.txt`
const defaultWordCount int = 25
const serverAddr string = ":9000"

var trie *Trie

func init() {
	f, err := os.Open(inputFilename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	trie = NewTrie()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		tokens := Parse(line)
		for _, token := range tokens {
			trie.Insert(token)
		}
	}
	err = scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("imported file %q, parsed %d words", inputFilename, len(trie.Words))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/autocomplete", autocompleteHandler)

	server := &http.Server{
		Addr:    serverAddr,
		Handler: mux,
	}
	log.Printf("started server on %s", serverAddr)
	server.ListenAndServe()
}

func autocompleteHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	terms, ok := query["term"]
	var term string
	if len(terms) > 0 {
		term = terms[0]
	}
	if !ok || len(term) < 2 {
		badTerm := struct {
			Error string `json:"error"`
		}{"missing term query param or too short"}
		js, err := json.Marshal(badTerm)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusPreconditionFailed)
		w.Write(js)
		return
	}
	countStrs, ok := query["count"]
	var count = defaultWordCount
	if ok && len(countStrs) > 0 {
		var err error
		count, err = strconv.Atoi(countStrs[0])
		if err != nil {
			count = defaultWordCount
		}
	}
	wordCounts := trie.MostFrequentWordsMatching(term, count)
	body := struct {
		Matches []WordCount `json:"matches"`
	}{wordCounts}
	js, err := json.Marshal(body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}
