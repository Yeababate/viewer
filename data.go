package main

import (
	"net/http"
	"io"
	"encoding/json"
	"strconv"
)


func getCars() ([]Car, error) {
	url := "http://localhost:3000/api/models"
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	body, errRead := io.ReadAll(resp.Body)
	if errRead != nil {
		return nil, errRead
	}
	var cars []Car
    err = json.Unmarshal(body, &cars)
    if err != nil {
        return nil, err
    }

	return cars, nil
}

func getManufacturers() ([]Manufacturer, error) {
        url := "http://localhost:3000/api/manufacturers/"
        resp, err := http.Get(url)
        if err != nil {
                return nil, err
        }

        body, errRead := io.ReadAll(resp.Body)
        if errRead != nil {
                return nil, errRead
        }

        var manufacturers []Manufacturer
        err = json.Unmarshal(body, &manufacturers)
        if err != nil {
                return nil, err
        }

        return manufacturers, nil
}

func getCategories() ([]Category, error) {
        url := "http://localhost:3000/api/categories"
        resp, err := http.Get(url)
        if err != nil {
                return nil, err
        }

        body, errRead := io.ReadAll(resp.Body)
        if errRead != nil {
                return nil, errRead
        }

        var categories []Category
        err = json.Unmarshal(body, &categories)
        if err != nil {
                return nil, err
        }

        return categories, nil
}

func getCarModel(id string) (Car, error) {
	url := "http://localhost:3000/api/models/" + id
	resp, errResp := http.Get(url)
	if errResp != nil {
		return Car{}, errResp
	}

	body, errRead := io.ReadAll(resp.Body)
	if errRead != nil {
		return Car{}, errRead
	}

	var model Car
	errJson := json.Unmarshal(body, &model)
	if errJson != nil {
		return Car{}, errJson
	}

	return model, nil
}

func getCarManufacturer(car string) (Manufacturer, error) {
	url := "http://localhost:3000/api/manufacturers/" + car
	resp, errResp := http.Get(url)
	if errResp != nil {
		return Manufacturer{}, errResp
	}

	body, errRead := io.ReadAll(resp.Body)
	if errRead != nil {
		return Manufacturer{}, errRead
	}

	var manufacturer Manufacturer
	errJson := json.Unmarshal(body, &manufacturer)
	if errJson != nil {
		return Manufacturer{}, errJson
	}

	return manufacturer, nil
}


func getCarCategory(car string) (Category, error) {
	url := "http://localhost:3000/api/categories/" + car
	resp, errResp := http.Get(url)
	if errResp != nil {
		return Category{}, errResp
	}

	body, errRead := io.ReadAll(resp.Body)
	if errRead != nil {
		return Category{}, errRead
	}

	var category Category
	errJson := json.Unmarshal(body, &category)
	if errJson != nil {
		return Category{}, errJson
	}

	return category, nil
}

func getYears(cars []Car) []int {
    yearMap := map[int]bool{}
    var years []int
    for _, car := range cars {
        if !yearMap[car.Year] {
            yearMap[car.Year] = true
            years = append(years, car.Year)
        }
    }
    return years
}

func createPages(landingPageCars []Car, w http.ResponseWriter, r* http.Request) ([]Car, []int) {
	//break cars up into pages
	var numPages int
	var pageNumbers []int
	var index int
	resultsPerPage := 10

	carsOnPage := landingPageCars

	if len(landingPageCars) > resultsPerPage {
	
	if len(landingPageCars)%resultsPerPage == 0 {
		numPages = len(landingPageCars)/resultsPerPage
	} else {
		numPages = len(landingPageCars)/resultsPerPage + 1
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

	carsOnPage = landingPageCars[index:index+resultsPerPage]
}

	return carsOnPage, pageNumbers
}
