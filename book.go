package inkdrop

type Books []Book

type Book struct {
	Id           string `json:"_id"`
	Rev          string `json:"_rev"`
	ParentBookId string `json:"parentBookId"`
	Name         string `json:"name"`
	Count        int    `json:"count"`
	CreatedAt    int    `json:"createdAt"`
	UpdatedAt    int    `json:"updatedAt"`
}
