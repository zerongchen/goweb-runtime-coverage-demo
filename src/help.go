package main

import (
	"fmt"
	"net/http"
	"strconv"
)

 func handleRadon(i string, w http.ResponseWriter) {
	parseInt, error := strconv.ParseInt(i, 0, 32)
	if error != nil {
		fmt.Fprintf(w, error.Error())
	}
	if parseInt < 0 {
		fmt.Fprintf(w, "you query number is less than zero")
	} else if parseInt > 0 {
		fmt.Fprintf(w, "you query number is greater than zero")
	} else {
		fmt.Fprintf(w, "you query number is equal zero")
	}

}
