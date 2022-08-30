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
  note, err := inkdrop.Get[inkdrop.Note](inkdrop.Parameter{
    "docid": "note:cacpQeu6G",
  })

  // https://docs.inkdrop.app/manual/accessing-the-local-database#get-notes
  notes, _ := inkdrop.Get[inkdrop.Notes](inkdrop.Parameter{
    "sort":    "title",
    "keyword": "inkdrop",
  })

  // https://docs.inkdrop.app/manual/accessing-the-local-database#get-books
  books, _ := inkdrop.Get[inkdrop.Books](inkdrop.Parameter{})

  // https://docs.inkdrop.app/manual/accessing-the-local-database#get-tags
  tags, _ := inkdrop.Get[inkdrop.Tags](inkdrop.Parameter{})
}
```
