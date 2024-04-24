export OUTPUT_PATH=$(shell pwd)/test-output.out
export COVERAGE_PATH=$(shell pwd)/coverage.out
export LICENSES_PATH=$(shell pwd)/licenses.csv

.PHONY: help
.DEFAULT_GOAL := help
help: ## Show this help.
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: build
build: ## Build the executable modules project.
	mkdir -p bin
	GOARCH=amd64 GOOS=linux go build -o bin ./...

.PHONY: lint
lint: ## Execute the lint command to validate the code.
	go run github.com/golangci/golangci-lint/cmd/golangci-lint@latest run --config .golangci.yaml

.PHONY: vulnerabilities
vulnerabilities:
	go run golang.org/x/vuln/cmd/govulncheck@latest ./...

.PHONY: tidy
tidy: ## Execute the tidy command to clean the project.
	go fmt ./...
	go mod tidy -v

.PHONY: licenses
licenses: ## Execute the commands to generate license report.
	go run github.com/google/go-licenses@latest csv ./... > $(LICENSES_PATH)

.PHONY: unit-test
unit-test: ## Execute the unit tests of the different modules of the project.
	go test -v -coverprofile=$(COVERAGE_PATH) ./...

.PHONY: clean
clean: ## Remove the executable modules of the project.
	rm -rf ./bin
	rm -f $(OUTPUT_PATH)
	rm -f $(COVERAGE_PATH)
	rm -f $(LICENSES_PATH)