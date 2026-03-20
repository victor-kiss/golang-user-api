

# Go User Management API (Em Desenvolvimento 🚧)

Este projeto consiste em uma API REST de gerenciamento de usuários desenvolvida em **Go**, focada em alta performance e escalabilidade. Atualmente, o projeto utiliza a infraestrutura do **Supabase** (PostgreSQL) e segue uma arquitetura desacoplada em camadas.

## 🚀 Tecnologias em Uso

* **Linguagem:** Go (Golang) 1.26+
* **Web Framework:** [Gin Gonic](https://github.com/gin-gonic/gin)
* **Database Driver:** [pgx/v5](https://github.com/jackc/pgx) (Gerenciamento de Pool de conexões)
* **Banco de Dados:** PostgreSQL (Supabase)
* **Configuração:** Godotenv para variáveis de ambiente

## 🛠️ Desafios Técnicos e Arquitetura

O projeto está sendo estruturado sob o padrão **Repository Pattern**, o que garante a separação entre a lógica de roteamento (Handlers) e a persistência de dados.

### Transição de GORM para PGX
Uma das principais decisões arquiteturais deste projeto é a utilização do **`pgx/v5`** em vez de um ORM tradicional como o GORM. 

**Por que estou utilizando PGX?**
1.  **Performance:** Menor overhead de memória e execução de queries mais rápidas.
2.  **Controle de SQL:** Implementação direta de comandos SQL (como a cláusula `RETURNING`) para garantir integridade dos dados retornados pelo banco.
3.  **Resiliência:** Configuração manual do `pgxpool` para lidar com a latência e conexões persistentes com o Supabase.

### Estratégia de Identificação Híbrida
A tabela `users_db` está sendo implementada com uma estrutura de **ID Duplo**:
* **ID (Serial):** Chave primária para indexação eficiente e relacionamentos internos.
* **UUID (v4):** Identificador aleatório exposto na API para garantir a segurança dos dados e evitar a enumeração de usuários por terceiros.

## 📋 Como Configurar e Testar

1. **Clone o repositório:**
   ```bash
   git clone [https://github.com/seu-usuario/seu-repositorio.git](https://github.com/seu-usuario/seu-repositorio.git)
   ```

2. **Configure o arquivo `.env`:**
   Na raiz do projeto, adicione suas credenciais do Supabase:
   ```env
   DB_HOST=seu-projeto.pooler.supabase.com
   DB_PORT=6543
   DB_USERNAME=postgres.seu-id
   DB_PASSWORD=sua-senha
   DB_NAME=postgres
   DB_SSL_MODE=require
   ```

3. **Inicie o servidor:**
   ```bash
   go run cmd/server/main.go
   ```

## 📍 Roadmap de Desenvolvimento

- [x] Configuração de conexão com Supabase via pgxpool.
- [x] Implementação de Handler de Status da API.
- [x] Criação de usuários com retorno de ID, UUID e Timestamps.
- [x] Tratamento de erros de banco (Conflitos 409).
- [ ] Implementação de busca de usuários por UUID.
- [ ] Hash de senhas com Bcrypt.
- [ ] Implementação de Soft Delete.

## 🧠 Foco Atual
No momento, o foco está na **otimização do parse de dados JSON** e na garantia de que a comunicação entre o **Gin** e o **pgx** ocorra sem vazamento de recursos ou ponteiros nulos.

