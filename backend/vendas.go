package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Representa uma venda que será enviada em JSON
type Venda struct {
	ID        int     `json:"id"`
	ClienteID int     `json:"cliente_id"`
	Valor     float64 `json:"valor"`
	Data      string  `json:"data"`
	Produto   string  `json:"produto"`
}

// Busca todas as vendas no banco de dados
func ListarVendas(c *gin.Context) {
	rows, err := DB.Query(`
        SELECT id, cliente_id, valor, data, produto
        FROM vendas
        ORDER BY data DESC
    `)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"erro": "Erro ao buscar vendas",
		})
		return
	}

	defer rows.Close()

	vendas := make([]Venda, 0)

	for rows.Next() {
		var venda Venda

		err := rows.Scan(
			&venda.ID,
			&venda.ClienteID,
			&venda.Valor,
			&venda.Data,
			&venda.Produto,
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"erro": "Erro ao ler vendas",
			})
			return
		}

		vendas = append(vendas, venda)
	}

	c.JSON(http.StatusOK, vendas)
}
