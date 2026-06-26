package main

import (
	"net/http"
	"sync"
	"bytes"
	"log"
)


var bufferPool = sync.Pool {
	New: func() interface{} {
		return new(bytes.Buffer) //fresh buffer on demand
	},
}


func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w,r)
		return
	}

	landingPageCars := CarsData
	var searchOn bool
	var notFound bool
	var filterOn bool
	var errFilter error

	landingPageCars, searchOn, notFound = checkSearch(landingPageCars, w, r)

	landingPageCars, filterOn, errFilter  = checkFilter(landingPageCars, w, r)
	if errFilter != nil {
		log.Println(errFilter)
		w.WriteHeader(http.StatusInternalServerError)
		tpl.ExecuteTemplate(w, "error.html", nil)
		return
	}

	carsOnPage, pageNumbers := createPages(landingPageCars, w, r)

	data := HomeData{
		Cars: carsOnPage,
		Manufacturers: Manufacturers,
		Categories: Categories,
		Years: Years,
		PageNumbers: pageNumbers,
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

func carDetailHandler(w http.ResponseWriter, r *http.Request){
	//var carsFound []Car
	var searchOn bool
	var notFound bool
	var filterOn bool
	var errFilter error
	//var err error

	carsFound := CarsData

	carsFound, searchOn, notFound = checkSearch(carsFound, w, r)

	carsFound, filterOn, errFilter = checkFilter(carsFound, w, r)
	if errFilter != nil {
		log.Println(errFilter)
		w.WriteHeader(http.StatusInternalServerError)
		tpl.ExecuteTemplate(w, "error.html", nil)
		return
	}

	if !searchOn && !filterOn {
		id := r.PathValue("id")
		car, errCarModel := getCarModel(id)
		carsFound = append([]Car{}, car)
		if errCarModel != nil {
			log.Println(errCarModel)
			w.WriteHeader(http.StatusInternalServerError)
			tpl.ExecuteTemplate(w, "error.html", nil)
			return
		}
	}

	carsOnPage, pageNumbers := createPages(carsFound, w, r)

	cars, errCarDetails := getCarDetails(carsOnPage, w, r)
	if errCarDetails != nil {
		log.Println(errCarDetails)
		w.WriteHeader(http.StatusInternalServerError)
		tpl.ExecuteTemplate(w, "error.html", nil)
		return
	}

		data := HomeData{
			Cars: cars,
			Manufacturers: Manufacturers,
			Categories: Categories,
			Years: Years,
			PageNumbers: pageNumbers,
			SearchOn: searchOn,
			SearchValue: SearchValue,
			NotFound: notFound,
			FilterOn: filterOn,
			Filters: Filters,
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
