package main

import (
	"net/http"
	"strings"
	"fmt"
	"strconv"
)

func getSearchedCars(term string) ([]Car, error) {
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
		return []Car{}, fmt.Errorf("car not found")
	} 
	
	return carsFound, nil

}

func checkSearch(landingPageCars []Car, w http.ResponseWriter, r *http.Request) ([]Car, bool) {
	var searchOn bool
	var err error
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
        	landingPageCars, err = getSearchedCars(SearchValue)
        	if err != nil {
		 	http.Error(w, "car not found", http.StatusInternalServerError)
                	 return landingPageCars, searchOn
        	}
		searchOn = true
	}

	return landingPageCars, searchOn
}

func getFilteredCars(cars []Car) []Car{
    /*manufacturer := r.URL.Query().Get("manufacturer")
    category := r.URL.Query().Get("category")
    year := r.URL.Query().Get("year")*/

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
	if  r.URL.Query().Get("manufacturer")!= "" || r.URL.Query().Get("category")!= "" || r.URL.Query().Get("year")!= ""{
		Filters.Manufacturer = r.URL.Query().Get("manufacturer")
		Filters.Category = r.URL.Query().Get("category")
		Filters.Year = r.URL.Query().Get("year")
		filterOn = true
	}
	if Filters.Manufacturer != "" || Filters.Category != "" || Filters.Year != "" {
		landingPageCars = getFilteredCars(landingPageCars)
		filterOn = true
	}

	return landingPageCars, filterOn


}


