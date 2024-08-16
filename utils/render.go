package utils

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func Render(w http.ResponseWriter, templateName string, data any) {
	partials := []string{
		"views/base.layout.gohtml",
	}

	var templateSlice []string
	templateSlice = append(templateSlice, fmt.Sprintf("views/%s", templateName))

	templateSlice = append(templateSlice, partials...)

	tmpl, err := template.ParseFiles(templateSlice...)
	if err != nil {
		log.Printf("could not parse template: %v\n", err)
		http.Error(w, "could not parse template", http.StatusInternalServerError)
	}
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
