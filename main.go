// TODO: Make everything configurable with spf13/viper or something.
package main

import (
	"log"
	"net/http"
	"text/template"
)

type Config struct {
	Url string
}

const bash = `
#!/usr/bin/env bash
echo "Please don't blindly curl stuff into your shell! You should not trust us!"
echo "We publish signed releases here:"
echo "{{.Url}}"
`

func handler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("bash").Parse(bash)
	if err != nil {
		log.Fatal(err)
	}

	cfg := Config{Url: "https://example.com"}

	if err := tmpl.Execute(w, cfg); err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
