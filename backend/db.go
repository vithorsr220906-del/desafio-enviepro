package main

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func ConectarBanco() {
	var err error

	DB, err = sql.Open("sqlite", "./enviepro.db")
	if err != nil {
		log.Fatal("Erro ao abrir banco de dados:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados:", err)
	}

	log.Println("Banco de dados conectado com sucesso!")
}
