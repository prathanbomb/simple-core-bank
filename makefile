run:
	go build -o simple-core-bank src/main.go
	./simple-core-bank migrate-db
	./simple-core-bank serve-http-api

start-db:
	./tools/start_postgresql_docker.sh
	sleep 5
	./tools/drop_and_create_db.sh

stop-db:
	./tools/stop_postgresql_docker.sh
