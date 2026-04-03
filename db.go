package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var dbPool *pgxpool.Pool

func addToDB(ime string, cena float64) {
	query := `
        INSERT INTO valute (ime, cena, created_at) 
        VALUES ($1, $2, NOW()) 
        ON CONFLICT (ime) 
        DO UPDATE SET 
            cena = EXCLUDED.cena, 
            created_at = NOW();
    `
	_, err := dbPool.Exec(context.Background(), query, ime, cena)
	if err != nil {
		fmt.Println("Greska prilikom upisa u bazi", err)
	}
}

func DBCon() {
	connString := "postgres://postgres:postgres@localhost:5432/test"
	var err error
	dbPool, err = pgxpool.New(context.Background(), connString)
	if err != nil {
		fmt.Println("Greska pri konekciji na bazi")
		os.Exit(1)
	}

}
