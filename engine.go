package main

import (
	"fmt"
	"strings"
)

type Document struct {
	ID    int
	Title string
}

type SearchEngine struct {
	Docs     []Document
	Inverted map[string][]int
}

func NewSearchEngine() *SearchEngine {
	return &SearchEngine{
		Docs:     make([]Document, 0),
		Inverted: make(map[string][]int),
	}
}

func (s *SearchEngine) AddDoc(doc Document) {
	s.Docs = append(s.Docs, doc)
	for _, word := range strings.Fields(doc.Title) {
		word = strings.ToLower(word)
		s.Inverted[word] = append(s.Inverted[word], doc.ID)
	}
}

func levenshtein(a, b string) int {
	if len(a) == 0 {
		return len(b)
	}
	if len(b) == 0 {
		return len(a)
	}

	matrix := make([][]int, len(a)+1)
	for i := range matrix {
		matrix[i] = make([]int, len(b)+1)
	}

	for i := range matrix {
		matrix[i][0] = i
	}

	for j := range matrix[0] {
		matrix[0][j] = j
	}

	for i := 1; i <= len(a); i++ {
		for j := 1; j <= len(b); j++ {
			cost := 0
			if a[i-1] != b[j-1] {
				cost = 1
			}
			matrix[i][j] = min(matrix[i-1][j]+1, min(matrix[i][j-1]+1, matrix[i-1][j-1]+cost))
		}
	}

	return matrix[len(a)][len(b)]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (s *SearchEngine) Search(query string) []Document {
	query = strings.ToLower(query)
	ids, found := s.Inverted[query]
	if !found {
		return nil
	}

	var docs []Document
	for _, id := range ids {
		docs = append(docs, s.Docs[id])
	}
	return docs
}

func (s *SearchEngine) FuzzySearch(query string, maxDistance int) []Document {
	var docs []Document

	for _, doc := range s.Docs {
		for _, word := range strings.Fields(doc.Title) {
			if levenshtein(query, word) <= maxDistance {
				docs = append(docs, doc)
				break 
			}
		}
	}

	return docs
}

func main() {
	engine := NewSearchEngine()

	doc1 := Document{ID: 0, Title: "Hello World"}
	doc2 := Document{ID: 1, Title: "Go is fun"}
	doc3 := Document{ID: 2, Title: "How to use this"}

	engine.AddDoc(doc1)
	engine.AddDoc(doc2)
	engine.AddDoc(doc3)

	results := engine.FuzzySearch("Helol", 2)

	for _, doc := range results {
		fmt.Println(doc.Title)
	}
}