# Desafio CleanArch

### 1. Iniciar o servidor da base de dados e do RabbitMQ:  
```docker composer up --build```
### 2. Rodar as migrations:  
Para isso é necessário ter instalado o [go-migrate](https://github.com/golang-migrate/migrate) e na sequência rodar o seginte comando:  
```migrate -path=migrations -database "mysql://root:root@tcp(localhost:33306)/orders" -verbose up```
### 3. Rodar os Servidores
Entrar na pasta cmd/ordersystem e executar o programa abaixo:  
```go run main.go wire_gen.go```   
Os servidores estarão rodando nas pastas especificadas
### 4. Para testar:
Para testar o servidor REST existe um arquivo em /api/api.http contendo os exemplos de chamadas testáveis
Para testar o servido gRPC pode-se utilizar o evans
Para testar o servidor de GraphQL entrar em http://localhost:8080/
