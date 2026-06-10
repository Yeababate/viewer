package main
import (
	"fmt"
	"net/http"
	"html/template"
	"log"

)

var tpl *template.Template

func main(){
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	tpl,_ = tpl.ParseGlob("templates/*html")

	http.HandleFunc("/", homeHandler) 
	fmt.Println("server is running")
	http.HandleFunc("/{id}", AllcarDetail)
	http.HandleFunc("GET /models/", allmodelsHandler)
	http.HandleFunc("GET /models/{id}", modelsHandler)
	http.HandleFunc("GET /manufacturers", allManufacturersHandler)
	http.HandleFunc("GET /manufacturers/{id}", manufacturersHandler)
	http.HandleFunc("GET /categories", allCategoriesHandler)
	http.HandleFunc("GET /categories/{id}", categoriesHandler)
	
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}
