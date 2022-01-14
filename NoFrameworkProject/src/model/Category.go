package model

type Category struct {
	Id   int
	Name string
}

func (c *Category) SetName(name string) {
	c.Name = "{" + name + "}"
}
