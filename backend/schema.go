package main

import "log"

func CriarTabelas() {
	comandos := []string{

		`CREATE TABLE IF NOT EXISTS clientes (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			nome TEXT NOT NULL,
			empresa TEXT,
			email TEXT,
			telefone TEXT,
			status TEXT NOT NULL,
			responsavel TEXT,
			ultimo_contato DATE
		);`,

		`CREATE TABLE IF NOT EXISTS oportunidades (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			cliente_id INTEGER NOT NULL,
			titulo TEXT NOT NULL,
			etapa_funil TEXT NOT NULL,
			valor REAL NOT NULL,
			data_criacao DATE,
			FOREIGN KEY (cliente_id) REFERENCES clientes(id)
		);`,

		`CREATE TABLE IF NOT EXISTS vendas (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			cliente_id INTEGER NOT NULL,
			valor REAL NOT NULL,
			data DATE NOT NULL,
			produto TEXT,
			FOREIGN KEY (cliente_id) REFERENCES clientes(id)
		);`,
	}

	for _, comando := range comandos {
		_, err := DB.Exec(comando)
		if err != nil {
			log.Fatal("Erro ao criar tabela:", err)
		}
	}

	log.Println("Tabelas criadas com sucesso!")
}
