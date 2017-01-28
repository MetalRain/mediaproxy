# Mediaproxy API

Mediaproxy API is responsible of:
* Managing media sources
* Fetching media feeds

## How to build

API can be built as Docker image with:

```mediaproxy/api$ docker build -t mediaproxy-api .```

## How to run locally

API container exposes port 8080

Run built docker image:

```docker run -p 127.0.0.1:80:8080 mediaproxy-api```

and browse to http://localhost

## How to deploy

Deploy to kubernetes cluster

```kubectl create -f deployments/api```

## Dependencies

- [Go](https://golang.org/)
- [Iris Framework](https://github.com/kataras/iris)
