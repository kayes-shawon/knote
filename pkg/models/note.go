package models

import "fmt"

type Note struct {
	Id int64 `json:"id"`
	Title string `json:"title"`
	Content string `json:"content"`
	Tag string `json:"tag"`
}

func (n *Note) String() string  {
	return fmt.Sprintf("Title: %s, Content: %s", n.Title, n.Content)
}


