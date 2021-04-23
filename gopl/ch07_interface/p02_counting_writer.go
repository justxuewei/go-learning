package ch07_interface

import "io"

type counting struct {
	w io.Writer
	l int64
}

func (c *counting) Write(b []byte) (int, error) {
	n, err := c.w.Write(b)
	c.l = int64(n)
	return n, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	counting := counting{w, 0}
	return &counting, &counting.l
}
