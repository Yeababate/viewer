package main

import (
	"fmt"
	"net/http"
	"io"
	"encoding/json"
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

	carsData := string(body)
	return carsData, nil
}

func getCarModel(car string) (Car, error) {
	url := "http://localhost:3000/api/models/" + car
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

func getCarsManufacturer() (string, error) {
	url := "http://localhost:3000/api/manufacturers"
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

func manufacturerCars(car string) (Car, error) {
	url := "http://localhost:3000/api/manufacturers/" + car
	resp, errResp := http.Get(url)
	if errResp != nil {
		return Car{}, errResp
	}

	body, errRead := io.ReadAll(resp.Body)
	if errRead != nil {
		return Car{}, errRead
	}

	var manufacturer Car
	errJson := json.Unmarshal(body, &manufacturer)
	if errJson != nil {
		return Car{}, errJson
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

func categoriesCars(car string) (Car, error) {
	url := "http://localhost:3000/api/categories/" + car
	resp, errResp := http.Get(url)
	if errResp != nil {
		return Car{}, errResp
	}

	body, errRead := io.ReadAll(resp.Body)
	if errRead != nil {
		return Car{}, errRead
	}

	var category Car
	errJson := json.Unmarshal(body, &category)
	if errJson != nil {
		return Car{}, errJson
	}

	return category, nil
}

func main() {

	http.HandleFunc("GET /", homeHandler)
	http.HandleFunc("GET /models", modelsHandler)
	http.HandleFunc("GET /manufacturers", manufacturersHandler)
	http.HandleFunc("GET /categories", categoriesHandler)
	
	func homeHandler(w http.ResponseWriter, r *http.Request) {
		carsData, err := getCars()
		if err != nil {
			http.Error(w, "error message", http.StatusInternalServerError)
			return
		}
		w.Header().Set("content-type", "application/json")
		w.Write([]byte(carsData))
	}

	func modelsHandler(w http.ResponseWriter, r *http.Request) {
		carsData, err := getCars()
		if err != nil {
			http.Error(w, "error message", http.StatusInternalServerError)
			return
		}
		w.Header().Set("content-type", "application/json")
		w.Write([]byte(carsData))
	}
	
	func manufacturersHandler(w http.ResponseWriter, r *http.Request) {
		manufacturersData, err := getCarsManufacturer()
		if err != nil {
			http.Error(w, "error message", http.StatusInternalServerError)
			return
		}
		w.Header().Set("content-type", "application/json")
		w.Write([]byte(manufacturersData))
	}

	func categoriesHandler(w http.ResponseWriter, r *http.Request) {
		categoriesData, err := getCarsCategories()
		if err != nil {
			http.Error(w, "error message", http.StatusInternalServerError)
			return
		}
		w.Header().Set("content-type", "application/json")
		w.Write([]byte(categoriesData))
	}

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}

	// resp, err := getCarModel("3")
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(resp)
	// }
}
