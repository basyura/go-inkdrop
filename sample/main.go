package main

import (
	"encoding/json"
	"fmt"

	"github.com/basyura/go-inkdrop"
)

func main() {
	fmt.Println("********************")
	fmt.Println("*    go-inkdrop    *")
	fmt.Println("********************")

	note, err := inkdrop.Get[inkdrop.Note](inkdrop.Parameter{
		"docid": "note:cacpQeu6G",
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	printJson(note)

	notes, _ := inkdrop.Get[inkdrop.Notes](inkdrop.Parameter{
		"sort":    "title",
		"keyword": "inkdrop",
	})
	printJson(len(notes))
	printJson(notes[0])

	books, _ := inkdrop.Get[inkdrop.Books](inkdrop.Parameter{})
	printJson(books[0])

	tags, _ := inkdrop.Get[inkdrop.Tags](inkdrop.Parameter{})
	printJson(tags[0])
}

func printJson(v any) {
	fmt.Println("-------------------------")
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
}
