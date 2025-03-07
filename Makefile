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
	cd $(PWD)/clients/cli/cmd/bookctl && \
	go build -o bookctl main.go && \
	mv bookctl $(PWD)/bin

openapiclient:
	openapi-generator generate -i docs/swagger/v1/openapi.yaml -g go -o ./clients/openapi --additional-properties=moduleName=github.com/vishalanarase/bookstore/openapi

test:
	go test -v ./...

docker:
	docker build -t bookstore-api -f build/api/Dockerfile .


## Kustomize build
kustiomizebuild: kustiomizebuildstaging kustiomizebuildproduction kustiomizebuilddev

kustiomizebuilddev:
	kustomize build deploy/overlays/dev > config/dev/dev.yaml

kustiomizebuildstaging:
	kustomize build deploy/overlays/staging > config/staging/staging.yaml

kustiomizebuildproduction:
	kustomize build deploy/overlays/production > config/production/production.yaml

## Kustomize apply
kustomizeapply: kustomizeapplydev kustomizeapplystaging kustomizeapplyproduction

kustomizeapplydev:
	kustomize build deploy/overlays/dev/ | kubectl apply -f -

kustomizeapplystaging:
	kustomize build deploy/overlays/staging/ | kubectl apply -f -

kustomizeapplyproduction:
	kustomize build deploy/overlays/production/ | kubectl apply -f -

## Kustomize delete
kustomizedelete: kustomizedeletedev kustomizedeletestaging kustomizedeleteproduction

kustomizedeletedev:
	kustomize build deploy/overlays/dev/ | kubectl delete -f -

kustomizedeletestaging:
	kustomize build deploy/overlays/staging/ | kubectl delete -f -

kustomizedeleteproduction:
	kustomize build deploy/overlays/production/ | kubectl delete -f -