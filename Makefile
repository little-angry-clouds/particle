# Create base for installing binaries
BIN = $(CURDIR)/bin
$(BIN):
	@mkdir -p $@
$(BIN)/golangci-lint: | $(BIN)
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.24.0
$(BIN)/gopherbadger: | $(BIN)
	GOBIN=$(BIN) go get github.com/jpoles1/gopherbadger
$(BIN)/goreleaser: | $(BIN)
	curl -sfL https://install.goreleaser.com/github.com/goreleaser/goreleaser.sh | sh

# Binaries to install
GOLANGCI-LINT = $(BIN)/golangci-lint
GOPHERBADGER = $(BIN)/gopherbadger
GORELEASER = $(BIN)/goreleaser

###############################################################################
###############################################################################
###############################################################################

all: clean static unit build

# Build binaries
build: | $(GORELEASER)
	$(GORELEASER) build --rm-dist --skip-validate

clean:
	rm -r bin/
	rm -r dist/

static: | $(GOLANGCI-LINT) $(GOPHERBADGER)
	$(GOLANGCI-LINT) run ./... --fix
	$(GOPHERBADGER) -md="README.md"

unit:
	go test ./... -cover

releases:
	$(GORELEASER) release
