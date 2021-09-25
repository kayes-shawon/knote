package models

import "fmt"

type Note struct {
	tableName struct{} `pg:"knote"`
	Id int64 `json:"id" pg:"id"`
	Title string `json:"title" pg:"title"`
	Content string `json:"content" pg:"content"`
	Tag string `json:"tag" pg:"tag"`
}

func (n *Note) String() string  {
	return fmt.Sprintf("Id: %d, Title: %s, Content: %s", n.Id, n.Title, n.Content)
}


