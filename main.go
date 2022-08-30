package main

import (
	"encoding/json"
	"fmt"

	"github.com/basyura/go-inkdrop/client"
	. "github.com/basyura/go-inkdrop/inkdrop"
)

func main() {
	fmt.Println("********************")
	fmt.Println("*    go-inkdrop    *")
	fmt.Println("********************")

	note, err := client.Get[Note](client.Parameter{
		"docid": "note:cacpQeu6G",
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	printJson(note)

	notes, _ := client.Get[Notes](client.Parameter{
		"sort":    "title",
		"keyword": "inkdrop",
	})
	printJson(len(notes))
	printJson(notes[0])

	books, _ := client.Get[Books](client.Parameter{})
	printJson(books[0])

	tags, _ := client.Get[Tags](client.Parameter{})
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
