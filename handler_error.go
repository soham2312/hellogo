package main

import (
	"net/http"
)

func handleeErr(w http.ResponseWriter, r *http.Request) {
	resonsewitherror(w, 400, "Something went wrong!")
}