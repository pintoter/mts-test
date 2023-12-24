# mts-test
Microservices

How to run local:
  1) touch ./store-service/.env
  2) open ./store-service/.env and fill:
      # Database
      DB_USER = username
      DB_PASSWORD = password
      DB_HOST = "localhost" 
      DB_PORT = 5432
      DB_NAME = "store"
      DB_SSLMODE = "disable"

  3) make run-local
  4) cd store-service && make run
  5) cd order-service && make run

  Simple test:
  1) Open Postman for gRPC Request:
    Addr: 0.0.0.0:7001
    Example request: 
      {
        "item_id": "2",
        "user_id": "1"
      }
  2) Kafka-UI: http://localhost:8080


Create .env in directory ./ and fill:
  DB_USER = postgres
  DB_PASSWORD = 123qweASD
  DB_PORT = 5432
  DB_NAME = store
Create .env in dicrectory ./store-service/ and fill:
  DB_USER = username
  DB_PASSWORD = password
  DB_HOST = "localhost" 
  DB_PORT = 5432
  DB_NAME = "store"
  DB_SSLMODE = "disable"
