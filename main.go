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
	//req, errReq := http.NewRequest("GET", url, nil)
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

func main() {
	resp, err := getCarModel("3")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resp)
	}
}
