.PHONY: help
.DEFAULT_GOAL := help

help: ## Show this help.
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

build: ## Build the executable modules project.
	mkdir -p bin
	GOARCH=amd64 GOOS=linux go build -o bin ./...

test: build ## Execute the test of the different modules of the project.
	echo 'Bootstrap output file' > $(OUTPUT_PATH)
	go test ./... -coverprofile=bin/coverage.out

clean: ## Remove the executable modules of the project.
	rm -rf ./bin

sonar: test ## Run the sonar scanner to generate quality report.
	sonar-scanner -Dsonar.projectVersion="$(version)" -Dsonar.login="$(token)"

licenses: ## Execute the commands to generate license report.
	go install github.com/google/go-licenses@latest
	go-licenses csv ./...