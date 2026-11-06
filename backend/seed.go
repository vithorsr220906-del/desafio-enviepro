package main

import (
	"log"
)

func CriarClientes() {
	var total int

	// Verifica quantos clientes já existem
	err := DB.QueryRow("SELECT COUNT(*) FROM clientes").Scan(&total)
	if err != nil {
		log.Fatal("Erro ao verificar clientes:", err)
	}

	// Evita cadastrar os mesmos clientes novamente
	if total > 0 {
		log.Println("Clientes já cadastrados!")
		return
	}

	clientes := []struct {
		nome        string
		empresa     string
		email       string
		telefone    string
		status      string
		responsavel string
		contato     string
	}{
		{"Ana Souza", "Tech Solutions", "ana@email.com", "11999990001", "Ativo", "Carlos", "2026-11-05"},
		{"Bruno Lima", "Lima Transportes", "bruno@email.com", "11999990002", "Ativo", "Mariana", "2026-11-04"},
		{"Carla Mendes", "Mendes Design", "carla@email.com", "11999990003", "Ativo", "Carlos", "2026-11-03"},
		{"Daniel Costa", "Costa Logística", "daniel@email.com", "11999990004", "Inativo", "Mariana", "2026-10-10"},
		{"Eduarda Alves", "Alves Comércio", "eduarda@email.com", "11999990005", "Ativo", "Pedro", "2026-11-02"},
		{"Felipe Rocha", "Rocha Sistemas", "felipe@email.com", "11999990006", "Ativo", "Pedro", "2026-11-01"},
		{"Gabriela Santos", "GS Marketing", "gabriela@email.com", "11999990007", "Ativo", "Carlos", "2026-10-20"},
		{"Henrique Martins", "HM Construções", "henrique@email.com", "11999990008", "Inativo", "Mariana", "2026-10-15"},
		{"Isabela Ferreira", "IF Consultoria", "isabela@email.com", "11999990009", "Ativo", "Pedro", "2026-11-05"},
		{"João Oliveira", "Oliveira Tech", "joao@email.com", "11999990010", "Ativo", "Carlos", "2026-11-04"},
		{"Karen Ribeiro", "Ribeiro Saúde", "karen@email.com", "11999990011", "Ativo", "Mariana", "2026-10-18"},
		{"Lucas Almeida", "Almeida Digital", "lucas@email.com", "11999990012", "Ativo", "Pedro", "2026-11-03"},
		{"Mariana Carvalho", "MC Eventos", "mariana@email.com", "11999990013", "Inativo", "Carlos", "2026-10-01"},
		{"Nicolas Gomes", "Gomes Financeiro", "nicolas@email.com", "11999990014", "Ativo", "Mariana", "2026-11-02"},
		{"Olivia Barros", "Barros Educação", "olivia@email.com", "11999990015", "Ativo", "Pedro", "2026-11-01"},
		{"Paulo Nunes", "Nunes Automação", "paulo@email.com", "11999990016", "Ativo", "Carlos", "2026-10-25"},
		{"Rafaela Dias", "Dias Imóveis", "rafaela@email.com", "11999990017", "Ativo", "Mariana", "2026-11-05"},
		{"Samuel Castro", "Castro Energia", "samuel@email.com", "11999990018", "Inativo", "Pedro", "2026-09-20"},
		{"Tatiane Lopes", "Lopes Moda", "tatiane@email.com", "11999990019", "Ativo", "Carlos", "2026-11-04"},
		{"Vinicius Moreira", "Moreira Serviços", "vinicius@email.com", "11999990020", "Ativo", "Mariana", "2026-11-03"},
	}

	for _, cliente := range clientes {
		_, err := DB.Exec(`
			INSERT INTO clientes
			(nome, empresa, email, telefone, status, responsavel, ultimo_contato)
			VALUES (?, ?, ?, ?, ?, ?, ?)
		`,
			cliente.nome,
			cliente.empresa,
			cliente.email,
			cliente.telefone,
			cliente.status,
			cliente.responsavel,
			cliente.contato,
		)

		if err != nil {
			log.Fatal("Erro ao cadastrar cliente:", err)
		}
	}

	log.Println("20 clientes cadastrados com sucesso!")
}
func CriarOportunidades() {
	var total int

	err := DB.QueryRow("SELECT COUNT(*) FROM oportunidades").Scan(&total)
	if err != nil {
		log.Fatal("Erro ao verificar oportunidades:", err)
	}

	if total > 0 {
		log.Println("Oportunidades já cadastradas!")
		return
	}

	oportunidades := []struct {
		clienteID int
		titulo    string
		etapa     string
		valor     float64
		data      string
	}{
		{1, "Desenvolvimento de sistema", "Lead", 5000, "2026-11-01"},
		{2, "Automação logística", "Contato", 7500, "2026-11-02"},
		{3, "Site institucional", "Proposta", 3000, "2026-11-03"},
		{4, "Sistema de estoque", "Negociação", 8500, "2026-10-20"},
		{5, "Plataforma de vendas", "Fechado", 12000, "2026-10-15"},
		{6, "Aplicativo empresarial", "Lead", 9500, "2026-11-01"},
		{7, "Marketing digital", "Contato", 2500, "2026-11-02"},
		{8, "Sistema financeiro", "Proposta", 6000, "2026-10-25"},
		{9, "Consultoria de TI", "Negociação", 4000, "2026-11-03"},
		{10, "Integração de sistemas", "Fechado", 15000, "2026-10-18"},
		{11, "Plataforma de saúde", "Lead", 11000, "2026-11-01"},
		{12, "Loja virtual", "Contato", 4500, "2026-11-02"},
		{13, "Sistema de eventos", "Proposta", 7000, "2026-10-10"},
		{14, "Controle financeiro", "Negociação", 5500, "2026-11-04"},
		{15, "Plataforma educacional", "Fechado", 18000, "2026-10-12"},
		{16, "Automação industrial", "Lead", 13000, "2026-11-01"},
		{17, "Sistema imobiliário", "Contato", 8000, "2026-11-02"},
		{18, "Consultoria energética", "Proposta", 6500, "2026-10-22"},
		{19, "E-commerce", "Negociação", 9000, "2026-11-05"},
		{20, "Sistema de atendimento", "Fechado", 10000, "2026-10-14"},
		{1, "Aplicativo mobile", "Contato", 4000, "2026-11-03"},
		{2, "Sistema de rastreamento", "Proposta", 7200, "2026-11-04"},
		{3, "Identidade visual", "Negociação", 2800, "2026-11-05"},
		{4, "Plataforma de entregas", "Lead", 14000, "2026-11-02"},
		{5, "Sistema de pagamentos", "Contato", 16000, "2026-11-03"},
		{6, "Consultoria empresarial", "Proposta", 3500, "2026-11-04"},
		{7, "Gestão de redes sociais", "Negociação", 2200, "2026-11-05"},
		{8, "Sistema de RH", "Fechado", 9500, "2026-10-16"},
		{9, "Plataforma de cursos", "Lead", 6800, "2026-11-03"},
		{10, "Automação comercial", "Contato", 11500, "2026-11-05"},
	}

	for _, oportunidade := range oportunidades {
		_, err := DB.Exec(`
			INSERT INTO oportunidades
			(cliente_id, titulo, etapa_funil, valor, data_criacao)
			VALUES (?, ?, ?, ?, ?)
		`,
			oportunidade.clienteID,
			oportunidade.titulo,
			oportunidade.etapa,
			oportunidade.valor,
			oportunidade.data,
		)

		if err != nil {
			log.Fatal("Erro ao cadastrar oportunidade:", err)
		}
	}

	log.Println("30 oportunidades cadastradas com sucesso!")
}
func CriarVendas() {
	var total int

	err := DB.QueryRow("SELECT COUNT(*) FROM vendas").Scan(&total)
	if err != nil {
		log.Fatal("Erro ao verificar vendas:", err)
	}

	if total > 0 {
		log.Println("Vendas já cadastradas!")
		return
	}

	vendas := []struct {
		clienteID int
		valor     float64
		data      string
		produto   string
	}{
		{1, 5000, "2026-11-01", "Sistema empresarial"},
		{2, 7500, "2026-11-01", "Automação logística"},
		{3, 3000, "2026-11-02", "Site institucional"},
		{5, 12000, "2026-11-02", "Plataforma de vendas"},
		{6, 9500, "2026-11-02", "Aplicativo empresarial"},
		{7, 2500, "2026-11-03", "Marketing digital"},
		{9, 4000, "2026-11-03", "Consultoria de TI"},
		{10, 15000, "2026-11-03", "Integração de sistemas"},
		{11, 11000, "2026-11-03", "Plataforma de saúde"},
		{12, 4500, "2026-11-04", "Loja virtual"},
		{14, 5500, "2026-11-04", "Controle financeiro"},
		{15, 18000, "2026-11-04", "Plataforma educacional"},
		{16, 13000, "2026-11-04", "Automação industrial"},
		{17, 8000, "2026-11-05", "Sistema imobiliário"},
		{19, 9000, "2026-11-05", "E-commerce"},
		{20, 10000, "2026-11-05", "Sistema de atendimento"},
		{1, 4000, "2026-11-05", "Aplicativo mobile"},
		{2, 7200, "2026-11-05", "Sistema de rastreamento"},
		{5, 16000, "2026-11-06", "Sistema de pagamentos"},
		{8, 9500, "2026-11-06", "Sistema de RH"},
		{9, 6800, "2026-11-06", "Plataforma de cursos"},
		{10, 11500, "2026-11-06", "Automação comercial"},
		{13, 7000, "2026-11-06", "Sistema de eventos"},
		{18, 6500, "2026-11-06", "Consultoria energética"},
		{4, 8500, "2026-11-06", "Sistema de estoque"},
	}

	for _, venda := range vendas {
		_, err := DB.Exec(`
			INSERT INTO vendas
			(cliente_id, valor, data, produto)
			VALUES (?, ?, ?, ?)
		`,
			venda.clienteID,
			venda.valor,
			venda.data,
			venda.produto,
		)

		if err != nil {
			log.Fatal("Erro ao cadastrar venda:", err)
		}
	}

	log.Println("25 vendas cadastradas com sucesso!")
}
