package model

type Category struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (c *Category) SetName(name string) {
	c.Name = "{" + name + "}"
}
