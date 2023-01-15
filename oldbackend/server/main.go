package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/backslash-security-tests/golang-supermarket/oldbackend"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type OldbackendServer struct {
	postgresqlDB *sql.DB
	restClient   http.Client
}

func (srv *OldbackendServer) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var input oldbackend.CreateUserInput
	json.NewDecoder(r.Body).Decode(&input)

	output, err := srv.CreateUser(r.Context(), input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}

func (srv *OldbackendServer) CreateUser(ctx context.Context, input oldbackend.CreateUserInput) (
	output oldbackend.CreateUserOutput,
	err error,
) {
	username := input.Username
	password := input.Password

	srv.postgresqlDB.QueryRow(fmt.Sprintf("INSERT INTO users (username, password) VALUES ('%s', '%s')", username, password)).Scan()
	return output, nil
}

func main() {
	srv := &OldbackendServer{}

	// connect to the postgresql database using environment variables
	srv.postgresqlDB, err = sql.Open("postgres", "")
	if err != nil {
		log.Fatalf("Failed connecting to postgresql database: %v", err)
	}
	r := mux.NewRouter()
	r.Methods("POST").PathPrefix("/users").Handler(srv.CreateUserHandler)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":80", nil))
}
