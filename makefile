run: build run-without-build

start-db:
	./tools/start_postgresql_docker.sh
	sleep 5
	./tools/drop_and_create_db.sh

stop-db:
	./tools/stop_postgresql_docker.sh

build-docker-image:
	docker build -t simple-core-bank .
	./tools/remove_all_none_image.sh

build:
	go build -o simple-core-bank src/main.go

run-without-build:
	./simple-core-bank migrate-db
	./simple-core-bank serve-http-api
