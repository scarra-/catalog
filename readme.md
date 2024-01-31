## Rest API microservice example written in Fiber Go.
This is example Fiber service to demonstrate microservice written in Go.

## Available routes:
- `GET /blueprints/:id` Get blueprint by ID.
- `GET /blueprints` - List blueprints (offset pagination).
- `POST /blueprints` - Create blueprint.
- `GET /metrics` - Get prometheus metrics.

## Used libraries
- `gofiber/fiber` HTTP API framework.
- `joho/godotenv` Lib to retrieve .env variables.
- `prometheus/client_golang` - Go Prometheus support.
- `go-redis/extra/redisprometheus/v9` - Prometheus redis exporter.
- `go.uber.org/zap` - Logger.
- `gorm.io/gorm` - ORM library.

## Installation
To install service locally run `make dcu` which will run docker-compose up.
Application has 3 components:
- Fiber server
- Redis
- Mysql

## Helm chart
There is helm chart in `./deployments` folder but for now it only works with minikube.

## Postman
<img width="2040" alt="Screenshot 2024-01-31 at 19 56 01" src="https://github.com/scarra-/catalog/assets/10938706/6a7b5d68-89fd-4d51-85e4-62c39c24cdb6">
