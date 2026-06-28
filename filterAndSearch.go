package main

import (
	"net/http"
	"strconv"
	"strings"
	)

func getSearchedCars(term string, cars []Car) ([]Car, bool) {
	var notFound bool
	//strings.Fields() splits strings around whitespace chars (incl mult spaces, tabs, newlines)
	searchTerms := strings.Fields(strings.ToLower(term))

	//[]Car in case more than one car is found by search
	var carsFound []Car
	var name string
	var year string
	if len(searchTerms) == 1 {
		for _, car := range cars {
			name = strings.ToLower(car.Name)
			year = strconv.Itoa(car.Year)
			if strings.Contains(name, term) || strings.Contains(year, term) {
				carsFound = append(carsFound, car)
			}
		}
	} else {
		for _, car := range cars {
			name = strings.ToLower(car.Name)
			year = strconv.Itoa(car.Year)
			matchCount := 0
			for _, term := range searchTerms {
				if strings.Contains(name, term) || strings.Contains(year, term) {
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

func checkSearch(landingPageCars []Car, h *HomeHandler, w http.ResponseWriter, r *http.Request) ([]Car) {
	h.Search.On = false
	h.Search.NotFound = false

	//check if Clear Search button clicked
	if r.FormValue("clearSearch") == "clear" {
		h.Search.Value = ""
	}

	//look for search results
	if r.FormValue("q") != "" {
		h.Search.Value = strings.ToLower(r.FormValue("q"))
		h.Search.On = true

	}
	if h.Search.Value != "" {
			var notFound bool
        	landingPageCars, notFound = getSearchedCars(h.Search.Value, landingPageCars)
        	if notFound {
                	 h.Search.NotFound = true
        	}
		h.Search.On = true
	}

	return landingPageCars
}

func getFilteredCars(cars []Car, h *HomeHandler) []Car {

	var filtered []Car
	for _, car := range cars {
		if h.Filters.Manufacturer != "" && strconv.Itoa(car.ManufacturerID) != h.Filters.Manufacturer {
			continue
		}
		if h.Filters.Category != "" && strconv.Itoa(car.CategoryID) != h.Filters.Category {
			continue
		}
		if h.Filters.Year != "" && strconv.Itoa(car.Year) != h.Filters.Year {
			continue
		}
		filtered = append(filtered, car)
	}

	return filtered
}

func checkFilter(landingPageCars []Car, h *HomeHandler, w http.ResponseWriter, r* http.Request) ([]Car, error) {
	h.Filters.On = false

	//check if Clear Filter button clicked
	if r.FormValue("clearFilter") == "clear" {
		h.Filters.Manufacturer = ""
		h.Filters.Category = ""
		h.Filters.Year = ""
	}

	//look for filter results
	if  r.URL.Query().Get("manufacturer")!= "" {
		h.Filters.Manufacturer = r.URL.Query().Get("manufacturer")
		manufacturerID, err := strconv.Atoi(h.Filters.Manufacturer)
		if err != nil {
			return landingPageCars, err
		}
		h.Filters.ManufacturerName = h.ManufacturerIDMap[manufacturerID]
		h.Filters.On = true
	}
	if r.URL.Query().Get("category") != "" {
		h.Filters.Category = r.URL.Query().Get("category")
		categoryID, err := strconv.Atoi(h.Filters.Category)
		if err != nil {
			return landingPageCars, err
		}
		h.Filters.CategoryName = h.CategoryIDMap[categoryID]
		h.Filters.On = true
	}
	if r.URL.Query().Get("year") != "" {
		h.Filters.Year = r.URL.Query().Get("year")
		h.Filters.On = true
	}

	if h.Filters.Manufacturer != "" || h.Filters.Category != "" || h.Filters.Year != "" {
		landingPageCars = getFilteredCars(landingPageCars, h)
		h.Filters.On = true
	}

	return landingPageCars, nil
}
