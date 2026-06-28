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
	AllCars []Car
	CarsOnPage []Car
	Manufacturers []Manufacturer
	Categories []Category
	Years []int
	ManufacturerIDMap map[int]string
	CategoryIDMap map[int]string
	Search Search
	Filters Filters
	Pagination Pagination
}

func (h *HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	landingPageCars := h.AllCars

	landingPageCars, h.Search.On, h.Search.NotFound = checkSearch(landingPageCars, h, w, r)

	landingPageCars, h.Filters.On, h.Filters.Err  = checkFilter(landingPageCars, h, w, r)
	if h.Filters.Err != nil {
		log.Println(h.Filters.Err)
		w.WriteHeader(http.StatusInternalServerError)
		tpl.ExecuteTemplate(w, "error.html", nil)
		return
	}

	h.CarsOnPage, h.Pagination.PageNumbers, h.Pagination.CurrentPage, h.Pagination.LastPage = createPages(landingPageCars, w, r)

	//get a buffer 
	buf := bufferPool.Get().(*bytes.Buffer)
	//reset buffer, in case it was used before
	buf.Reset()
	//defer returning buffer to pool
	defer bufferPool.Put(buf)

	//write to buffer instead of w
	err := tpl.ExecuteTemplate(buf, "home.html", h)

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
	DetailedCar []CarDetail
}

func (h* CarDetailHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	id := r.PathValue("id")
	car, errCarModel := getCarModel(id)
	carsFound := append([]Car{}, car)
	if errCarModel != nil {
		log.Println(errCarModel)
		w.WriteHeader(http.StatusInternalServerError)
		tpl.ExecuteTemplate(w, "error.html", nil)
		return
	}

	var errCarDetails error
	h.DetailedCar, errCarDetails = getCarDetails(carsFound, w, r)
	if errCarDetails != nil {
		log.Println(errCarDetails)
		w.WriteHeader(http.StatusInternalServerError)
		tpl.ExecuteTemplate(w, "error.html", nil)
		return
	}

	buf := bufferPool.Get().(*bytes.Buffer)
	buf.Reset()
	defer bufferPool.Put(buf)

	err := tpl.ExecuteTemplate(buf, "car.html", h)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		tpl.ExecuteTemplate(w, "error.html", nil)
	} else {
		buf.WriteTo(w)
	}
}

type CompareCarHandler struct {
	AllCars []Car
	DetailedCars []CarDetail
	Years []int
}

func (h *CompareCarHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	carDetails, _ := getCarDetails(h.AllCars, w, r)

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
