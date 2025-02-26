.PHONY: migrations testsetup test

devsetupdatabase:
	cd $(PWD)/migrations && \
	soda -e development -c config/database.yaml create

testsetupdatabase:
	cd $(PWD)/migrations && \
	soda -e test -c config/database.yaml create

devrunmigrations:
	cd $(PWD)/migrations && \
	soda -e development -c config/database.yaml migrate up

testrunmigrations:
	cd $(PWD)/migrations && \
	soda -e test -c config/database.yaml migrate up

setupdatabase: devsetupdatabase testsetupdatabase

runmigrations: devrunmigrations testrunmigrations

testsetup: testsetupdatabase testrunmigrations

run:
	cd $(PWD)/cmd/api && \
	go run main.go

cli:
	go build -o bin/cli client/cli/main.go
	
openapiclien:
	openapi-generator generate -i docs/swagger/v1/openapi.yaml -g go -o ./openapiclient

test:
	go test -v ./...

docker:
	docker build -t bookstore-api -f build/api/Dockerfile .