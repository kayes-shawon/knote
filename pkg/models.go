package pkg

import "fmt"

type Note struct {
	Id int64
	Title string
	Content string
}

func (n *Note) String() string  {
	return fmt.Sprintf("Title: %s, Content: %s", n.Title, n.Content)
}


