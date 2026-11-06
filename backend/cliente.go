package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Representa um cliente que será enviado em JSON
type Cliente struct {
	ID            int    `json:"id"`
	Nome          string `json:"nome"`
	Empresa       string `json:"empresa"`
	Email         string `json:"email"`
	Telefone      string `json:"telefone"`
	Status        string `json:"status"`
	Responsavel   string `json:"responsavel"`
	UltimoContato string `json:"ultimo_contato"`
}

// Busca todos os clientes no banco de dados
func ListarClientes(c *gin.Context) {
	rows, err := DB.Query("SELECT id, nome, empresa, email, telefone, status, responsavel, ultimo_contato FROM clientes")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"erro": "Erro ao buscar clientes",
		})
		return
	}

	defer rows.Close()

	clientes := make([]Cliente, 0)

	for rows.Next() {
		var cliente Cliente

		err := rows.Scan(
			&cliente.ID,
			&cliente.Nome,
			&cliente.Empresa,
			&cliente.Email,
			&cliente.Telefone,
			&cliente.Status,
			&cliente.Responsavel,
			&cliente.UltimoContato,
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"erro": "Erro ao ler clientes",
			})
			return
		}

		clientes = append(clientes, cliente)
	}

	c.JSON(http.StatusOK, clientes)
}
