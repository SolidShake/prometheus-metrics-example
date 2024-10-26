package main

import (
	"fmt"
	"net/http"
)

func healthHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Server is healthy")
}
