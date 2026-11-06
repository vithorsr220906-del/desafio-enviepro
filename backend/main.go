package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Conecta ao banco de dados
	ConectarBanco()
	defer DB.Close()

	// Cria as tabelas
	CriarTabelas()

	// Insere os dados iniciais
	CriarClientes()

	// Cadastra as oportunidades do funil
	CriarOportunidades()

	// Cadastra as vendas
	CriarVendas()

	// Cria o servidor
	servidor := gin.Default()

	// Configura o CORS para permitir acesso do frontend
	servidor.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type", "Authorization"},
	}))

	// Rota principal
	servidor.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"mensagem": "Servidor EnviePro funcionando!",
		})
	})

	// Rota para listar clientes
	servidor.GET("/clientes", ListarClientes)

	// Rota para listar oportunidades do funil
	servidor.GET("/funil", ListarFunil)

	// Rota para listar vendas
	servidor.GET("/vendas", ListarVendas)

	// Rota para resumo mensal das vendas
	servidor.GET("/vendas/resumo", ResumoVendas)

	// Inicia o servidor
	servidor.Run(":8080")
}
