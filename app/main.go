package main

import (
	"fmt"
	"net/http"

	"proxy"

	"google.golang.org/appengine"
)

func main() {
	http.HandleFunc("/get_index", func(w http.ResponseWriter, r *http.Request) {
		ctx := appengine.NewContext(r)
		if indexBody, err := proxy.GetIndex(ctx); err == nil {
			fmt.Fprintf(w, indexBody)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	appengine.Main()
}
