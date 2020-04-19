# note: call scripts from /scripts

#input
NAME ?= zacks
SWAGGER_SPEC ?= swagger.yml
YAML_FILE := $(shell pwd)/api/${SWAGGER_SPEC}
CLIENT_PATH := $(shell pwd)/internal/app/server/gen/client
CLIENT_PKG_SUFFIX ?= client

SERVER_PATH := $(shell pwd)/internal/app/server/gen/server

#static
TEMPLATES_DIR := ../go-swagger-template/templates

#dynamic
MODELS_PKG_SUFFIX ?= $(CLIENT_PKG_SUFFIX)models

gen-server:
	rm -rf $(SERVER_PATH)/*
	mkdir -p $(SERVER_PATH)
	docker run --rm -it -u `id -u $(USER)` \
		-e GOPATH=$(HOME)/go:/go \
		-v $(HOME):$(HOME) \
		-w $(shell pwd) quay.io/goswagger/swagger:v0.20.1 generate server \
		-f $(YAML_FILE) \
		-A $(NAME) \
		--template-dir ${TEMPLATES_DIR} \
		--exclude-main \
		-t $(SERVER_PATH)

gen-client:
	rm -rf $(CLIENT_PATH)/*
	mkdir -p $(CLIENT_PATH)
	docker run --rm -it -u `id -u $(USER)` \
		-e GOPATH=$(HOME)/go:/go \
		-v $(HOME):$(HOME) \
		-w $(shell pwd) quay.io/goswagger/swagger:v0.20.1 generate client \
		-f $(YAML_FILE) \
		-A $(NAME) \
		--template-dir ${TEMPLATES_DIR} \
		-c $(NAME)$(CLIENT_PKG_SUFFIX) \
		-m $(NAME)$(MODELS_PKG_SUFFIX) \
		-t $(CLIENT_PATH)

docker-build:
	docker build -t docker.io/iadolgov/zacks -f deployments/Dockerfile.multistage .

docker-deploy:
	docker push iadolgov/zacks

docker-run:
	docker run -p 8080:8080 \
        -e SERVER_PORT=8080 \
        -e CACHE_ENABLED=true \
        -e DBCACHE_ENABLED=true \
        -e DBCACHE_MONGODB_URL=mongodb://localhost:27017 \
        -e DBCACHE_MONGODB_DATABASE_NAME=zacks \
        -e FETCHER_ENABLED=true \
        -e FETCHER_TIMEOUT_SECONDS=10 \
        -e REFRESHER_ENABLED=true \
        -e REFRESHER_RESCAN_SECONDS=3600 \
        docker.io/iadolgov/zacks:latest

run:
	go run ./internal/app/server/.

up:
	cd api; docker-compose up

prod-build:
	docker build -t s1.local:32000/iadolgov/zacks:latest -f deployments/Dockerfile.multistage .

prod-push:
	docker push s1.local:32000/iadolgov/zacks:latest

prod-deploy:
	helm --kubeconfig ~/.kube/s1 upgrade --wait --install --force zacks-deployment ./deployments/k8s/s1/
