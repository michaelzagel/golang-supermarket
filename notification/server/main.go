package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/backslash-security-tests/golang-supermarket/notification"
	"github.com/gorilla/mux"
)

type NotificationServer struct {
	restClient http.Client
}

func (srv *NotificationServer) NotifyWindowHandler(w http.ResponseWriter, r *http.Request) {
	var input notification.NotifyWindowInput
	json.NewDecoder(r.Body).Decode(&input)

	output, err := srv.NotifyWindow(r.Context(), input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}

func (srv *NotificationServer) NotifyWindow(ctx context.Context, input notification.NotifyWindowInput) (
	output notification.NotifyWindowOutput,
	err error,
) {
	order_id := input.OrderId
	window_from := input.WindowFrom
	window_to := input.WindowTo

	output.Success = success
	return output, nil
}

func main() {
	srv := &NotificationServer{}
	r := mux.NewRouter()
	r.Methods("POST").PathPrefix("/notify_window").Handler(srv.NotifyWindowHandler)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":80", nil))
}
