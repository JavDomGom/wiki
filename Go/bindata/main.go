package main

import (
	"html/template"
	"log"
	"net/http"
)

var tmpl *template.Template

func init() {
	data, err := Asset("files/index.html")
	if err != nil {
		log.Fatal(err)
	}
	tmpl = template.Must(template.New("tmpl").Parse(string(data)))
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		tmpl.Execute(w, map[string]string{"Name": "Javier"})
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
