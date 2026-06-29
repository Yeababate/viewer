package main

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"fmt"
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

func getCarModel(id string) (Car, error, int) {
	url := "http://localhost:3000/api/models/" + id
	resp, errResp := http.Get(url)
	if errResp != nil {
		return Car{}, errResp, 500
	}

	_, err := strconv.Atoi(id)
	if err != nil{
		return Car{}, fmt.Errorf("failed to get car model %s: status %d", id, resp.StatusCode), resp.StatusCode
	}
	if resp.StatusCode == http.StatusInternalServerError {
    	return Car{}, fmt.Errorf("failed to get car model %s: status %d", id, resp.StatusCode), resp.StatusCode
	}

	if resp.StatusCode == http.StatusNotFound {
    	return Car{}, fmt.Errorf("failed to get car model %s: status %d", id, resp.StatusCode), resp.StatusCode
	}	
	

	body, errRead := io.ReadAll(resp.Body)
	if errRead != nil {
		return Car{}, errRead,0
	}

	var model Car
	errJson := json.Unmarshal(body, &model)
	if errJson != nil {
		return Car{}, errJson,0
	}

	return model, nil,0
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

func getManufacturerMap(manufacturers []Manufacturer) map[int]string {
	manufacturerMap := make(map[int]string)
	for i:= range manufacturers {
		manufacturerMap[manufacturers[i].Id] = manufacturers[i].Name
	}

	return manufacturerMap
}

func getCategoryMap(categories []Category) map[int]string {
	categoryMap := make(map[int]string)
	for i:= range categories {
		categoryMap[categories[i].Id] = categories[i].Name
	}

	return categoryMap
}

func getCarDetails(car Car) (CarDetail, error) {
	var detail CarDetail

	manufacturerDetails, err := getCarManufacturer(strconv.Itoa(car.ManufacturerID))
	if err != nil {
		return detail, err
	}
	categoryDetails, err := getCarCategory(strconv.Itoa(car.CategoryID))
	if err != nil {
		return detail, err
	}

	detail = CarDetail{
		Car:          car,
		Manufacturer: manufacturerDetails,
		Category:     categoryDetails,
	}

	return detail, nil
}
