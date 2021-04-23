package ch07_interface

import "io"

type LimitReader struct {
	ar io.Reader
	i int64
	l int64
}

func (r *LimitReader) Read(b []byte) (n int, err error) {
	tmpn, err := r.ar.Read(b)
	if err != nil {
		return 0, nil
	}

	maxn := r.l - r.i
	if int64(tmpn) > maxn {
		b = b[:maxn]
		return int(maxn), io.EOF
	}
	n = tmpn
	r.i += int64(n)
	return
}

func NewLimitReader(r io.Reader, n int) *LimitReader {
	return &LimitReader{
		ar: r,
		i: 0,
		l: int64(n),
	}
}
