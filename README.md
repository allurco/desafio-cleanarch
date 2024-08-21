# Desafio CleanArch

### 1. Iniciar o servidor da base de dados e do RabbitMQ:  
```docker compose up -d --build```
### 2. Rodar os Servidores
Entrar na pasta cmd/ordersystem e executar o programa abaixo:  
```go run main.go wire_gen.go```   
Os servidores estarão rodando nas pastas especificadas
### 3. Para testar:
- Para testar o servidor REST existe um arquivo em /api/api.http contendo os exemplos de chamadas testáveis  
- Para testar o servido gRPC pode-se utilizar o evans  
- Para testar o servidor de GraphQL entrar em http://localhost:8080/
