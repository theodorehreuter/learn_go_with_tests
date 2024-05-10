package main

import (
	"log"
	"net/http"

	ps "github.com/theodorehreuter/learn_go_with_tests/server"
)

func main() {
	handler := http.HandlerFunc(ps.PlayerServer)
	log.Fatal(http.ListenAndServe(":5000", handler))
}
