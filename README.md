# GO Platform

Esta é uma API base,desenvolvida em Golang, que pode ser utilizada para iniciar o desenvolvimento de plataformas.

## Estrutura do Projeto
```
/
├── cmd/portal/          # Código principal do aplicativo
│   └── main.go
├── pkg/                 # Pacotes reutilizáveis
├── internal/            # Lógica de negócios
│   ├── handlers/        # Manipuladores de rotas
│   ├── models/          # Modelos de dados
|   ├── routes/          # Definição de roteamento
│   └── services/        # Regras de negócio
├── migrations/          # Scripts de migração do banco de dados
├── configs/             # Configurações
└── go.mod               # Gerenciador de dependências do Go
```

## Inicialização
1. Instale as dependências:
   ```bash
   go get -u github.com/gorilla/mux
   go get -u gorm.io/gorm
   go get -u gorm.io/driver/mysql
   ```
2. Execute o projeto:
   ```bash
   go run cmd/portal/main.go
   ```
