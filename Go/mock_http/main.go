package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/files", getJSON)
	http.ListenAndServe(":3000", nil)
}

func getJSON(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./file.json")
}
