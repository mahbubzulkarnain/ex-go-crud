package model

// Todo godoc.
type Todo struct {
	Id          int64  `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
}
