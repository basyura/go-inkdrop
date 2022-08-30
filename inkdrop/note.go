package inkdrop

type Notes []Note

type Note struct {
	Id                string   `json:"_id"`
	Rev               string   `json:"_rev"`
	BookId            string   `json:"bookId"`
	Doctype           string   `json:"doctype"`
	Status            string   `json:"status"`
	Title             string   `json:"title"`
	Body              string   `json:"body"`
	Tags              []string `json:"tags"`
	NumOfTasks        int      `json:"numOfTasks"`
	NumOfCheckedTasks int      `json:"numOfCheckedTasks"`
	Pinned            bool     `json:"pinned"`
	Share             string   `json:"share"`
	CreatedAt         int      `json:"createdAt"`
	UpdatedAt         int      `json:"updatedAt"`
}
