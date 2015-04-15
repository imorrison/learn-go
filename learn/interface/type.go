package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type webPage struct {
	url  string
	body []byte
	err  error
}

func (w *webPage) get() {
	resp, err := http.Get(w.url)

	if err != nil {
		w.err = err
		return
	}

	defer resp.Body.Close()

	w.body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		w.err = err
	}
}

func (w *webPage) isOk() bool {
	return w.err == nil
}

func main() {
	// Could have also used new(webPage)
	w := &webPage{url: "http://localhost:8000"}
	w.get()
	fmt.Printf("URL: %s Error %s Length %d\n", w.url, w.err, len(w.body))
}
