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
	cd $(PWD)/clients/cli/cmd && \
	go build -o cli main.go && \
	mv cli ../../../bin/cli
	
openapiclien:
	openapi-generator generate -i docs/swagger/v1/openapi.yaml -g go -o ./clients/openapi --additional-properties=moduleName=github.com/vishalanarase/bookstore/openapiclient 

test:
	go test -v ./...

docker:
	docker build -t bookstore-api -f build/api/Dockerfile .