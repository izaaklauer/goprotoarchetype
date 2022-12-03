.PHONY: start
start: # When you first clone the repo, start here!
	go mod tidy

.PHONY: serve
serve: # Runs the server for local development
	go run cmd/goprotoarchetype-api/main.go config/config_local.hcl

.PHONY: proto/lint
proto/lint: # Lint protobufs
	buf lint .

.PHONY: proto/generate
proto/compile: # Compile protobufs
	rm -rf ./gen
	buf generate

.PHONY: test
test: # Run tests
	go test ./...