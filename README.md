# GoBank

Example of a REST application using the following technologies:

- [Go](https://go.dev/) [Language]
- [Gin Web Framework](https://github.com/gin-gonic/gin) [Router]
- [GORM](https://github.com/go-gorm/gorm) [ORM]
- [Testify](https://github.com/stretchr/testify) [Testing-toolkit]
- [Postman](https://www.postman.com/) [API testing tool]
- [k6](https://k6.io/) [API load testing tool]

## Packages overview
![](./documentation/images/packages.svg)

## How to run

### Docker / Docker-Compose
You need [Docker](https://docs.docker.com/engine/install/) version 20.10.11 or greater and [Docker Compose](https://docs.docker.com/compose/install/) version 1.29.2 or greater.

     docker-compose up
     
## Endpoints

The list of endpoints are available in the [Postman collection](./documentation/postman/collection.json). 

## Benchmark
    k6 run --vus 10 --duration 60s script/k6-script.js 
![](./documentation/images/benchmark.png)