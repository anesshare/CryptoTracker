package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func fetchData() {

	url := "https://min-api.cryptocompare.com/data/price?fsym=BTC&tsyms=USD,EUR,ETH,SOL,BNB"
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Greska prilikom preuzimanja podataka!")
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Greska prilikom ocitavanja podataka")
		return
	}

	var record map[string]float64
	er := json.Unmarshal(body, &record)
	if er != nil {
		fmt.Println("Greska prilikom smestanja podataka")
		return
	}

	for ime, podaci := range record {
		fmt.Println("Valuta:", ime, " Cena: ", podaci)
		addToDB(ime, podaci)
	}

}
func main() {
	DBCon()
	go func() {
		fetchData()
		timeTicker := time.NewTicker(10 * time.Second)
		for range timeTicker.C {
			fmt.Println("Osvezavam podatke...")
			fetchData()
		}
	}()
	http.HandleFunc("/data", getCurrencies)
	fmt.Println("Server pokrenut na 8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println("Greska pri pokretanju servera")
	}

}
