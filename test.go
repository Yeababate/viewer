package main

import (
	//"fmt"
	"net/http"
	//"io"
	//"encoding/json"
)

func testHandler(w http.ResponseWriter, r *http.Request) {
	//car := CarsData.Name

	tpl.ExecuteTemplate(w, "test.html", CarsData)
}
