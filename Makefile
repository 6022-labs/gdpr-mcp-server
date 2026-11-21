.PHONY: build test integration-test lint fmt mocks hooks help

# build compiles every Go package under src/.
build:
	go build ./src/...

# test runs the unit-test packages; the $(...) form propagates a failing exit code.
test:
	go test $$(go list ./tests/... | grep '_unit_tests')

# integration-test runs the integration packages; they read the GDPR JSON under data/, no Docker needed.
integration-test:
	go test $$(go list ./tests/... | grep '_integration_tests')

# lint requires golangci-lint v2 on PATH (go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@latest).
lint:
	golangci-lint run ./...

# fmt runs golines via golangci-lint's built-in formatter; no standalone golines needed.
fmt:
	golangci-lint fmt

# mocks requires the deepsearch-mockgen CLI (github.com/FournyP/deepsearch-mockgen-cli).
mocks:
	deepsearch-mockgen -S ./src -O ./tests -A -P

# hooks installs pinned lefthook + gitleaks into GOPATH/bin and wires the git hooks.
hooks:
	go install github.com/evilmartians/lefthook/v2@v2.2.2
	go install github.com/zricethezav/gitleaks/v8@v8.30.1
	$$(go env GOPATH)/bin/lefthook install

help:
	@echo "build       - go build ./src/..."
	@echo "test        - go test unit-test packages"
	@echo "integration-test - go test integration packages (reads data/)"
	@echo "lint        - golangci-lint run ./... (needs golangci-lint v2)"
	@echo "fmt         - golangci-lint fmt, 120-char cap"
	@echo "mocks       - regenerate mocks (needs deepsearch-mockgen)"
	@echo "hooks       - install git hooks (format + secret scan on commit, build/lint/test on push)"
