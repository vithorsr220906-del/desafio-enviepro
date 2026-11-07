# EnviePro — Sistema CRM

Desafio técnico EnviePro: desenvolvimento de um sistema CRM para gerenciamento de clientes, oportunidades e vendas.

## Tecnologias utilizadas

### Frontend
- Next.js
- React
- TypeScript
- Tailwind CSS
- Lucide React

### Backend
- Go
- Gin
- SQLite

## Funcionalidades

- Dashboard com indicadores de vendas.
- Visualização da quantidade de vendas, faturamento e ticket médio.
- Funil de vendas com as etapas Lead, Contato, Proposta, Negociação e Fechado.
- Gerenciamento de clientes.
- Exibição de informações dos clientes, incluindo nome, empresa, e-mail, telefone, status, responsável e último contato.
- Barra lateral de navegação recolhível.
- Página de configurações.
- Integração entre frontend e backend por API REST.
- Banco de dados SQLite com dados fictícios.

## Como executar o projeto

### Requisitos

- Node.js e npm.
- Go.
- Git.

### 1. Clonar o repositório

```bash
git clone https://github.com/vithorsr220906-del/desafio-enviepro.git
```

Entrar na pasta do projeto:

```bash
cd desafio-enviepro
```

### 2. Executar o backend

Abrir um terminal na pasta principal e executar:

```bash
cd backend
go run .
```

O servidor será iniciado em:

`http://localhost:8080`

O banco de dados SQLite e as tabelas são inicializados automaticamente.

### 3. Executar o frontend

Abrir um segundo terminal na pasta principal do projeto e executar:

```bash
cd frontend
npm install
npm run dev
```

O frontend estará disponível em:

`http://localhost:3000`

## Endpoints da API

| Método | Rota | Descrição |
|---|---|---|
| GET | `/` | Verificação do servidor |
| GET | `/clientes` | Lista de clientes |
| GET | `/funil` | Oportunidades do funil de vendas |
| GET | `/vendas` | Lista de vendas |
| GET | `/vendas/resumo` | Resumo dos indicadores de vendas |

## Validação do frontend

Para verificar o código:

```bash
npm run lint
```

Para gerar a versão de produção:

```bash
npm run build
```

## Observações

O sistema utiliza dados fictícios para demonstração das funcionalidades. O banco de dados SQLite é criado automaticamente durante a inicialização do backend.