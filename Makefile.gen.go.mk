# DO NOT EDIT. Generated with:
#
#    devctl
#
#    https://github.com/giantswarm/devctl/blob/bf7f386ac6a4e807dde959892df1369fee6d789f/pkg/gen/input/makefile/internal/file/Makefile.gen.go.mk.template
#

APPLICATION    := $(shell go list -m | cut -d '/' -f 3)
BUILDTIMESTAMP := $(shell date -u '+%FT%TZ')
GITSHA1        := $(shell git rev-parse --verify HEAD)
MODULE         := $(shell go list -m)
OS             := $(shell go env GOOS)
SOURCES        := $(shell find . -name '*.go')
VERSION        := $(shell architect project version)
ifeq ($(OS), linux)
EXTLDFLAGS := -static
endif
LDFLAGS        ?= -w -linkmode 'auto' -extldflags '$(EXTLDFLAGS)' \
  -X '$(shell go list -m)/pkg/project.buildTimestamp=${BUILDTIMESTAMP}' \
  -X '$(shell go list -m)/pkg/project.gitSHA=${GITSHA1}'

.DEFAULT_GOAL := build

##@ Go

.PHONY: build build-darwin build-darwin-64 build-linux build-linux-arm64 build-windows-amd64
build: $(APPLICATION) ## Builds a local binary.
	@echo "====> $@"
build-darwin: $(APPLICATION)-darwin ## Builds a local binary for darwin/amd64.
	@echo "====> $@"
build-darwin-arm64: $(APPLICATION)-darwin-arm64 ## Builds a local binary for darwin/arm64.
	@echo "====> $@"
build-linux: $(APPLICATION)-linux ## Builds a local binary for linux/amd64.
	@echo "====> $@"
build-linux-arm64: $(APPLICATION)-linux-arm64 ## Builds a local binary for linux/arm64.
	@echo "====> $@"
build-windows-amd64: $(APPLICATION)-windows-amd64.exe ## Builds a local binary for windows/amd64.
	@echo "====> $@"

$(APPLICATION): $(APPLICATION)-v$(VERSION)-$(OS)-amd64
	@echo "====> $@"
	cp -a $< $@

$(APPLICATION)-darwin: $(APPLICATION)-v$(VERSION)-darwin-amd64
	@echo "====> $@"
	cp -a $< $@

$(APPLICATION)-darwin-arm64: $(APPLICATION)-v$(VERSION)-darwin-arm64
	@echo "====> $@"
	cp -a $< $@

$(APPLICATION)-linux: $(APPLICATION)-v$(VERSION)-linux-amd64
	@echo "====> $@"
	cp -a $< $@

$(APPLICATION)-linux-arm64: $(APPLICATION)-v$(VERSION)-linux-arm64
	@echo "====> $@"
	cp -a $< $@

$(APPLICATION)-windows-amd64.exe: $(APPLICATION)-v$(VERSION)-windows-amd64.exe
	@echo "====> $@"
	cp -a $< $@

$(APPLICATION)-v$(VERSION)-%-amd64: $(SOURCES)
	@echo "====> $@"
	CGO_ENABLED=0 GOOS=$* GOARCH=amd64 go build -trimpath -ldflags "$(LDFLAGS)" -o $@ .

$(APPLICATION)-v$(VERSION)-%-arm64: $(SOURCES)
	@echo "====> $@"
	CGO_ENABLED=0 GOOS=$* GOARCH=arm64 go build -trimpath -ldflags "$(LDFLAGS)" -o $@ .

$(APPLICATION)-v$(VERSION)-windows-amd64.exe: $(SOURCES)
	@echo "====> $@"
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -trimpath -ldflags "$(LDFLAGS)" -o $@ .

.PHONY: install
install: ## Install the application.
	@echo "====> $@"
	go install -ldflags "$(LDFLAGS)" .

.PHONY: run
run: ## Runs go run main.go.
	@echo "====> $@"
	go run -ldflags "$(LDFLAGS)" -race .

.PHONY: clean
clean: ## Cleans the binary.
	@echo "====> $@"
	rm -f $(APPLICATION)*
	go clean

.PHONY: imports
imports: ## Runs goimports.
	@echo "====> $@"
	goimports -local $(MODULE) -w .

.PHONY: lint
lint: ## Runs golangci-lint.
	@echo "====> $@"
	golangci-lint run -E gosec -E goconst --timeout=15m ./...

.PHONY: fmt
fmt: ## Run go fmt against code.
	go fmt ./...

.PHONY: vet
vet: ## Run go vet against code.
	go vet ./...

.PHONY: nancy
nancy: ## Runs nancy (requires v1.0.37 or newer).
	@echo "====> $@"
	CGO_ENABLED=0 go list -json -deps ./... | nancy sleuth --skip-update-check --quiet --exclude-vulnerability-file ./.nancy-ignore --additional-exclude-vulnerability-files ./.nancy-ignore.generated

.PHONY: test
test: ## Runs go test with default values.
	@echo "====> $@"
	go test -ldflags "$(LDFLAGS)" -race ./...

.PHONY: build-docker
build-docker: build-linux ## Builds docker image to registry.
	@echo "====> $@"
	cp -a $(APPLICATION)-linux $(APPLICATION)
	docker build -t ${APPLICATION}:${VERSION} .
