package ch01_tutorial

import (
	"fmt"
	"io"
	"net/http"
)

type URLFetcher struct {
}

func (f *URLFetcher) fetch(url string, w io.Writer) (n int64) {
	resp, err := http.Get(url)
	if err != nil {
		panic(fmt.Sprintf("fetch: %v", err))
	}
	n, err = io.Copy(w, resp.Body)
	if err != nil {
		panic(fmt.Sprintf("fetch: copying to writer occurs an error, %v", err))
	}
	return
}
