# Create base for installing binaries
BIN = $(CURDIR)/bin
$(BIN):
	@mkdir -p $@
$(BIN)/golangci-lint: | $(BIN)
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.24.0
$(BIN)/gopherbadger: | $(BIN)
	GOBIN=$(BIN) go get github.com/jpoles1/gopherbadger

# Binaries to install
GOLANGCI-LINT = $(BIN)/golangci-lint
GOPHERBADGER = $(BIN)/gopherbadger

###############################################################################
###############################################################################
###############################################################################

all: clean static test build

# Build binaries
build:
	go build -a -o bin/particle main.go;

clean:
	-rm -r bin/
	-rm -r releases/

static: | $(GOLANGCI-LINT) $(GOPHERBADGER)
	$(GOLANGCI-LINT) run ./... --fix
	$(GOPHERBADGER) -md="README.md"

unit:
	go test ./... -cover -v

PLATFORMS := linux-amd64 linux-386 darwin-amd64 windows-amd64 windows-386
temp = $(subst -, ,$@)
os = $(word 1, $(temp))
arch = $(word 2, $(temp))
releases: $(PLATFORMS)
$(PLATFORMS):
	@mkdir -p releases; \
	CGO_ENABLED=0 GOOS=$(os) GOARCH=$(arch) go build -a -o bin/particle-$(os)-$(arch) main.go; \
	tar -C bin -cvzf releases/particle-$(os)-$(arch).tar.gz particle-$(os)-$(arch) particle-$(os)-$(arch); \
