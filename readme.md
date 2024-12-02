UNICAMPAIGN/
├── cmd/
│   └── server/
│       └── main.go                   
├── infrastructure/
│   ├── persistence/
│   │   └── database.go               
│   └── router/
│       └── router.go                
├── internal/
│   ├── adapter/
│   │   ├── controllers/
│   │   │   └── opportunities_controller.go # Controlador de oportunidades
│   │   └── db/
│   │       └── opportunities_repository.go # Repositório de oportunidades
│   ├── domain/
│   │   ├── benefits.go               # Entidade: Benefícios
│   │   └── opportunities.go          # Entidade: Oportunidades
│   └── usecases/
│       └── opportunities_usecase.go  # Caso de uso de oportunidades
├── go.mod                            # Arquivo de dependências Go
├── go.sum                            # Resumo das dependências Go
└── readme.md                         # Documentação do projeto
