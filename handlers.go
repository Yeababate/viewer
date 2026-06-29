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
		w.WriteHeader(http.StatusNotFound)
		tpl.ExecuteTemplate(w, "error404.html", nil)
		return
	}

	landingPageCars := h.AllCars

	landingPageCars = checkSearch(landingPageCars, h, w, r)

	h.Filters.Err = nil
	landingPageCars, h.Filters.Err  = checkFilter(landingPageCars, h, r)
	if h.Filters.Err != nil {
		log.Println(h.Filters.Err)
		w.WriteHeader(http.StatusInternalServerError)
		tpl.ExecuteTemplate(w, "error.html", nil)
		return
	}

	createPages(landingPageCars, h, w, r)

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
	DetailedCar CarDetail
}

func (h* CarDetailHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	id := r.PathValue("id")
	car, errCarModel, status := getCarModel(id)
	if status == 500{
		log.Println("failed to get car model")
		w.WriteHeader(http.StatusInternalServerError)
		tpl.ExecuteTemplate(w, "error.html", nil)
		return
	}
	
	
	if errCarModel != nil {
		log.Println("url not found")
		w.WriteHeader(http.StatusNotFound)
		tpl.ExecuteTemplate(w, "error404.html", nil)
		return
	}

	var errCarDetails error
	h.DetailedCar, errCarDetails = getCarDetails(car)
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
	SelectedCars []CarDetail
}

func (h *CompareCarHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var selectedCars []CarDetail
	if r.Method == "POST" {
		r.ParseForm()
		selectedIDs := r.Form["carIds"]
			for _, idStr := range selectedIDs {
				carModel, err,status := getCarModel(idStr)
				if status == 500{
					log.Println("failed to get car model")
					w.WriteHeader(http.StatusInternalServerError)
					tpl.ExecuteTemplate(w, "error.html", nil)
					return
				}
				if err != nil {
					continue
				}
				manufacturer, err:= getCarManufacturer(strconv.Itoa(carModel.ManufacturerID))
				if err != nil{
					w.WriteHeader(http.StatusInternalServerError)
    				tpl.ExecuteTemplate(w, "error.html", nil)
    				return
				}
				category, err := getCarCategory(strconv.Itoa(carModel.CategoryID))
				if err != nil{
					w.WriteHeader(http.StatusInternalServerError)
    				tpl.ExecuteTemplate(w, "error.html", nil)
    				return
				}
				selectedCars = append(selectedCars, CarDetail{
					Car:          carModel,
					Manufacturer: manufacturer,
					Category:     category,
				})
			}
	}
buf := bufferPool.Get().(*bytes.Buffer)
	buf.Reset()
	defer bufferPool.Put(buf)

	err := tpl.ExecuteTemplate(buf, "compare.html", selectedCars)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		tpl.ExecuteTemplate(w, "error.html", nil)
		
	} else {
		buf.WriteTo(w)
	}

}


