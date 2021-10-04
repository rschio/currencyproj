.PHONY: all
all: docker-up build-server build-gateway

.PHONY: build-server
build-server:
	cd currency/cmd/server && go build -o ../../../bin/server

.PHONY: build-gateway
build-gateway:
	cd currency/cmd/gateway && go build -o ../../../bin/gateway

.PHONY: docker-up
docker-up:
	docker-compose up -d

.PHONY:
clean:
	rm -rf ./bin
	docker-compose down

.PHONY:
destroy:
	sudo rm -rf ./schema/mysql/data/
