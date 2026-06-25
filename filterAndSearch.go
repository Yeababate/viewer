package main

import (
	"net/http"
	"strings"
	//"fmt"
	"strconv"
)

func getSearchedCars(term string) ([]Car, bool) {
	var notFound bool
	//strings.Fields() splits strings around whitespace chars (incl mult spaces, tabs, newlines)
	searchTerms := strings.Fields(strings.ToLower(term))
	
	//[]Car in case more than one car is found by search
	var carsFound []Car
	var name string
	var year string
	if len(searchTerms) == 1 {
		for _,car := range CarsData {
			name = strings.ToLower(car.Name)
			year = strconv.Itoa(car.Year)
			if strings.Contains(name, term)||strings.Contains(year,term) {
				carsFound = append(carsFound, car)
			}
		}
	} else {
	for _,car := range CarsData {
		name = strings.ToLower(car.Name)
		year = strconv.Itoa(car.Year)
		matchCount := 0
		for _,term := range searchTerms {
			if strings.Contains(name, term) || strings.Contains(year,term) {
				matchCount += 1
			}
			
		}
		if matchCount == 2 {
			carsFound = append(carsFound, car)
		}
		matchCount = 0
	}
	}
	if len(carsFound)<1 {
		notFound = true
		return []Car{}, notFound
	} 
	
	return carsFound, notFound

}

func checkSearch(landingPageCars []Car, w http.ResponseWriter, r *http.Request) ([]Car, bool, bool) {
	var searchOn bool
	var notFound bool
	//check if Clear Search button clicked
	if r.FormValue("clearSearch") == "clear" {
		SearchValue = ""
	}

	//look for search results
	if r.FormValue("q") != "" {
		SearchValue = strings.ToLower(r.FormValue("q"))
		searchOn = true

	}
	if SearchValue != "" {
        	landingPageCars, notFound = getSearchedCars(SearchValue)
        	if notFound {
                	 notFound = true
        	}
		searchOn = true
	}

	return landingPageCars, searchOn, notFound
}

func getFilteredCars(cars []Car) []Car{

    var filtered []Car
    for _, car := range cars {
        if Filters.Manufacturer != "" && strconv.Itoa(car.ManufacturerID) != Filters.Manufacturer {
            continue
        }
        if Filters.Category != "" && strconv.Itoa(car.CategoryID) != Filters.Category {
            continue
        }
        if Filters.Year != "" && strconv.Itoa(car.Year) != Filters.Year {
            continue
        }
        filtered = append(filtered, car)
    }
	
   return filtered
}

func checkFilter(landingPageCars []Car, w http.ResponseWriter, r* http.Request) ([]Car, bool) {
	var filterOn bool

	//check if Clear Filter button clicked
	if r.FormValue("clearFilter") == "clear" {
		Filters.Manufacturer = ""
		Filters.Category = ""
		Filters.Year = ""
	}

		//look for filter results
		//IDEA make a map of manufacturers[ID]Name, categories[ID]Name to reduce API calls
	if  r.URL.Query().Get("manufacturer")!= "" {
		Filters.Manufacturer = r.URL.Query().Get("manufacturer")
		//HANDLE HTIS ERROR!!!
		manufacturerID, _ := strconv.Atoi(Filters.Manufacturer)
		Filters.ManufacturerName = ManufacturerMap[manufacturerID]
		filterOn = true
	}
	if r.URL.Query().Get("category") != "" {
		Filters.Category = r.URL.Query().Get("category")
		//HANDLE THIS ERROR!!!!
		categoryID,_ := strconv.Atoi(Filters.Category)
		Filters.CategoryName = CategoryMap[categoryID]
		filterOn = true
	}
	if r.URL.Query().Get("year") != "" {
		Filters.Year = r.URL.Query().Get("year")
		filterOn = true
	}

	if Filters.Manufacturer != "" || Filters.Category != "" || Filters.Year != "" {
		landingPageCars = getFilteredCars(landingPageCars)
		filterOn = true
	}

	return landingPageCars, filterOn


	}


