APP_NAME=app
DB_NAME=storage

run:
	@echo "RUN postgres"
	docker run -d --name $(DB_NAME) --rm postgres:10.5

build:
	@echo "BUILD and RUN app"
	go mod download && go build -o $(APP_NAME) .
	./app

clean:
	rm $(APP_NAME)
	docker stop $(DB_NAME)

