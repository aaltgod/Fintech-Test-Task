APP_NAME=app
DB_NAME=storage

run:
	@echo "====== RUN postgres ======"
	docker run -d --name $(DB_NAME) --rm -p 5432:5432 postgres:10.5

build:
	@echo "====== BUILD and RUN app ======"
	go mod download && go build -o ./.bin/$(APP_NAME) cmd/api/main.go
	./.bin/$(APP_NAME)

clean:
	docker stop $(DB_NAME)
	rm ./.bin/$(APP_NAME)

