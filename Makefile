run:
	go run ./cmd
mock:
	mockery --all --dir internal/adapter
test:
	go test ./... --cover
docker-sql:
	docker run --name mysql -p 6603:3306 -e MYSQL_ROOT_PASSWORD=admin -d mysql:latest
migrate-status:
	goose -dir files/migration mysql "root:admin@tcp(localhost:6603)/mysql?parseTime=true" status
migrate-up:
	goose -dir files/migration mysql "root:admin@tcp(localhost:6603)/mysql?parseTime=true" up
migrate-down:
	goose -dir files/migration mysql "root:admin@tcp(localhost:6603)/mysql?parseTime=true" down