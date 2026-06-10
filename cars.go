package main

import (
	//"fmt"
	"net/http"
	"io"
	"encoding/json"
	"fmt"
)

type Car struct {
	Id int `json: "id"`
	Name string `json: "name"`
	ManufacturerID int `json: manufacturerId"` 
	CategoryID int `json: "categoryId"`
	Year int `json: "year"`
	Specifications Specifications `json: "specifications"`
	Image string `json: "image"`
}

type Specifications struct {
	Engine string `json: "engine"`
	Horsepower int `json: "horsepower"`
	Transmission string `json: "transmission"`
	Drivetrain string `json: "drivetrain"`
}

type Manufacturers struct {
	Id int `json: "id"`
	Name string `json: "name"`
	Country string `json: "country"`
	FoundingYear int `json: FoundingYear`
}

type Catagories struct {
	Id int `json: "id"`
	Name string `json: "name"`
}

type CarDetail struct {
        Car          Car
        Manufacturer Manufacturers
        Category     Catagories
    }

func getCars() (string, error) {
	url := "http://localhost:3000/api/models"
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	body, errRead := io.ReadAll(resp.Body)
	if errRead != nil {
		return "", errRead
	}

	carsData := string(body)  // why chage to string because we're changing it back 
	return carsData, nil
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

func getCarsManufacturer() ([]byte, error) {
	url := "http://localhost:3000/api/manufacturers/"
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	body, errRead := io.ReadAll(resp.Body)
	if errRead != nil {
		return nil, errRead
	}

	return body, nil
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

func getCarsCategories() (string, error) {
	url := "http://localhost:3000/api/categories"
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	body, errRead := io.ReadAll(resp.Body)
	if errRead != nil {
		return "", errRead
	}

	carsData := string(body)
	return carsData, nil
}

func categoriesCars(car string) (Catagories, error) {
	url := "http://localhost:3000/api/categories/" + car
	resp, errResp := http.Get(url)
	if errResp != nil {
		return Catagories{}, errResp
	}

	body, errRead := io.ReadAll(resp.Body)
	if errRead != nil {
		return Catagories{}, errRead
	}

	var category Catagories
	errJson := json.Unmarshal(body, &category)
	if errJson != nil {
		return Catagories{}, errJson
	}

	return category, nil
}

type FrontPage struct {
	ModelName string 
	ManuYear int
	Image string
}

func landingPage() ([]FrontPage, error){
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
	var landing []FrontPage
	for _, car := range model{
		landing = append(landing, FrontPage{
			ModelName: car.Name,
			ManuYear: car.Year,
			Image: car.Image,
		})
	}
	

	return landing, nil
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
		carsData, err := landingPage()
		if err != nil {
			http.Error(w, "error message", http.StatusInternalServerError)
			return
		}
		/*jsonData,err := json.Marshal(carsData)
		if err != nil {
			http.Error(w, "error message", http.StatusInternalServerError)
			return
		}
		w.Header().Set("content-type", "application/json")
		w.Write(jsonData)*/

		tpl.ExecuteTemplate(w, "home.html", carsData)
}

func allmodelsHandler(w http.ResponseWriter, r *http.Request) {
		carsData, err := getCars()
		if err != nil {
			http.Error(w, "error message", http.StatusInternalServerError)
			return
		}
		w.Header().Set("content-type", "application/json")
		w.Write([]byte(carsData))
	}

func modelsHandler(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		carsData, err := getCarModel(id)
		if err != nil {
			http.Error(w, "error message", http.StatusInternalServerError)
			return
		}
		jsonData, err := json.Marshal(carsData)
		if err != nil {
			http.Error(w, "error message", http.StatusInternalServerError)
			return
		}
		w.Header().Set("content-type", "application/json")
		w.Write(jsonData)
	}
	
func allManufacturersHandler(w http.ResponseWriter, r *http.Request) {
		manufacturersData, err := getCarsManufacturer()
		if err != nil {
			http.Error(w, "error message", http.StatusInternalServerError)
			return
		}
		w.Header().Set("content-type", "application/json")
		w.Write(manufacturersData)
	}

func manufacturersHandler(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		manufacturersData, err := manufacturerCars(id)
		if err != nil {
			http.Error(w, "error message", http.StatusInternalServerError)
			return
		}
		jsonData, err := json.Marshal(manufacturersData)
		if err != nil {
			http.Error(w, "error message", http.StatusInternalServerError)
			return
		}
		w.Header().Set("content-type", "application/json")
		w.Write(jsonData)
	}

func allCategoriesHandler(w http.ResponseWriter, r *http.Request) {
		categoriesData, err := getCarsCategories()
		if err != nil {
			http.Error(w, "error message", http.StatusInternalServerError)
			return
		}
		w.Header().Set("content-type", "application/json")
		w.Write([]byte(categoriesData))
	}
func categoriesHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	categoriesData, err := categoriesCars(id)
	if err != nil {
			http.Error(w, "error message", http.StatusInternalServerError)
			return
		}
	jsonData,err := json.Marshal(categoriesData)
	if err != nil {
		http.Error(w, "error message", http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(jsonData)
}

func AllcarDetail(w http.ResponseWriter, r *http.Request){
	id := r.PathValue("id")
	car ,err := getCarModel(id)
	if err != nil {
		http.Error(w, "error getting car", http.StatusInternalServerError)
		return
	}
	manufacturer,err:= manufacturerCars(fmt.Sprintf("%d", car.ManufacturerID))
	if err != nil {
		http.Error(w, "error getting manufacturer", http.StatusInternalServerError)
		return
	}
	category,err:= categoriesCars(fmt.Sprintf("%d", car.CategoryID))
	if err != nil {
		http.Error(w, "error getting category", http.StatusInternalServerError)
		return
	}
	
    detail := CarDetail{
        Car:          car,
        Manufacturer: manufacturer,
        Category:     category,
    }
	tpl.ExecuteTemplate(w, "car.html", detail) 
}


