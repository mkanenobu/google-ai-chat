package filehistory

import "os"

type Container struct {
	buffer []string
}

func (c *Container) At(n int) string {
	return c.buffer[n]
}

func (c *Container) Len() int {
	return len(c.buffer)
}

func (c *Container) Add(s string) {
	c.buffer = append(c.buffer, s)
}

func openFile() (*os.File, error) {
	return os.OpenFile("history.txt", os.O_RDWR|os.O_CREATE, 0666)

}

func New() *Container {
	return &Container{}
}
