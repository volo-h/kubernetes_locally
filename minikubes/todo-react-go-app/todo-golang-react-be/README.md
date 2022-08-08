### localy in dev mode
  1. config grants for database, connect to db-service and create database for project
  2. go run main.go

### localy with docker-compose

### localy over kubernetes
```sh
  eval $(minikube docker-env)

  docker build -t todo-golang-react-be .
    minikube ssh
    docker images todo-golang-react-be:latest
```
