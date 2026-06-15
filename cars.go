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

type Manufacturers struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Country string `json:"country"`
	FoundingYear int `json:"FoundingYear"`
}

type Categories struct {
	Id int `json:"id"`
	Name string `json:"name"`
}

type CarDetail struct {
        Car          Car
        Manufacturer Manufacturers
        Category     Categories
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

func manufacturerCars(car string) (Manufacturers, error) {
	url := "http://localhost:3000/api/manufacturers/" + car
	resp, errResp := http.Get(url)
	if errResp != nil {
		return Manufacturers{}, errResp
	}

	body, errRead := io.ReadAll(resp.Body)
	if errRead != nil {
		return Manufacturers{}, errRead
	}

	var manufacturer Manufacturers
	errJson := json.Unmarshal(body, &manufacturer)
	if errJson != nil {
		return Manufacturers{}, errJson
	}

	return manufacturer, nil
}


func categoriesCars(car string) (Categories, error) {
	url := "http://localhost:3000/api/categories/" + car
	resp, errResp := http.Get(url)
	if errResp != nil {
		return Categories{}, errResp
	}

	body, errRead := io.ReadAll(resp.Body)
	if errRead != nil {
		return Categories{}, errRead
	}

	var category Categories
	errJson := json.Unmarshal(body, &category)
	if errJson != nil {
		return Categories{}, errJson
	}

	return category, nil
}

/*type FrontPage struct {
	Id int
	ModelName string 
	ManuYear int
	Image string
}*/

func landingPage() ([]Car, error){
	url := "http://localhost:3000/api/models"
	resp, errResp := http.Get(url)
	if errResp != nil {
		return nil, errResp
	}

	body, errRead := io.ReadAll(resp.Body)
	if errRead != nil {
		return nil, errRead
	}

	var model []Car
	errJson := json.Unmarshal(body, &model)
	if errJson != nil {
		return nil, errJson
	}
/*	var landing []FrontPage
	for _, car := range model{
		landing = append(landing, FrontPage{
			Id: car.Id,
			ModelName: car.Name,
			ManuYear: car.Year,
			Image: car.Image,
		})
	}*/

	return model, nil
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

type HomeData struct {
	Cars []FrontPage
	Manufacturers []Manufacturers
	Categories []Categories
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
=======
	var err error
	CarsData, err = landingPage()
	landingPageCars := CarsData
	if err != nil {
		http.Error(w, "error message", http.StatusInternalServerError)
		return
	}

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

	tpl.ExecuteTemplate(w, "home.html", landingPageCars)
}

func AllcarDetail(w http.ResponseWriter, r *http.Request){
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
		manufacturer,err:= manufacturerCars(fmt.Sprintf("%d", carsFound[i].ManufacturerID))
		if err != nil {
			http.Error(w, "error getting manufacturer", http.StatusInternalServerError)
			return
		}
		category,err:= categoriesCars(fmt.Sprintf("%d", carsFound[i].CategoryID))
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

    var filtered []FrontPage
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
        filtered = append(filtered, FrontPage{
            Id: car.Id,
            ModelName: car.Name,
            ManuYear: car.Year,
            Image: car.Image,
        })
    }
	
    manufacturersBody, _ := http.Get("http://localhost:3000/api/manufacturers/")
    var manufacturers []Manufacturers
    body, _ := io.ReadAll(manufacturersBody.Body)
    json.Unmarshal(body, &manufacturers)

    categoriesBody, _ := http.Get("http://localhost:3000/api/categories")
    var categories []Categories
    body, _ = io.ReadAll(categoriesBody.Body)
    json.Unmarshal(body, &categories)

    data := HomeData{
        Cars: filtered,
        Manufacturers: manufacturers,
        Categories: categories,
    }

    tpl.ExecuteTemplate(w, "home.html", data)
}
