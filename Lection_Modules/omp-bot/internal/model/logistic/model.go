package logistic

import "fmt"

type Pack struct {
	Title string
}

func (c *Pack) String() string {
	return fmt.Sprintf("Title of product: %v", c.Title)
}
