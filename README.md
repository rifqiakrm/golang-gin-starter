# Golang Gin Starter

Gin Starter is a microservices boilerplate with RS256 JWT

It comes pre-configured with :

1. Gin Gonic Router (https://github.com/gin-gonic/gin)
2. JWT-GO (https://github.com/dgrijalva/jwt-go)
3. Testify (https://github.com/stretchr/testify)
4. Jaeger (https://github.com/uber/jaeger-lib)
5. Jaeger Client for Go (https://github.com/uber/jaeger-client-go)
6. Redigo (https://github.com/gomodule/redigo)
7. GORM (https://gorm.io/gorm)
8. Google Pub Sub (https://cloud.google.com/go/pubsub) 
9. Go SQL Mock (https://github.com/DATA-DOG/go-sqlmock) 
10. Go Humanize (https://github.com/dustin/go-humanize) 
11. Golang Mock (https://github.com/golang/mock) 
12. Google UUID(https://github.com/google/uuid) 
13. Env Decode(https://github.com/joeshaw/envdecode) 
14. godotenv (https://github.com/joho/godotenv)
15. Mailgun (https://github.com/mailgun/mailgun-go/v4)

This starter also contains pre-developed API such as user and cms authentication, backend cms, forgot password feature, otp and more.


## Setup

Use this command to install the blueprint

```bash
go get github.com/rifqiakrm/golang-gin-starter
```

or manually clone the repo and then run `go run main.go`.

## How to Run

- Read [Prerequisites](doc/PREREQUISITES.md).
- Then, read [How to Run](doc/HOW_TO_RUN.md).

## Development Guide

- Read [Prerequisites](doc/PREREQUISITES.md).
- Then, read [Development Guide](doc/DEVELOPMENT_GUIDE.md).

## Test

### Unit Test

```sh
$ make tidy
$ make cover
```

### Integration Test / API Test

To run integration test, we need to start all dependencies needed. We provide all dependencies via [Docker Compose](https://docs.docker.com/compose/).
Make sure to install [Docker Compose](https://docs.docker.com/compose/install/) before running integration test.

Also, we need to build the docker image for golang-gin-starter first.

```sh
$ make compile-server
$ make docker-build-server
```

After that, run all images needed using `docker-compose` and run the integration test.

```sh
$ docker-compose up
$ make test.integration
```

## Deployment

Read [Deployment](doc/DEPLOYMENT.md).

### Staging

TBD

### Production

TBD

### Postman Collection

[![Run in Postman](https://run.pstmn.io/button.svg)](https://app.getpostman.com/run-collection/3441134-30c3a8b9-46a3-4ef9-b214-5f47f76709e2?action=collection%2Ffork&collection-url=entityId%3D3441134-30c3a8b9-46a3-4ef9-b214-5f47f76709e2%26entityType%3Dcollection%26workspaceId%3D4c7bdc52-a80d-44cd-9636-496de826d460#?env%5BGin%20Starter%5D=W3sia2V5IjoiYmFzZV91cmwiLCJ2YWx1ZSI6ImxvY2FsaG9zdDoxNzQ0Ni92MSIsImVuYWJsZWQiOnRydWUsInR5cGUiOiJkZWZhdWx0In0seyJrZXkiOiJhY2Nlc3NfdG9rZW4iLCJ2YWx1ZSI6IiIsImVuYWJsZWQiOnRydWUsInR5cGUiOiJkZWZhdWx0In1d)

## FAQs

- Read [FAQs](doc/FAQS.md)