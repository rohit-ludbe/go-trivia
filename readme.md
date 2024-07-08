# Trivia app in go

Docker tutorial using go.

```terminal

  docker compose up // to create  a container

  docker compose run --service-ports web bash // to run a container

```

### create a container for go project

```terminal
  docker compose up
```

### access docker image to configure project

```terminal
   docker compose run --service-ports web bash
```

### run go project and bind it to our localhost

```terminal
   go run cmd/main.go -b 0.0.0.0
```

### update docker configuration is changed

```terminal
    docker compose build
```
