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
	// http.HandleFunc("/cars/models", carModels)
	// http.HandleFunc("/cars/manufacturers", carManufacturers)
	// http.HandleFunc("/cars/catagories", carCatagories)
	
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}
