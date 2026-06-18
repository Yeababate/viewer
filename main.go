package main
import (
	"fmt"
	"net/http"
	"html/template"
	"log"
	"sync"
)

var tpl *template.Template

func main(){
	var err error
	//use WaitGroup to wait for all goroutines fetching API data to complete before starting server
        var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		CarsData, err = getCars()
		if err != nil {
			log.Fatal(err)
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
        	Manufacturers, err = getManufacturers()
		if err != nil {
			log.Fatal(err)
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
        	Categories, err = getCategories()
		if err != nil {
			log.Fatal(err)
		}
	}()
	//wait until above functions complete before continuing
	wg.Wait()

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	tpl,_ = tpl.ParseGlob("templates/*html")
	Years = getYears(CarsData)
	fmt.Println("server is running")
	http.HandleFunc("/", homeHandler) 
	http.HandleFunc("GET /car/{id}", carDetailHandler)
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
