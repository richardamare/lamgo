validate:
	@sam validate --lint

build: validate
	@go env -w CGO_ENABLED=0 GOOS=linux GOARCH=amd64
	@sam build -p

deploy: build
	@sam deploy --config-env dev