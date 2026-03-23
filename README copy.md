# Desafio Clean Architecture 03 — Clean Architecture: REST + gRPC + GraphQL

Order management system using Clean Architecture, exposing a single `ListOrders` use case through three interfaces simultaneously.

## Starting the Application

```sh
docker compose up --build
```

Migrations run automatically on startup. No manual steps required.

## Service Ports

| Service | Port  | URL                         |
| ------- | ----- | --------------------------- |
| REST    | 8080  | http://localhost:8080/order |
| GraphQL | 8081  | http://localhost:8081/query |
| gRPC    | 50051 | localhost:50051             |

## API

### REST

```sh
# Create order
curl -X POST http://localhost:8080/order \
  -H "Content-Type: application/json" \
  -d '{"name":"Fernando","amount":123.45,"status":"pending"}'

# List orders
curl http://localhost:8080/order
```

### GraphQL

Playground available at http://localhost:8081

```graphql
# List orders
{
  listOrders {
    id
    name
    amount
    status
    createdAt
  }
}

# Create order
mutation {
  createOrder(input: { name: "Fernando", amount: 123.45, status: "pending" }) {
    id
    name
    amount
    status
    createdAt
  }
}
```

### gRPC

Test with [Evans](https://github.com/ktr0731/evans):

```sh
evans -r repl --host localhost --port 50051
```

See `api.http` for ready-to-use HTTP requests (VS Code REST Client).
