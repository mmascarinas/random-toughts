package main

import (
	"log"
	"net/http"
)

const PORT = ":8080"

type api struct {
	addr string
}

func (s *api) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		switch r.URL.Path {
		case "/":
			w.Write([]byte("home page"))
			return
		case "/users":
			w.Write([]byte("users page"))
			return
		}
	default:
		w.Write([]byte("method not allowed"))
		return
	}
}

func main() {
	api := &api{addr: PORT}

	srv := &http.Server{
		Addr:    api.addr,
		Handler: api,
	}

	if err :=  srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
