.PHONY: start
start: # When you first clone the repo, start here!
	go mod tidy

.PHONY: serve
serve: # Runs the server for local development
	CONFIG_PATH=config/config_local.hcl go run cmd/goprotoarchetype-api/main.go

.PHONY: proto/lint
proto/lint: # Lint protobufs
	buf lint .

.PHONY: proto/generate
proto/generate: # Compile protobufs
	rm -rf ./gen
	buf generate

.PHONY: test
test: # Run tests
	go test ./...
