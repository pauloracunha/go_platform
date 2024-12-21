# portal-ong

Este é o backend do Portal para Gestão de ONGs, desenvolvido em Golang.

## Estrutura do Projeto
```
portal-ong/
├── cmd/portal/          # Código principal do aplicativo
│   └── main.go
├── pkg/                 # Pacotes reutilizáveis
├── internal/            # Lógica de negócios
│   ├── handlers/        # Manipuladores de rotas
│   ├── models/          # Modelos de dados
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
   go get -u gorm.io/driver/sqlite
   ```
2. Execute o projeto:
   ```bash
   go run cmd/portal/main.go
   ```
