package main

import (
	"net/http"
	"strconv"
)

func createPages(landingPageCars []Car, w http.ResponseWriter, r* http.Request) ([]Car, []int) {
	//break cars up into pages
	var numPages int
	var pageNumbers []int
	var pageNum int
	var index int
	resultsPerPage := 10
	overflow := 0
	var err error

	carsOnPage := landingPageCars

	if len(landingPageCars) > resultsPerPage {
	
	if len(landingPageCars)%resultsPerPage == 0 {
		numPages = len(landingPageCars)/resultsPerPage
	} else {
		numPages = len(landingPageCars)/resultsPerPage + 1
		overflow = len(landingPageCars)%resultsPerPage
	}

	for num := range numPages {
		//start pageNumbers at 1 instead of 0
		pageNumbers = append(pageNumbers, num+1)
	}
	if r.URL.Query().Get("page") != "" {
		pageNum, err = strconv.Atoi(r.URL.Query().Get("page"))
		if err != nil {
			http.Error(w, "error getting page number", http.StatusInternalServerError)
		}
		index = (pageNum * resultsPerPage) - resultsPerPage
	} else {
		index = 0
	}
	if pageNum == len(pageNumbers) && overflow != 0 {
		carsOnPage = landingPageCars[index:index+overflow]
	} else {
		carsOnPage = landingPageCars[index:index+resultsPerPage]
	}
}

	return carsOnPage, pageNumbers
}