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

type CarDetail struct {
	Car          Car
	Manufacturer Manufacturer
	Category     Category
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

type Search struct {
	On bool
	Value string
	NotFound bool
}

type Filters struct {
	Manufacturer     string
	ManufacturerName string
	Category         string
	CategoryName     string
	Year             string
	On bool
	Err error
}
