.PHONY: migrations testmigrations test

migrations:
	cd $(PWD)/db && \
	soda -c config/database.yaml create && \
	soda -c config/database.yaml migrate up

testmigrations:
	cd $(PWD)/db && \
	soda -e test -c config/database.yaml create && \
	soda -e test -c config/database.yaml migrate up

run:
	cd $(PWD)/cmd/api && \
	go run main.go

test:
	go test -v ./...