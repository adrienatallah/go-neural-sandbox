package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Println("Welcome!\n")

	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local)
	fmt.Println("Go launched at \n", t.Local())

	now := time.Now()
	fmt.Println("The time is now \n", now.Local())

	priceResponse := new(PriceResponse)

	priceResponse.CurrentDate = t
	priceResponse.PredictionDate = now
	priceResponse.Currency = "BTC"
	priceResponse.Value = 12390

	out, err := json.Marshal(priceResponse)

	if err != nil {
		fmt.Println("error marshalling json: ", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}