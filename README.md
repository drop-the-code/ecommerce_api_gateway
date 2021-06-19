# Ecommerce Gateway

Esta aplicação esta contida em um container docker,
e sua responsabilidade é interligar o front-end via requisições HTTP,
e repassa as requisições recebida do front-end para os micro-serviços via gRPC.

### tecnologias utilizada
docker
gRPC
HTTP
GOlang




# Run Api in development  environment

```
$ air -c air.conf
```
```
$ protoc --go_out=. --go_opt=paths=source_relative \                
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    protos/user.proto
```
