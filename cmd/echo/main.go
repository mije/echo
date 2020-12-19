package main

import (
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gofrs/uuid"
)

const version = "0.0.2"

func main() {
	host := os.Getenv("HOST")
	if host == "" {
		host = "0.0.0.0"
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Header().Set("Version", version)
		w.Header().Set("RequestId", uuid.Must(uuid.NewV4()).String())
		w.WriteHeader(200)
		w.Write(b)
	})
	http.ListenAndServe(host+":"+port, nil)
}
