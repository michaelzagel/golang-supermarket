package main

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/backslash-security-tests/golang-supermarket/oldbackend"
	"github.com/backslash-security-tests/golang-supermarket/oldfrontend"
	"github.com/gorilla/mux"
)

type OldfrontendServer struct {
	restClient http.Client
}

func (srv *OldfrontendServer) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var input oldfrontend.LoginInput
	json.NewDecoder(r.Body).Decode(&input)

	output, err := srv.Login(r.Context(), input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}

func (srv *OldfrontendServer) Login(ctx context.Context, input oldfrontend.LoginInput) (
	output oldfrontend.LoginOutput,
	err error,
) {
	username := input.Username
	password := input.Password

	call0Input, _ := json.Marshal(&oldbackend.CreateUserInput{
		Username: username,
		Password: password,
	})

	req0, err := http.NewRequest("POST", "http://oldbackend/users", bytes.NewReader(call0Input))
	if err != nil {
		return output, err
	}

	res0, err := srv.restClient.Do(req0)
	if err != nil {
		return output, err
	}

	defer res0.Body.Close()

	var call0 oldbackend.CreateUserOutput
	err = json.NewDecoder(res0.Body).Decode(&call0)
	if err != nil {
		return output, err
	}

	output.AccessToken = access_token
	return output, nil
}

func main() {
	srv := &OldfrontendServer{}
	r := mux.NewRouter()
	r.Methods("POST").PathPrefix("/login").Handler(srv.LoginHandler)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":80", nil))
}
