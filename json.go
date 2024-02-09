package main

import (
	"log"
	"encoding/json"
	"net/http"
)

func resonsewitherror(w http.ResponseWriter, code int, msg string){
	if code>499{
		log.Printf("Responding with error: %v)", msg)
	}
	type errorResponse struct{
		Error string `json:"error"`
	}
	respondwithJSON(w, code, errorResponse{Error: msg})
}

func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil{
		log.Printf("Failed to marshal the response payload: %v", payload)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}