run:
	go run cmd/main.go
	
createdb:
	docker run --name=todo-db -e POSTGRES_PASSWORD='qwerty' -p 5436:5432 -d postgres

migrateup:
	migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' up

migratedown: 
	migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' up

startdb:
	docker container start todo-db

stopdb:
	docker container stop todo-db

#signdb
#	docker exec -it todo-db /bin/bash
#	psql -U postgres


.PHONY: run createdb migrateup migratedown startdb stopdb