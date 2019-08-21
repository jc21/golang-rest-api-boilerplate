# Go REST API Boilerplate

This project is meant to be a feature rich and somewhat simple boilerplate for a
REST JSON API microservice.


## Features

- Router of choice is [Chi](https://github.com/go-chi/chi)
- standard error responses
- swagger and json-schema validation
- jwt



## Usage

environment variables


## Building

```bash
go build -ldflags="-X main.commit=$(git log -n 1 --format=%h)" -o bin/server ./cmd/server/main.go
```


## Development

For easy development, [Tasker](https://taskfile.dev) support is included, refer to the
[installation instructions](https://taskfile.dev/#/installation) then:

```bash
task -w
```

And start writing!


## Justifications

Chi router was selected after much research and deliberation on the available routers.
Chi offers a balance of rich features with performance. The ready-to-go middleware
makes life easier too.

