package main
import (
	"fmt"
	"net/http"
	"html/template"
	"log"

)

var tpl *template.Template

func main(){
	var err error
        CarsData, err = getCars()
        Manufacturers, err = getManufacturers()
        Categories, err = getCategories()
        if err != nil {
                log.Fatal(err)
        }

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	tpl,_ = tpl.ParseGlob("templates/*html")

	fmt.Println("server is running")

	http.HandleFunc("/", homeHandler) 
	http.HandleFunc("GET /car/{id}", carDetailHandler)
	http.HandleFunc("GET /filter", filterHandler)
	
	// http.HandleFunc("GET /models/", allmodelsHandler)
	// http.HandleFunc("GET /models/{id}", modelsHandler)
	// http.HandleFunc("GET /manufacturers", allManufacturersHandler)
	// http.HandleFunc("GET /manufacturers/{id}", manufacturersHandler)
	// http.HandleFunc("GET /categories", allCategoriesHandler)
	// http.HandleFunc("GET /categories/{id}", categoriesHandler)
	
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}
