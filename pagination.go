package main

import (
	"net/http"
	"strconv"
)

func (p Pagination) GetPreviousPage(currentPage int) int {
	previousPage := currentPage - 1
	return previousPage
}

func (p Pagination) GetNextPage(currentPage int) int {
	nextPage := currentPage + 1
	return nextPage
}

func (p Pagination) NotBeginning(currentPage int) bool {
	var notBeg bool
	if currentPage != 1 {
		notBeg = true
	}
	return notBeg
}

func (p Pagination) NotEnd(currentPage, lastPage int) bool {
	var notEnd bool
	if currentPage != lastPage {
		notEnd = true
	}
	return notEnd
}

func createPages(landingPageCars []Car, h *HomeHandler, w http.ResponseWriter, r* http.Request) {
	//break cars up into pages
	var numPages int
	var pageNumbers []int
	var index int
	var err error
	var pagesShown []int 

	carsOnPage := landingPageCars
	pageNum := 1
	pageNumbers = []int{1}
	resultsPerPage := 5
	overflow := 0

	if len(landingPageCars) > resultsPerPage {
	
		if len(landingPageCars)%resultsPerPage == 0 {
			numPages = len(landingPageCars)/resultsPerPage
		} else {
			numPages = len(landingPageCars)/resultsPerPage + 1
			overflow = len(landingPageCars)%resultsPerPage
		}
		pageNumbers = []int{}
		for num := range numPages {
			//start pageNumbers at 1 instead of 0
			pageNumbers = append(pageNumbers, num+1)
		}
		if r.URL.Query().Get("page") != "" {
			pageNum, err = strconv.Atoi(r.URL.Query().Get("page"))
			if err != nil {
				http.Error(w, "error getting page number", http.StatusInternalServerError)
			}
			if pageNum > len(pageNumbers) {
				pageNum = len(pageNumbers)
			} else if pageNum < 1 {
				pageNum = 1
			}
			index = (pageNum * resultsPerPage) - resultsPerPage
		} else {
			index = 0
		}
		
		numPagesShown := 9
		if len(pageNumbers) < numPagesShown {
			numPagesShown = len(pageNumbers)
		}
		//if current page is in the first 4 pages
		if pageNum <= numPagesShown/2 {
			for i:=1; i<=numPagesShown; i++ {
				pagesShown = append(pagesShown, i)
			}	
		} else if pageNum >= len(pageNumbers) - numPagesShown/2{ //if current page is in the last four pages
			for i:=len(pageNumbers)-numPagesShown; i<=len(pageNumbers); i++ {
				if i<=0 {
					continue
				}
				pagesShown = append(pagesShown, i)
			}
		} else {
			for i:=pageNum-4; i<=pageNum+4; i++ {
				pagesShown = append(pagesShown, i)
			}
		}
		if pageNum == len(pageNumbers) && overflow != 0 {
			carsOnPage = landingPageCars[index:index+overflow]
		} else {
			carsOnPage = landingPageCars[index:index+resultsPerPage]
		}
	}
	h.CarsOnPage = carsOnPage
	h.Pagination.PageNumbers = pagesShown
	h.Pagination.CurrentPage = pageNum
	h.Pagination.LastPage = len(pageNumbers)
}
