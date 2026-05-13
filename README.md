# Go API - Aplicação de Estudo

Esta é uma API simples desenvolvida em Go para fins de estudo da linguagem. A aplicação permite operações CRUD (Criar, Ler, Atualizar, Deletar) em produtos, utilizando PostgreSQL como banco de dados.

**Nota:** Esta aplicação é apenas para fins educacionais. No futuro, será implementada autenticação para proteger os endpoints.

## Pré-requisitos

- [Go](https://golang.org/dl/) (versão 1.25 ou superior)
- [Docker](https://www.docker.com/get-started) e [Docker Compose](https://docs.docker.com/compose/install/)

## Como executar

### Usando Docker Compose (Recomendado)

1. Clone o repositório:
   ```bash
   git clone <url-do-repositorio>
   cd go-api
   ```

2. Execute o comando para iniciar os serviços:
   ```bash
   docker-compose up --build
   ```

   Isso irá:
   - Construir a imagem da aplicação Go
   - Iniciar um container PostgreSQL
   - Executar a aplicação na porta 8000

3. A API estará disponível em `http://localhost:8000`

### Executando localmente

1. Certifique-se de ter PostgreSQL rodando localmente ou em um container separado.

2. Configure as variáveis de conexão no arquivo `db/conn.go` se necessário (atualmente configurado para Docker).

3. Instale as dependências:
   ```bash
   go mod tidy
   ```

4. Execute a aplicação:
   ```bash
   go run cmd/main.go
   ```

5. A API estará disponível em `http://localhost:8000`

## Endpoints

- `GET /ping` - Verifica se a API está funcionando
- `GET /products` - Lista todos os produtos
- `POST /product` - Cria um novo produto
- `GET /product/:productId` - Obtém um produto por ID
- `PUT /product/:productId` - Atualiza um produto por ID
- `DELETE /product/:productId` - Deleta um produto por ID

## Estrutura do Projeto

- `cmd/` - Ponto de entrada da aplicação
- `controller/` - Controladores da API
- `db/` - Conexão com o banco de dados
- `model/` - Modelos de dados
- `repository/` - Camada de acesso aos dados
- `usecase/` - Lógica de negócio

## Futuro

- Implementação de autenticação e autorização
- Validação de entrada
- Tratamento de erros aprimorado
- Testes unitários e de integração</content>
<parameter name="filePath">/home/guifeh/Documents/go-api/README.md