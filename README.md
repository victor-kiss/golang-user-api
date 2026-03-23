# Go User Management API (Em Desenvolvimento 🚧)

Este projeto é uma API REST de gerenciamento de usuários desenvolvida em **Go**, focada em alta performance, segurança e escalabilidade. A aplicação utiliza a infraestrutura do **Supabase** (PostgreSQL) e segue as melhores práticas de arquitetura da comunidade Go.

## 🚀 Stack Tecnológica

* **Linguagem:** Go (Golang) 1.26+
* **Web Framework:** [Gin Gonic](https://github.com/gin-gonic/gin)
* **Query Builder:** [Squirrel](https://github.com/Masterminds/squirrel) (Geração de SQL fluente e seguro)
* **Database Driver:** [pgx/v5](https://github.com/jackc/pgx) (Gerenciamento de Pool de conexões de alta performance)
* **Banco de Dados:** PostgreSQL (Hospedado no Supabase)
* **Configuração:** Godotenv para gerenciamento de variáveis de ambiente

## 🛠️ Arquitetura e Decisões Técnicas

O projeto utiliza o **Standard Go Project Layout**, garantindo que a lógica de negócio esteja protegida dentro da pasta `internal/`, seguindo princípios de encapsulamento.

### PGX & Squirrel vs ORMs Tradicionais
Optei por não utilizar um ORM (como GORM) para ter controle total sobre a performance e o comportamento do banco de dados:
1.  **Performance:** Menor overhead de memória e execução direta de queries sem camadas de abstração desnecessárias.
2.  **SQL Dinâmico:** O Squirrel permite construir updates parciais de forma segura e legível, facilitando a manutenção.
3.  **Cláusula RETURNING:** Uso nativo do PostgreSQL para retornar IDs, UUIDs e Timestamps no momento da inserção, garantindo sincronia entre Go e o Banco.

### Estratégia de Identificação Híbrida
* **ID (Serial):** Chave primária para indexação interna eficiente e performance em JOINS.
* **UUID (v4):** Identificador público exposto na API, garantindo que usuários não consigam prever ou enumerar IDs de outros registros (Segurança por Design).

## 📂 Estrutura de Pastas

```text
.
├── cmd/server/main.go        # Ponto de entrada e configuração de rotas
├── internal/                 # Código privado da aplicação (encapsulado)
│   ├── database/             # Lógica de conexão e pool (pgxpool)
│   ├── handlers/             # Camada HTTP (Handlers do Gin)
│   ├── repository/           # Camada de Dados (Persistência com Squirrel)
│   ├── models/               # Estruturas de dados (User struct)
│   └── utils/                # Utilitários (Verificação de UUID)
└── README.md
```

## 📋 Como Configurar

1.  **Clone o repositório:**
    ```bash
    git clone [https://github.com/seu-usuario/seu-repositorio.git](https://github.com/seu-usuario/seu-repositorio.git)
    ```
2.  **Configure o arquivo `.env`:**
    Crie um arquivo `.env` na raiz do projeto com sua string de conexão:
    ```env
    DB_URL=postgres://postgres:[SENHA]@db.[ID-PROJETO].supabase.co:5432/postgres
    ```
3.  **Inicie o servidor:**
    ```bash
    go run cmd/server/main.go
    ```

## 📍 Roadmap de Desenvolvimento

- [x] Configuração de conexão resiliente com Supabase via `pgxpool`.
- [x] Implementação do padrão `internal/` para segurança do código.
- [x] **Operações CRUD Completas:**
    - [x] **Create:** Registro de usuário com mapeamento de campos gerados.
    - [x] **Read:** Listagem filtrada (non-deleted) e busca por UUID.
    - [x] **Update:** Atualização dinâmica de campos via `map[string]any`.
    - [x] **Delete:** Implementação de **Soft Delete** (coluna `deleted_at`).
- [x] Utilitário centralizado de verificação de UUID.
- [ ] Implementação de Hash de senhas com Bcrypt.
- [ ] Autenticação JWT e Middlewares de proteção de rotas.
- [ ] Testes de integração com `testify` e `httptest`.

## 🧠 Foco Atual

```
No momento, o foco está na migração total dos comentários internos para **Inglês**, visando alinhar o projeto aos padrões globais de desenvolvimento e facilitar a colaboração em código aberto.
```
