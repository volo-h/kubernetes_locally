PROJECTNAME=$(shell basename "$(PWD)")
BINARY_NAME=main.out

# ==================================================================================== #
# HELPERS
# ==================================================================================== #

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #

## run/api: run the application
.PHONY: run/api
run/api:
	go run ./main.go

## vendor: tidy and vendor dependencies
.PHONY: vendor
vendor:
	@echo 'Tidying and verifying module dependencies...'
	go mod tidy
	go mod verify
	@echo 'Vendoring dependencies...'
	go mod vendor

# ==================================================================================== #
# BUILD
# ==================================================================================== #

## build/api: build the application
.PHONY: build/api
build/api:
	@echo 'Building ...'
	CGO_ENABLED=0 GOOS=linux go build -o ./bin/${BINARY_NAME} ./main.go

# ==================================================================================== #
# PRODUCTION
# ==================================================================================== #

## production/minikube: build, push to local minikube and up
.PHONY: production/minikube
production/minikube:
	@echo 'Building ...'
	eval $(minikube docker-env)
	docker build --rm -t klaus-hello-world .

# 	kubectl get deployment klaus-hello-world
# 	kubectl delete deployment klaus-hello-world
# 	kubectl delete pod klaus-hello-world
# kubectl get services klaus-hello-world-service

	kubectl delete services klaus-hello-world-service

	kubectl apply -f k8s/deployment.yaml
	kubectl expose deployment klaus-hello-world --type=LoadBalancer --name=klaus-hello-world-service
# 	@echo 'pls enter password for correct set up tunnel...'
	minikube tunnel

# 	@echo 'pls go to dashboard in browser...'
# 	minikube dashboard

