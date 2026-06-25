package main

import (
	"net/http"
	"strconv"
	"strings"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
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
		Cars:          carsOnPage,
		Manufacturers: Manufacturers,
		Categories:    Categories,
		Years:         Years,
		PageNumbers:   pageNumbers,
		SearchOn:      searchOn,
		SearchValue:   SearchValue,
		FilterOn:      filterOn,
		Filters:       Filters,
	}

	tpl.ExecuteTemplate(w, "home.html", data)
}

func carDetailHandler(w http.ResponseWriter, r *http.Request) {
	//var carsFound []Car
	var searchOn bool
	var filterOn bool
	//var err error

	carsFound := CarsData

	carsFound, searchOn = checkSearch(carsFound, w, r)

	carsFound, filterOn = checkFilter(carsFound, w, r)

	if !searchOn && !filterOn {
		id := strings.TrimPrefix(r.URL.Path, "/car/")
		car, err := getCarModel(id)
		carsFound = append([]Car{}, car)
		if err != nil {
			http.Error(w, "error getting car", http.StatusInternalServerError)
		}
	}

	carsOnPage, pageNumbers := createPages(carsFound, w, r)

	//HANDLE THIS ERROR!!!
	cars, _ := getCarDetails(carsOnPage, w, r)

	data := HomeData{
		Cars:          cars,
		Manufacturers: Manufacturers,
		Categories:    Categories,
		Years:         Years,
		PageNumbers:   pageNumbers,
		SearchOn:      searchOn,
		SearchValue:   SearchValue,
		FilterOn:      filterOn,
		Filters:       Filters,
	}

	tpl.ExecuteTemplate(w, "car.html", data)
}
func compareCarHandler(w http.ResponseWriter, r *http.Request) {
	carDetails, _ := getCarDetails(CarsData, w, r)

	var selectedCars []CarDetail
	compareError := ""
	if r.Method == "POST" {
		r.ParseForm()
		selectedIDs := r.Form["carIds"]
		if len(selectedIDs) != 2 {
			compareError = "Please select exactly 2 cars to compare."
		} else {
			for _, idStr := range selectedIDs {
				carModel, err := getCarModel(idStr)
				if err != nil {
					compareError = "Error loading selected cars, Please try again."
					selectedCars = nil
					break
				}
				manufacturer, _ := getCarManufacturer(strconv.Itoa(carModel.ManufacturerID))
				category, _ := getCarCategory(strconv.Itoa(carModel.CategoryID))
				selectedCars = append(selectedCars, CarDetail{
					Car:          carModel,
					Manufacturer: manufacturer,
					Category:     category,
				})
			}
		}
	}

	data := CompareData{
		Cars:         carDetails,
		SelectedCars: selectedCars,
		CompareError: compareError,
	}

	tpl.ExecuteTemplate(w, "compare.html", data)
}
