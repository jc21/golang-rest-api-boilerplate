# Go REST API Boilerplate

This project is meant to be a feature rich and somewhat simple boilerplate for a
REST JSON API microservice.


## Features

- router
- standard error responses
- swagger and json-schema validation


## Usage

environment variables


## Building

```bash
go build -ldflags="-X main.commit=$(git log -n 1 --format=%h)" -o bin/server ./cmd/server/main.go
```


## Development

For easy development, Tasker support is included, refer to the installation instructions then:

```bash
task -w
```

And start writing!
