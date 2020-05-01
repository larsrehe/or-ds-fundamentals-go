package main

import (
	"fmt"
	"regexp"
	"strings"

	"gopkg.in/jdkato/prose.v2"
)

func preProcessforNLP(text string) string {
	// Find all chars that are not alphabets
	reg := regexp.MustCompile("[^a-zA-Z ]+")

	// Replace those chars wtih spaces
	text = reg.ReplaceAllString(text, " ")

	// Lower case
	text = strings.ToLower(text)
	return text
}

func main() {
	str := "I am analyzing the tweets in    Golang and emailing 1231 2 to @myself&*!"
	str = preProcessforNLP(str)

	// Tokenization 1
	tokens := strings.Fields(str)
	for idx, token := range tokens {
		fmt.Println(idx, token, len(tokens))
	}

	// Create a new document with the default configuration:
	text := "Yesterday was the coldest day in New York City."
	doc, err := prose.NewDocument(text)
	if err != nil {
		panic(err)
	}

	// Tokenization 2
	for _, tok := range doc.Tokens() {
		fmt.Println(tok.Text, tok.Tag, tok.Label)
	}

	// Get Named Entities from document
	for _, ent := range doc.Entities() {
		fmt.Println(ent.Text, ent.Label)
	}

	// Get document's sentences
	for idx, sent := range doc.Sentences() {
		fmt.Println(idx, sent.Text)
	}
}
