.PHONY: migrations

migrations:
	cd $(PWD)/db && \
	soda -c config/database.yaml create && \
	soda -c config/database.yaml migrate up

run:
	cd $(PWD)/cmd/api && \
	go run main.go