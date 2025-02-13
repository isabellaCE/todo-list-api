# 📌 Todo List API

Uma API simples para gerenciar uma lista de tarefas, desenvolvida em **Golang** com **PostgreSQL**. Este projeto tem como objetivo estudar e aplicar os princípios do **Clean Architecture** e **Clean Code**, garantindo um código mais organizado, testável e sustentável.

## 🚀 Tecnologias Utilizadas

- **Golang** (sem o framework Gin)
- **PostgreSQL**
- **DBeaver** (para gerenciamento do banco de dados)
- **Postman** (para testes de API)

## 📂 Estrutura do Projeto

```
│── /entities       # Definição das entidades do domínio
│── /usecases       # Regras de negócio
│── /repositories   # Acesso ao banco de dados
│── /controllers    # Handlers (Controllers)
│── /config         # Configuração da aplicação
```

## ⚙️ Configuração do Banco de Dados

1. **Criar o banco de dados no PostgreSQL:**
   ```sql
   CREATE DATABASE todo_list;
   ```
2. **Criar a tabela de tarefas:**
   ```sql
   CREATE TABLE public.tasks (
       id SERIAL PRIMARY KEY,
       title VARCHAR(255) NOT NULL,
       description TEXT,
       completed BOOLEAN DEFAULT false,
       created_at TIMESTAMP DEFAULT NOW()
   );
   ```

## 🔧 Configuração do Projeto

1. Clone o repositório:
   ```sh
   git clone https://github.com/SEU_USUARIO/todo-list-api.git
   cd todo-list-api
   ```
2. Configure as variáveis de ambiente (`.env`):
   ```env
   DB_HOST=localhost
   DB_USER=seu_usuario
   DB_PASSWORD=sua_senha
   DB_NAME=todo_list
   DB_PORT=5432
   ```
3. Execute a aplicação:
   ```sh
   go run main.go
   ```

## 📌 Endpoints da API

### 🔹 Criar uma nova tarefa
**`POST /tasks`**

**Exemplo de requisição:**
```json
{
  "title": "Aprender Golang",
  "description": "Criar um projeto de Lista de Tarefas seguindo Clean Code",
  "completed": false
}
```

