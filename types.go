	package main

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
	SearchOn bool
	SearchValue string
	FilterOn bool
	Filters struct{Manufacturer string; Category string; Year string}
}

type CarDetail struct {
        Car          Car
        Manufacturer Manufacturer
        Category     Category
		Years []int
    }

