package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	a := &api{addr: ":8080", mux: mux}

	mux.HandleFunc("GET /users", a.GetUserHandler)
	mux.HandleFunc("POST /users", a.CreateUserHandler)

	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			a.GetUserHandler(w, r)
		case http.MethodPost:
			a.CreateUserHandler(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	srv := &http.Server{
		Addr:    a.addr,
		Handler: a,
	}

	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
