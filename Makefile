.PHONY: migrations testsetup test

devsetupdatabase:
	cd $(PWD)/migrations && \
	soda -c config/database.yaml create

testsetupdatabase:
	cd $(PWD)/migrations && \
	soda -e test -c config/database.yaml create

devrunmigrations:
	cd $(PWD)/migrations && \
	soda -e test -c config/database.yaml migrate up

testrunmigrations:
	cd $(PWD)/migrations && \
	soda -e test -c config/database.yaml migrate up

setupdatabase: devsetupdatabase testsetupdatabase

runmigrations: devrunmigrations testrunmigrations

testsetup: testsetupdatabase testrunmigrations

run:
	cd $(PWD)/cmd/api && \
	go run main.go

test:
	go test -v ./...