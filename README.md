## Dêe uma estreal! :star:
Se vc gostou do projeto temperature-system-by-cep-gcp, por favor dêe uma estrela

## Como executar:
Execute do docker-compose:  
docker-compose up -d

Para rodar o projeto execute no prompt de comando na pasta raiz do projeto:  
go run ./cmd/server/main.go

Instale o 'REST Client' no 'VS Code' e execute o teste da pasta:  
./api/api.http

Deploy realizado no Google Cloud Run (free tier) e endereço ativo para ser acessado:  
https://temperature-system-by-cep-gcp-lsf6zzbwaa-uc.a.run.app/v1/temperature/29902555

## Tecnologias implementadas:

go 1.20
 - Router [chi](https://github.com/go-chi/chi)
 