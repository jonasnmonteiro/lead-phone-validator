APP_NAME=leadphone-validator
PORT?=3007

.PHONY: build run test clean docker-build docker-run web-build web-dev

web-build:
	cd web && npm install && npm run build

web-dev:
	cd web && npm run dev

build: web-build
	go build -ldflags="-w -s" -o bin/$(APP_NAME) .

run:
	go build -ldflags="-w -s" -o bin/$(APP_NAME) .
	./bin/$(APP_NAME)

test:
	go test -v ./...

clean:
	rm -rf bin/ web/dist/ web/node_modules/

docker-build:
	docker build -t $(APP_NAME):latest .

docker-run:
	docker run -p $(PORT):$(PORT) -e PORT=$(PORT) $(APP_NAME):latest
