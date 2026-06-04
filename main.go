package main
import (
	"fmt"
	"net/http"

)

func main(){
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		fmt.Println("server is running")
	})
	// http.HandleFunc("/cars/models", carModels)
	// http.HandleFunc("/cars/manufacturers", carManufacturers)
	// http.HandleFunc("/cars/catagories", carCatagories)
	
	http.ListenAndServe(":8080", nil)

}