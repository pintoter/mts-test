# order-path

## Examples

[![Golang](https://img.shields.io/badge/Go-v1.21-EEEEEE?logo=go&logoColor=white&labelColor=00ADD8)](https://go.dev/)
[![License](https://img.shields.io/badge/license-MIT-green)](LICENSE)

<div align="center">
    <h1>Order path</h1>
    <h5>
        Microservices written in Go as test task 
    </h5>
</div>

## Navigation
* **[Task](#task)**
* **[Installation](#installation)**
* **[Getting started](#Getting_started)**

---

## Task

1. Implement 2 services: `order-service` & `store-service`
2. The first service `order-service` should receive a request via **[gRPC](https://grpc.io)** and send data to **[Kafka](https://kafka.apache.org)**

```proto
message CreateOrderRequest{
 int64 user_id = 1;
 int64 item_id = 2;
}

message CreateOrderResponse{
  string message = 1;
}
```

3. The second service `store-service` should reviece a message from **[Kafka](https://kafka.apache.org)** and write it to **[PostgreSQL](https://www.postgresql.org)**

```sql
CREATE TABLE IF NOT EXISTS store (
  id SERIAL PRIMARY KEY,
  user_id BIGINT NOT NULL,
  item_id BIGINT NOT NULL,
  created_at DATE NOT NULL DEFAULT CURRENT_DATE
);
```

---

## Installation
```shell
git clone https://github.com/pintoter/mts-test.git
```

---

## Getting started

1. **Create .env file with filename ".env" in the project root and setting up environment your own variables:**
```dotenv
  DB_USER = postgres
  DB_PASSWORD = 123qweASD
  DB_PORT = 5432
  DB_NAME = store
```

2. **Create .env file with filename ".env" in the `store-service` root and setting up environment your own variables:**
```dotenv
  # Database
  DB_USER = username
  DB_PASSWORD = password
  DB_HOST = "localhost" 
  DB_PORT = 5432
  DB_NAME = "store"
  DB_SSLMODE = "disable"
```
> **Hint:**
if you are running the project using Docker, set `DB_HOST` to "**store-service-postgres**" (as the service name of Postgres in the docker-compose).

3. **Compile and run the project:**
For starting:
```shell
make version=prod
```

4. **To test functionality, you can open `Postman` and create request to order-service: `0.0.0.0:8001`
  Access to Kafka-UI: http://localhost:8080/ .**

* Example request in Postman:
```json
{
    "item_id": "1",
    "user_id": "10"
}
```

* Example response in Postman:
```json
{
    "message": "Your order has been successfully created!"
}
```

* Example record in PostgreSQL:
```sql
| id | user_id | item_id | item_id    |
| -- | ------- | ------- | ---------- |
| 1  | 10      | 1       | 2023-12-24 |
```
