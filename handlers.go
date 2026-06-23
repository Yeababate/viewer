package main

import (
	"net/http"
	"strconv"
)
var CarsData []Car
var Manufacturers []Manufacturer
var Categories []Category
var Years []int
var ManufacturerMap map[int]string
var CategoryMap map[int]string

var Filters struct {
	Manufacturer string
	ManufacturerName string
	Category string
	CategoryName string
	Year string
}

var SearchValue string

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w,r)
		return
	}
	//var err error


	landingPageCars := CarsData
	var searchOn bool
	var filterOn bool

	landingPageCars, searchOn = checkSearch(landingPageCars, w, r)

	landingPageCars, filterOn = checkFilter(landingPageCars, w, r)

	carsOnPage, pageNumbers := createPages(landingPageCars, w, r)

	data := HomeData{
		Cars: carsOnPage,
		Manufacturers: Manufacturers,
		Categories: Categories,
		Years: Years,
		PageNumbers: pageNumbers,
		SearchOn: searchOn,
		SearchValue: SearchValue,
		FilterOn: filterOn,
		Filters: Filters,
	}

	tpl.ExecuteTemplate(w, "home.html", data)
}

func carDetailHandler(w http.ResponseWriter, r *http.Request){
	//var carsFound []Car
	var searchOn bool
	var filterOn bool
	//var err error

	carsFound := CarsData

	carsFound, searchOn = checkSearch(carsFound, w, r)

	carsFound, filterOn = checkFilter(carsFound, w, r)

	if !searchOn && !filterOn {
		id := r.PathValue("id")
		car, err := getCarModel(id)
		carsFound = append([]Car{}, car)
		if err != nil {
			http.Error(w, "error getting car", http.StatusInternalServerError)
		}
	}

	carsOnPage, pageNumbers := createPages(carsFound, w, r)

	//filteredCar := getFilteredCars(carsFound)
	cars := []CarDetail{}
	var detail CarDetail

	//get all details for each car in carsFound
	for i := range carsOnPage {
		manufacturerDetails,err:= getCarManufacturer(strconv.Itoa(carsOnPage[i].ManufacturerID))
		if err != nil {
			http.Error(w, "error getting manufacturer", http.StatusInternalServerError)
			return
		}
		categoryDetails,err:= getCarCategory(strconv.Itoa(carsOnPage[i].CategoryID))
		if err != nil {
			http.Error(w, "error getting category", http.StatusInternalServerError)
			return
		}
	
   	 	detail = CarDetail{
       	 		Car: carsOnPage[i],
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
			PageNumbers: pageNumbers,
			SearchOn: searchOn,
			SearchValue: SearchValue,
			FilterOn: filterOn,
			Filters: Filters,
		}

	tpl.ExecuteTemplate(w, "car.html", data)
}

