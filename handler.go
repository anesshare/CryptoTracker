package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func getCurrencies(w http.ResponseWriter, r *http.Request) {
	row, err := dbPool.Query(context.Background(), "SELECT * FROM valute")
	if err != nil {
		fmt.Println("Greska u upitu")
	}
	defer row.Close()
	crypto := []Coin{}
	for row.Next() {
		var c Coin
		err = row.Scan(&c.ID, &c.Ime, &c.Cena, &c.Created_at)
		if err != nil {
			fmt.Println("Greska u fetchu iz baze")
		}
		crypto = append(crypto, c)
	}
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(crypto)
}
