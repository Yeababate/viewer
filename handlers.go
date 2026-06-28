package main

import (
	"net/http"
	"strconv"
	"sync"
	"bytes"
	"log"
)

var bufferPool = sync.Pool {
	New: func() interface{} {
		return new(bytes.Buffer) //fresh buffer on demand
	},
}

type HomeHandler struct {
	//Data *APIData
	Cars []Car
	Manufacturers []Manufacturer
	Categories []Category
	Years []int
	ManufacturerIDMap map[int]string
	CategoryIDMap map[int]string
}

func (h *HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	landingPageCars := h.Cars
	var searchOn bool
	var notFound bool
	var filterOn bool
	var errFilter error

	landingPageCars, searchOn, notFound = checkSearch(landingPageCars, w, r)

	landingPageCars, filterOn, errFilter  = checkFilter(landingPageCars, h, w, r)
	if errFilter != nil {
		log.Println(errFilter)
		w.WriteHeader(http.StatusInternalServerError)
		tpl.ExecuteTemplate(w, "error.html", nil)
		return
	}

	carsOnPage, pageNumbers, currentPage, numberOfPages/*, notBeg, notEnd */:= createPages(landingPageCars, w, r)


	data := HomeData{
		Cars:          carsOnPage,
		Manufacturers: h.Manufacturers,
		Categories: h.Categories,
		Years: h.Years,
		Pagination: Pagination{PageNumbers: pageNumbers, CurrentPage: currentPage, LastPage: numberOfPages, /*NotBeg: notBeg, NotEnd: notEnd*/},
		SearchOn: searchOn,
		SearchValue: SearchValue,
		NotFound: notFound,
		FilterOn: filterOn,
		Filters: Filters,
	}

	//get a buffer 
	buf := bufferPool.Get().(*bytes.Buffer)
	//reset buffer, in case it was used before
	buf.Reset()
	//defer returning buffer to pool
	defer bufferPool.Put(buf)

	//write to buffer instead of w
	err := tpl.ExecuteTemplate(buf, "home.html", data)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		tpl.ExecuteTemplate(w, "error.html", nil)
	} else {

		//if no error, write buffer to response writer
		buf.WriteTo(w)
	}
}

type CarDetailHandler struct {
	Cars []Car
	Manufacturers []Manufacturer
	Categories []Category
	Years []int
	ManufacturerIDMap map[int]string
	CategoryIDMap map[int]string
}

func (h* CarDetailHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	
	//var err error

	cars, errCarDetails := getCarDetails(h.Cars, h.Years, w, r)
	if errCarDetails != nil {
		log.Println(errCarDetails)
		w.WriteHeader(http.StatusInternalServerError)
		tpl.ExecuteTemplate(w, "error.html", nil)
		return
	}

		data := HomeData{
			Cars: cars,
			Manufacturers: h.Manufacturers,
			Categories: h.Categories,
			Years: h.Years,
		}
	
	buf := bufferPool.Get().(*bytes.Buffer)
	buf.Reset()
	defer bufferPool.Put(buf)

	err := tpl.ExecuteTemplate(buf, "car.html", data)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		tpl.ExecuteTemplate(w, "error.html", nil)
	} else {
		buf.WriteTo(w)
	}
}

type CompareCarHandler struct {
	Cars []Car
	Years []int
}

func (h *CompareCarHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	carDetails, _ := getCarDetails(h.Cars, h.Years, w, r)

	var selectedCars []CarDetail
	compareError := ""
	if r.Method == "POST" {
		r.ParseForm()
		selectedIDs := r.Form["carIds"]
		if len(selectedIDs) != 2 {
			compareError = "Please select exactly 2 cars to compare."
		} else {
			for _, idStr := range selectedIDs {
				carModel, err := getCarModel(idStr)
				if err != nil {
					compareError = "Error loading selected cars, Please try again."
					selectedCars = nil
					break
				}
				manufacturer, _ := getCarManufacturer(strconv.Itoa(carModel.ManufacturerID))
				category, _ := getCarCategory(strconv.Itoa(carModel.CategoryID))
				selectedCars = append(selectedCars, CarDetail{
					Car:          carModel,
					Manufacturer: manufacturer,
					Category:     category,
				})
			}
		}
	}

	data := CompareData{
		Cars:         carDetails,
		SelectedCars: selectedCars,
		CompareError: compareError,
	}

	tpl.ExecuteTemplate(w, "compare.html", data)
}
