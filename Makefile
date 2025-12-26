BIN_DIR := ./bin

GO      := go
GOFLAGS := -v
LDFLAGS :=

API_NAME := api
CLI_NAME := cli

API_CMD_DIR := ./cmd/api
CLI_CMD_DIR := ./cmd/cli

API_BIN := $(BIN_DIR)/$(API_NAME).exe
CLI_BIN := $(BIN_DIR)/$(CLI_NAME).exe

.PHONY: all build build-api build-cli run run-api run-cli tidy clean

all: build

build: build-api build-cli

build-api:
	@echo "==> build api"
	$(GO) build $(GOFLAGS) -ldflags="$(LDFLAGS)" -o $(API_BIN) $(API_CMD_DIR)

build-cli:
	@echo "==> build cli"
	$(GO) build $(GOFLAGS) -ldflags="$(LDFLAGS)" -o $(CLI_BIN) $(CLI_CMD_DIR)

run: run-api

run-api:
	@echo "==> run api"
	$(GO) run $(API_CMD_DIR)

tidy:
	@echo "==> tidy"
	$(GO) mod tidy

clean:
	@echo "==> clean"
	rm -rf $(BIN_DIR)
