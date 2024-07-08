# Starter Go Project with hot reloading

Starter project for go using docker

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

### .env

```terminal
    DB_USER=
    DB_PASSWORD=
    DB_NAME=
```
