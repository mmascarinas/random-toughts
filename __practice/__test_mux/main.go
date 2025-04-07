package main

import "net/http"

const PORT = ":8080"

func main() {
	api := &api{addr: PORT}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /users", api.getUsersHandler)
	mux.HandleFunc("POST /", api.createUsersHandler)

	srv := &http.Server{
		Addr:    api.addr,
		Handler: mux,
	}

	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
