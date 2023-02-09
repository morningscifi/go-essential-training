package main

import (
	"fmt"
	"io"
	"os"
)

// Capper implements io.Writer and turns everything to uppercase
type Capper struct {
	wtr io.Writer
}

func (c *Capper) Write(p []byte) (n int, err error) {
	buffer := make([]byte, len(p))
	for k, v := range p {
		if v >= 'a' && v <= 'z' {
			v -= 32
		}
		buffer[k] = v
	}
	return c.wtr.Write(buffer)
}

func main() {
	c := &Capper{os.Stdout}
	fmt.Fprintln(c, "Hello there")
}
