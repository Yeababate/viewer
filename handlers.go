package main

import (
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w,r)
		return
	}
	//var err error


	landingPageCars := CarsData
	var searchOn bool
	var notFound bool
	var filterOn bool

	landingPageCars, searchOn, notFound = checkSearch(landingPageCars, w, r)

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
		NotFound: notFound,
		FilterOn: filterOn,
		Filters: Filters,
	}

	tpl.ExecuteTemplate(w, "home.html", data)
}

func carDetailHandler(w http.ResponseWriter, r *http.Request){
	//var carsFound []Car
	var searchOn bool
	var notFound bool
	var filterOn bool
	//var err error

	carsFound := CarsData

	carsFound, searchOn, notFound = checkSearch(carsFound, w, r)

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

	//HANDLE THIS ERROR!!!
	cars,_ := getCarDetails(carsOnPage, w, r)

		data := HomeData{
			Cars: cars,
			Manufacturers: Manufacturers,
			Categories: Categories,
			Years: Years,
			PageNumbers: pageNumbers,
			SearchOn: searchOn,
			SearchValue: SearchValue,
			NotFound: notFound,
			FilterOn: filterOn,
			Filters: Filters,
		}

	tpl.ExecuteTemplate(w, "car.html", data)
}

