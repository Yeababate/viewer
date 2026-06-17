package main

import (
	"net/http"
	"strings"
	"strconv"
)
var CarsData []Car
var Manufacturers []Manufacturer
var Categories []Category

func homeHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	landingPageCars := CarsData
	if r.FormValue("q") != "" {
		landingPageCars, err = getSearchedCar(strings.ToLower(r.FormValue("q")))
		if err != nil {
			http.Error(w, "car not found", http.StatusInternalServerError)
			return
		}

	} else {
		if r.URL.Path != "/" {
			http.NotFound(w,r)
			return
		}
	}

	data := HomeData{
		Cars: landingPageCars,
		Manufacturers: Manufacturers,
		Categories: Categories,
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
		
	} else {
		id := r.PathValue("id")
		car, err := getCarModel(id)
		carsFound = append([]Car{}, car)
		if err != nil {
			http.Error(w, "error getting car", http.StatusInternalServerError)
			return
		}
	}

	
	cars := []CarDetail{}
	var detail CarDetail

	//get all details for each car in carsFound
	for i := range carsFound {
		manufacturer,err:= getCarManufacturer(strconv.Itoa(carsFound[i].ManufacturerID))
		if err != nil {
			http.Error(w, "error getting manufacturer", http.StatusInternalServerError)
			return
		}
		category,err:= getCarCategory(strconv.Itoa(carsFound[i].CategoryID))
		if err != nil {
			http.Error(w, "error getting category", http.StatusInternalServerError)
			return
		}
	
   	 	detail = CarDetail{
       	 		Car:          carsFound[i],
        		Manufacturer: manufacturer,
        		Category:     category,
    		}


    		cars = append(cars, detail)
    	}

	tpl.ExecuteTemplate(w, "car.html", cars)
}

func filterHandler(w http.ResponseWriter, r *http.Request) {
    manufacturer := r.URL.Query().Get("manufacturer")
    category := r.URL.Query().Get("category")
    year := r.URL.Query().Get("year")

    cars, err := getCars()
    if err != nil {
        http.Error(w, "error getting cars", http.StatusInternalServerError)
        return
    }

    var filtered []Car
    for _, car := range cars {
        if manufacturer != "" && strconv.Itoa(car.ManufacturerID) != manufacturer {
            continue
        }
        if category != "" && strconv.Itoa(car.CategoryID) != category {
            continue
        }
        if year != "" && strconv.Itoa(car.Year) != year {
            continue
        }
        filtered = append(filtered, Car{
            Id: car.Id,
            Name: car.Name,
            Year: car.Year,
            Image: car.Image,
        })
    }
	
    data := HomeData{
        Cars: filtered,
        Manufacturers: Manufacturers,
        Categories: Categories,
    }

    tpl.ExecuteTemplate(w, "home.html", data)
}
