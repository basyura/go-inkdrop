package inkdrop

type Tags []Tag

type Tag struct {
	Id        string `json:"_id"`
	Rev       string `json:"_rev"`
	Name      string `json:"name"`
	Color     string `json:"color"`
	Count     int    `json:"count"`
	CreatedAt int    `json:"createdAt"`
	UpdatedAt int    `json:"updatedAt"`
}
