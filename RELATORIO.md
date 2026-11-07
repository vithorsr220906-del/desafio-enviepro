# Relatório de Desenvolvimento — EnviePro

## 1. Organização do tempo

O desenvolvimento do projeto foi realizado em diferentes períodos do dia, sem um horário fixo. Trabalhei um pouco pela manhã, continuei durante a tarde e também em alguns momentos à noite. Organizei as atividades conforme o tempo disponível, avançando gradualmente em cada etapa do desafio.

## 2. Ferramentas utilizadas e inteligência artificial

A única ferramenta de inteligência artificial utilizada foi o ChatGPT. Ele foi utilizado como apoio durante o desenvolvimento, principalmente para explicar os procedimentos, sugerir códigos, orientar a configuração do ambiente e ajudar a identificar e resolver erros.

O Cursor foi utilizado como editor de código, enquanto o Git e o GitHub foram utilizados para controlar as versões do projeto e organizar o envio das alterações.

## 3. Três decisões técnicas importantes

**1. Utilização de Go e Gin no backend**

Foi escolhida a linguagem Go junto com o framework Gin para desenvolver o backend. Essa combinação permitiu criar uma API responsável por fornecer os dados utilizados pelo sistema.

**2. Utilização do SQLite como banco de dados**

O SQLite foi escolhido para armazenar os dados fictícios de clientes, oportunidades e vendas. Essa opção facilitou a configuração do banco de dados, sem a necessidade de instalar ou administrar um servidor de banco de dados separado.

**3. Utilização de Next.js, React e Tailwind CSS no frontend**

Essas tecnologias foram utilizadas para construir a interface do sistema. Com elas, foi possível desenvolver o painel inicial, a página de CRM, a página de configurações e a navegação lateral.

## 4. Dificuldades encontradas e soluções

Durante o desenvolvimento, encontrei algumas dificuldades:

- **Bloqueio ao executar o npm no PowerShell:** inicialmente, o Windows impediu a execução do comando `npm` devido à política de execução de scripts. A solução foi utilizar `npm.cmd`, permitindo executar os comandos necessários.
- **Página de configurações ausente na cópia local:** em determinado momento, a página de configurações não aparecia nos arquivos locais do projeto. Para resolver, atualizei as informações do repositório com `git fetch` e `git pull origin main`, recuperando os arquivos que já estavam no GitHub.

## 5. Aprendizados e limitações

Antes do desafio, eu não tinha conhecimento prático sobre o desenvolvimento de um sistema com frontend, backend, banco de dados e integração entre essas partes.

Durante o projeto, utilizei o ChatGPT como guia para executar cada etapa. Consegui acompanhar os procedimentos e compreender um pouco melhor a função de algumas ferramentas e tecnologias, mas reconheço que ainda preciso estudar e praticar para conseguir desenvolver um projeto semelhante de forma independente.

## 6. Testes realizados

Foram realizados testes para verificar o funcionamento das principais partes do sistema:

- Acesso pelo navegador às páginas Início, CRM e Configurações.
- Verificação das informações apresentadas no painel.
- Testes das rotas da API, incluindo clientes, oportunidades, vendas e resumo das vendas.
- Execução do comando `go test ./...` no backend.
- Execução do comando `npm.cmd run lint` no frontend.
- Execução do comando `npm.cmd run build` para verificar a compilação do frontend.

Os comandos de verificação do backend e do frontend foram executados com sucesso, e as páginas principais foram verificadas no navegador.

## 7. Melhorias futuras

Como possíveis melhorias para o sistema, considero:

- Implementar funcionalidades para cadastrar, editar e excluir clientes.
- Permitir atualizar as etapas das oportunidades no funil de vendas.
- Implementar um sistema de login e autenticação de usuários.
- Melhorar o visual e a responsividade da interface, especialmente em celulares e tablets.
- Aprofundar meus conhecimentos nas tecnologias utilizadas para conseguir desenvolver e manter o sistema com maior independência.