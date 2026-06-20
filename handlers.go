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

var Filters struct {
	Manufacturer string
	Category string
	Year string
}
var SearchValue string

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w,r)
		return
	}
	var err error
	var searchOn bool
	var filterOn bool
	landingPageCars := CarsData

	//check if Clear Search button clicked
	if r.FormValue("clearSearch") == "clear" {
		SearchValue = ""
	}

	//check if Clear Filter button clicked
	if r.FormValue("clearFilter") == "clear" {
		Filters.Manufacturer = ""
		Filters.Category = ""
		Filters.Year = ""
	}

	//look for search results
	if r.FormValue("q") != "" {
		SearchValue = strings.ToLower(r.FormValue("q"))
		searchOn = true

	}
	if SearchValue != "" {
        	landingPageCars, err = getSearchedCar(SearchValue)
        	if err != nil {
		 	http.Error(w, "car not found", http.StatusInternalServerError)
                	 return
        	}
		searchOn = true
	}

	//CHANGE THIS
	filteredCar := landingPageCars

	//look for filter results
	if  r.URL.Query().Get("manufacturer")!= "" || r.URL.Query().Get("category")!= "" || r.URL.Query().Get("year")!= ""{
		Filters.Manufacturer = r.URL.Query().Get("manufacturer")
		Filters.Category = r.URL.Query().Get("category")
		Filters.Year = r.URL.Query().Get("year")
		filterOn = true
	}
	if Filters.Manufacturer != "" || Filters.Category != "" || Filters.Year != "" {
		filteredCar = filter(landingPageCars)
		filterOn = true
	}

	//break cars up into pages
	var numPages int
	var pageNumbers []int
	var index int
	resultsPerPage := 5

	carsOnPage := filteredCar

	if len(filteredCar) > resultsPerPage {
		
		if len(filteredCar)%resultsPerPage == 0 {
			numPages = len(filteredCar)/resultsPerPage
		} else {
			numPages = len(filteredCar)/resultsPerPage + 1
		}

		for num := range numPages {
			//start pageNumbers at 1 instead of 0
			pageNumbers = append(pageNumbers, num+1)
		}
		if r.URL.Query().Get("page") != "" {
			pageNum, err := strconv.Atoi(r.URL.Query().Get("page"))
			if err != nil {
				http.Error(w, "error getting page number", http.StatusInternalServerError)
			}
			index = (pageNum * resultsPerPage) - resultsPerPage
		} else {
			index = 0
		}
	
		carsOnPage = filteredCar[index:index+resultsPerPage]
	}

	data := HomeData{
		Cars: carsOnPage,
		Manufacturers: Manufacturers,
		Categories: Categories,
		Years: Years,
		PageNumbers: pageNumbers,
		SearchOn: searchOn,
		FilterOn: filterOn,
	}

	tpl.ExecuteTemplate(w, "home.html", data)
}

func carDetailHandler(w http.ResponseWriter, r *http.Request){
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

	filteredCar := filter(carsFound)
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

