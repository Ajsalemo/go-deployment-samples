package main

import (
	"net/http"
	"log"
	"encoding/json"

	"goji.io"
	"goji.io/pat"
)

func hello(w http.ResponseWriter, r *http.Request) {
	res := make(map[string]string)
	res["message"] = "go-deployment-samples-goji"
	jsonResp, err := json.Marshal(res)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}

	w.Write(jsonResp)
}

func main() {
	mux := goji.NewMux()
	mux.HandleFunc(pat.Get("/"), hello)
	print("Listening on 0.0.0.0:8080")
	http.ListenAndServe(":8080", mux)
}