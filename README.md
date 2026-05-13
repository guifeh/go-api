# Go API - Aplicação de Estudo

Esta é uma API simples desenvolvida em Go para fins de estudo da linguagem. A aplicação permite operações CRUD em produtos com autenticação JWT, utilizando PostgreSQL como banco de dados.

## Pré-requisitos

- [Go](https://golang.org/dl/) (versão 1.25 ou superior)
- [Docker](https://www.docker.com/get-started) e [Docker Compose](https://docs.docker.com/compose/install/)

## Configuração

### Banco de Dados

Crie a tabela de usuários no PostgreSQL:

```sql
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL
);
```

### Variáveis de Ambiente

| Variável | Descrição | Valor Padrão |
|---|---|---|
| `JWT_SECRET` | Chave secreta para assinar os tokens JWT | `default_secret_key` |

## Como executar

### Usando Docker Compose (Recomendado)

1. Clone o repositório:
   ```bash
   git clone <url-do-repositorio>
   cd go-api
   ```

2. Execute o comando para iniciar os serviços:
   ```bash
   JWT_SECRET=sua-chave-secreta docker-compose up --build
   ```

   Isso irá:
   - Construir a imagem da aplicação Go
   - Iniciar um container PostgreSQL
   - Executar a aplicação na porta 8000

3. A API estará disponível em `http://localhost:8000`

### Executando localmente

1. Certifique-se de ter PostgreSQL rodando localmente ou em um container separado.

2. Configure as variáveis de conexão no arquivo `db/conn.go` se necessário.

3. Instale as dependências:
   ```bash
   go mod tidy
   ```

4. Execute a aplicação:
   ```bash
   JWT_SECRET=sua-chave-secreta go run cmd/main.go
   ```

5. A API estará disponível em `http://localhost:8000`

## Endpoints

### Públicos

| Método | Rota | Descrição |
|---|---|---|
| `GET` | `/ping` | Verifica se a API está funcionando |
| `POST` | `/register` | Registra um novo usuário |
| `POST` | `/login` | Realiza login e retorna um token JWT |

### Protegidos (requerem token JWT)

Todas as rotas abaixo exigem o header `Authorization: Bearer <token>`.

| Método | Rota | Descrição |
|---|---|---|
| `GET` | `/api/products` | Lista todos os produtos |
| `POST` | `/api/product` | Cria um novo produto |
| `GET` | `/api/product/:productId` | Obtém um produto por ID |
| `PUT` | `/api/product/:productId` | Atualiza um produto por ID |
| `DELETE` | `/api/product/:productId` | Deleta um produto por ID |

## Como usar a autenticação

### 1. Registrar um usuário

```bash
curl -X POST http://localhost:8000/register \
  -H "Content-Type: application/json" \
  -d '{"name":"João","email":"joao@email.com","password":"123456"}'
```

Resposta:
```json
{
  "id": 1,
  "name": "João",
  "email": "joao@email.com",
  "password": ""
}
```

### 2. Fazer login

```bash
curl -X POST http://localhost:8000/login \
  -H "Content-Type: application/json" \
  -d '{"email":"joao@email.com","password":"123456"}'
```

Resposta:
```json
{
  "token": "eyJhbGciOiJIUzI1NiIs..."
}
```

### 3. Acessar rotas protegidas

Use o token retornado no header `Authorization`:

```bash
curl http://localhost:8000/api/products \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIs..."
```

## Estrutura do Projeto

```
cmd/           - Ponto de entrada da aplicação
config/        - Configurações compartilhadas (JWT secret)
controller/    - Controladores da API
db/            - Conexão com o banco de dados
middleware/    - Middlewares (autenticação JWT)
model/         - Modelos de dados
repository/    - Camada de acesso aos dados
usecase/       - Lógica de negócio
```
