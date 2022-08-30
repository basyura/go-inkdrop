# go-inkdrop


## sample

```go
package main

import (
	"github.com/basyura/go-inkdrop/client"
	. "github.com/basyura/go-inkdrop/inkdrop"
)

func main() {
  // https://docs.inkdrop.app/manual/accessing-the-local-database#get-docid
  note, err := client.Get[Note](client.Parameter{
    "docid": "note:cacpQeu6G",
  })

  // https://docs.inkdrop.app/manual/accessing-the-local-database#get-notes
  notes, _ := client.Get[Notes](client.Parameter{
    "sort":    "title",
    "keyword": "inkdrop",
  })

  // https://docs.inkdrop.app/manual/accessing-the-local-database#get-books
  books, _ := client.Get[Books](client.Parameter{})

  // https://docs.inkdrop.app/manual/accessing-the-local-database#get-tags
  tags, _ := client.Get[Tags](client.Parameter{})
}
```
