package main

type Car struct {
	Id             int            `json:"id"`
	Name           string         `json:"name"`
	ManufacturerID int            `json:"manufacturerId"`
	CategoryID     int            `json:"categoryId"`
	Year           int            `json:"year"`
	Specifications Specifications `json:"specifications"`
	Image          string         `json:"image"`
}

type Specifications struct {
	Engine       string `json:"engine"`
	Horsepower   int    `json:"horsepower"`
	Transmission string `json:"transmission"`
	Drivetrain   string `json:"drivetrain"`
}

type Manufacturer struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Country      string `json:"country"`
	FoundingYear int    `json:"FoundingYear"`
}

type Category struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type HomeData struct {
	Cars          any
	Manufacturers []Manufacturer
	Categories []Category
	Years []int
	Pagination Pagination
	SearchOn bool
	SearchValue string
	NotFound bool
	FilterOn bool
	Filters struct{Manufacturer string; ManufacturerName string; Category string; CategoryName string; Year string}
}

type CarDetail struct {
	Car          Car
	Manufacturer Manufacturer
	Category     Category
	Years        []int
}

type CompareData struct {
	Cars         []CarDetail
	SelectedCars []CarDetail
	CompareError string
}

type Pagination struct {
	PageNumbers []int
	CurrentPage int
	LastPage int
}

var CarsData []Car
var Manufacturers []Manufacturer
var Categories []Category
var Years []int
var ManufacturerMap map[int]string
var CategoryMap map[int]string

var Filters struct {
	Manufacturer     string
	ManufacturerName string
	Category         string
	CategoryName     string
	Year             string
}

var SearchValue string
