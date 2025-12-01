package entity

type Post struct {
	Id       *int64  `json:"id,omitempty"`
	PersonId *int64  `json:"person_id,omitempty"`
	Content  *string `json:"content,omitempty"`
}
