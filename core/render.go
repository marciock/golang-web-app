package core

import (
	"net/http"
	"html/template"
)
var templates *template.Template
func Render(page string, embed string, template string, res http.ResponseWriter) {
	tmp, _ :=templates.ParseFiles(page, embed)
	 
	
	 tmp.ExecuteTemplate(res, template, nil)
}
