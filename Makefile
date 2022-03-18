BIN="./bin"
SRC=$(shell find . -name "*.go")

ifneq (, $(shell test -s ${BIN}/golangci-lint  || echo "exists" ))
$(warning "could not find golangci-lint in bin directory, run: curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh")
endif

ifeq (, $(shell which richgo))
$(warning "could not find richgo in path, run: go get -u github.com/kyoh86/richgo")
endif

ifeq (, $(shell which swag))
$(warning "could not find swag in path, run: go get -u github.com/swaggo/swag/cmd/swag")
endif

.PHONY: help fmt lint test swagger run run-swag build install clean

default: help

help: ## Show this message
	$(info )
	$(info Usage:  make COMMAND)
	$(info )
	$(info Go Rest Clean Boilerplate)
	$(info )
	$(info Commands:)

	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-10s\033[0m %s\n", $$1, $$2}'

all: fmt lint test swagger build ## Runs fmt, lint, test, swagger and build

fmt: ## Checks format of code
	$(info )
	$(info ******************** checking formatting ********************)
	$(info )

	@test -z $(shell gofmt -l $(SRC)) || (gofmt -d $(SRC); exit 1)

lint: ## Runs lint tools
	$(info )
	$(info ******************** running lint tools ********************)
	$(info )

	$(BIN)/golangci-lint run -v

test: install  ## Runs tests
	$(info )
	$(info ******************** running tests ********************)
	$(info )

	richgo test -v ./...

coverage: install  ## Runs tests
	$(info )
	$(info ******************** running coverage ********************)
	$(info )

	go test ./...  -coverpkg=./... -coverprofile ./coverage.out
	go tool cover -func ./coverage.out

swagger:  ## Generates swagger api docs
	$(info )
	$(info ******************** generating swagger docs ********************)
	$(info )

	swag fmt
	swag init -g cmd/server/server.go

run:  ## Runs api server
	$(info )
	$(info ******************** running go-rest-clean-boilerplate ********************)
	$(info )

	go run main.go

run-swag: swagger  ## Generates swagger api docs and runs api server
	$(info )
	$(info ******************** running go-rest-clean-boilerplate ********************)
	$(info )

	go run main.go

build: install ## Builds api server to build directory
	$(info )
	$(info ******************** building go-rest-clean-boilerplate ********************)
	$(info )

	go build -o build

install: ## Installs dependencies of code
	$(info )
	$(info ******************** downloading dependencies ********************)
	$(info )

	go get -v ./...

clean: ## Cleans directory
	rm -rf $(BIN)
