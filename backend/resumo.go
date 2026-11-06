package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Calcula o resumo mensal das vendas
func ResumoVendas(c *gin.Context) {
	var quantidade int
	var total float64
	var ticketMedio float64
	var clientesAtivos int

	// Calcula quantidade, total e média das vendas do mês atual
	err := DB.QueryRow(`
        SELECT
            COUNT(*),
            COALESCE(SUM(valor), 0),
            COALESCE(AVG(valor), 0)
        FROM vendas
        WHERE strftime('%Y-%m', data) = strftime('%Y-%m', 'now')
    `).Scan(&quantidade, &total, &ticketMedio)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"erro": "Erro ao calcular resumo das vendas",
		})
		return
	}

	// Conta os clientes ativos
	err = DB.QueryRow(`
        SELECT COUNT(*)
        FROM clientes
        WHERE status = 'Ativo'
    `).Scan(&clientesAtivos)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"erro": "Erro ao contar clientes ativos",
		})
		return
	}

	// Retorna os resultados em JSON
	c.JSON(http.StatusOK, gin.H{
		"quantidade_vendas": quantidade,
		"total_vendas":      total,
		"ticket_medio":      ticketMedio,
		"clientes_ativos":   clientesAtivos,
	})
}
