package main

import (
	"fmt"
	"net/http"
	"io"
	"encoding/json"
	"strings"
	"strconv"
)

type Car struct {
	Id int `json:"id"`
	Name string `json:"name"`
	ManufacturerID int `json:"manufacturerId"` 
	CategoryID int `json:"categoryId"`
	Year int `json:"year"`
	Specifications Specifications `json:"specifications"`
	Image string `json:"image"`
}

type Specifications struct {
	Engine string `json:"engine"`
	Horsepower int `json:"horsepower"`
	Transmission string `json:"transmission"`
	Drivetrain string `json:"drivetrain"`
}

type Manufacturer struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Country string `json:"country"`
	FoundingYear int `json:"FoundingYear"`
}

type Category struct {
	Id int `json:"id"`
	Name string `json:"name"`
}

type HomeData struct {
	Cars any
	Manufacturers []Manufacturer
	Categories []Category
	Years []int
	PageNumbers []int
}

type CarDetail struct {
        Car          Car
        Manufacturer Manufacturer
        Category     Category
		Years []int
    }


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

func getSearchedCar(term string) ([]Car, error) {
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

func filter(cars []Car, r *http.Request) []Car{
    manufacturer := r.URL.Query().Get("manufacturer")
    category := r.URL.Query().Get("category")
    year := r.URL.Query().Get("year")

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
        filtered = append(filtered, car)
    }
	
   return filtered
}
