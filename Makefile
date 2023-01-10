
protoc:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative pkg/pb/*.proto --go-grpc_opt=require_unimplemented_servers=false

proto:
	protoc pkg/pb/*.proto --go-grpc_out=.

start:
	go run main.go serve --config file://config.json

postgres:
	docker run --name postgres15 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:15-alpine

createdb:
	docker exec -it postgres15 createdb --username=root --owner=root invitation_auth

dropdb:
	docker exec -it postgres15 dropdb invitation_auth

migrateup:
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/invitation_auth?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/invitation_auth?sslmode=disable" -verbose down