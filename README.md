# Testando a Funcionalidade ListOrders

## GraphQL

Para testar a consulta `ListOrders` no GraphQL, execute a seguinte query, na porta http://localhost:8080/:

```graphql
query {
  listOrders {
    id
    Price
    FinalPrice
    Tax
  }
}
```

## gRPC


Para testar o serviço `ListOrders` com gRPC, use o seguinte comando no terminal:

```
grpcurl -plaintext -proto <caminho_para_protofile> <endereço_servidor> pb.OrderService/ListOrders
```

## HTTP

Para testar o endpoint HTTP, você pode usar o arquivo preparado list_orders.http. Execute a requisição neste arquivo usando uma ferramenta como HTTPie ou curl.