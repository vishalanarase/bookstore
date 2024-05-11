.PHONY: migrations

migrations:
	cd $(PWD)/db && \
	soda -c config/database.yaml create && \
	soda -c config/database.yaml migrate up

run: migrations
	go run cmd/api/main.go