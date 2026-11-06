package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Oportunidade struct {
	ID          int     `json:"id"`
	ClienteID   int     `json:"cliente_id"`
	Titulo      string  `json:"titulo"`
	Etapa       string  `json:"etapa_funil"`
	Valor       float64 `json:"valor"`
	DataCriacao string  `json:"data_criacao"`
}

func ListarFunil(c *gin.Context) {
	rows, err := DB.Query(`
		SELECT id, cliente_id, titulo, etapa_funil, valor, data_criacao
		FROM oportunidades
	`)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"erro": "Erro ao buscar oportunidades",
		})
		return
	}

	defer rows.Close()

	oportunidades := make([]Oportunidade, 0)

	for rows.Next() {
		var oportunidade Oportunidade

		err := rows.Scan(
			&oportunidade.ID,
			&oportunidade.ClienteID,
			&oportunidade.Titulo,
			&oportunidade.Etapa,
			&oportunidade.Valor,
			&oportunidade.DataCriacao,
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"erro": "Erro ao ler oportunidades",
			})
			return
		}

		oportunidades = append(oportunidades, oportunidade)
	}

	c.JSON(http.StatusOK, oportunidades)
}
