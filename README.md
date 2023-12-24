# mts-test
Microservices

How to run local:
  1) touch ./store-service/.env
  2) open ./store-service/.env and fill:
      # Database
export DB_USER = username
export DB_PASSWORD = password
export DB_HOST = "localhost" 
export DB_PORT = 5432
export DB_NAME = "store"
export DB_SSLMODE = "disable"

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
