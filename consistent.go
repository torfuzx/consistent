package consistent

import "stathat.com/c/consistent"

type Circle struct {
	*consistent.Consistent
}

func New() *Circle {
	return &Circle{
		consistent.New(),
	}
}

func (c *Circle) Add(elt string) {
	c.Consistent.Add(elt)
}

func (c *Circle) Remove(elt string) {
	c.Consistent.Remove(elt)
}

func (c *Circle) Get(name string) (string, error) {
	elt, err := c.Consistent.Get(name)
	if err != nil {
		return "", err
	}

	return elt, nil
}

func (c *Circle) Set(elts []string) {
	c.Consistent.Set(elts)
}

func (c *Circle) Members() []string {
	return c.Consistent.Members()
}
