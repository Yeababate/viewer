package main

import (
	"net/http"
	"strings"
	"strconv"
)
var CarsData []Car
var Manufacturers []Manufacturer
var Categories []Category
var Years []int

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w,r)
		return
	}
	var err error
	landingPageCars := CarsData
	if r.FormValue("q") != "" {
		landingPageCars, err = getSearchedCar(strings.ToLower(r.FormValue("q")))
		if err != nil {
			http.Error(w, "car not found", http.StatusInternalServerError)
			return
		}

	}
	filteredCar := filter(landingPageCars, r)
	data := HomeData{
		Cars: filteredCar,
		Manufacturers: Manufacturers,
		Categories: Categories,
		Years: Years,
	}

	tpl.ExecuteTemplate(w, "home.html", data)
}

func carDetailHandler(w http.ResponseWriter, r *http.Request){
	//make a []Car in case more than one car is found by search
	var carsFound []Car
	var err error

	//if r.FormValue("q") != "", something has been submitted in the search bar
	if r.FormValue("q") != "" {
		carsFound, err = getSearchedCar(strings.ToLower(r.FormValue("q")))
		if err != nil {
			http.Error(w, "car not found", http.StatusInternalServerError)
			return
		}
		
	}else if  r.URL.Query().Get("manufacturer")!= "" || r.URL.Query().Get("category")!= "" || r.URL.Query().Get("year")!= ""{
		carsFound = CarsData
	}else {
		id := r.PathValue("id")
		car, err := getCarModel(id)
		carsFound = append([]Car{}, car)
		if err != nil {
			http.Error(w, "error getting car", http.StatusInternalServerError)
			return
		}
	}

	filteredCar := filter(carsFound, r)
	cars := []CarDetail{}
	var detail CarDetail

	//get all details for each car in carsFound
	for i := range filteredCar {
		manufacturerDetails,err:= getCarManufacturer(strconv.Itoa(filteredCar[i].ManufacturerID))
		if err != nil {
			http.Error(w, "error getting manufacturer", http.StatusInternalServerError)
			return
		}
		categoryDetails,err:= getCarCategory(strconv.Itoa(filteredCar[i].CategoryID))
		if err != nil {
			http.Error(w, "error getting category", http.StatusInternalServerError)
			return
		}
	
   	 	detail = CarDetail{
       	 		Car: filteredCar[i],
        		Manufacturer: manufacturerDetails,
        		Category: categoryDetails,
				Years: Years,
    		}


    		cars = append(cars, detail)
    }
		data := HomeData{
			Cars: cars,
			Manufacturers: Manufacturers,
			Categories: Categories,
			Years: Years,
		}

	tpl.ExecuteTemplate(w, "car.html", data)
}

