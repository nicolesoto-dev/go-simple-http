package main

import (
	"encoding/json"
	"net/http"
)

type api struct {
	addr string
	mux  *http.ServeMux
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"first_name"`
}

var users = []User{}

func (a *api) getUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (a *api) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var payload User
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	u := User{
		ID:   len(users) + 1,
		Name: payload.Name,
	}
	users = append(users, u)

	w.WriteHeader(http.StatusCreated)
}

func (a *api) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" && r.Method == http.MethodGet {
		w.Write([]byte("Hello World"))
		return
	}

	a.mux.ServeHTTP(w, r)
}

func main() {
	a := &api{addr: ":8080"}
	mux := http.NewServeMux()

	mux.HandleFunc("GET /users", a.getUserHandler)
	mux.HandleFunc("POST /users", a.CreateUserHandler)

	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			a.getUserHandler(w, r)
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
