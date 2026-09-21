APP_NAME=leadphone-validator
PORT?=3007

.PHONY: build run test clean docker-build docker-run

build:
	go build -ldflags="-w -s" -o bin/$(APP_NAME) .

run: build
	./bin/$(APP_NAME)

test:
	go test -v ./...

clean:
	rm -rf bin/

docker-build:
	docker build -t $(APP_NAME):latest .

docker-run:
	docker run -p $(PORT):$(PORT) -e PORT=$(PORT) $(APP_NAME):latest
