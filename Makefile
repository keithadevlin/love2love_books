.PHONY: models modelgen

GIT_COMMIT_SHORT := $(shell git rev-parse --short HEAD)

ifneq ($(CODEBUILD_RESOLVED_SOURCE_VERSION),)
	GIT_COMMIT_SHORT := $(shell echo $(CODEBUILD_RESOLVED_SOURCE_VERSION) | head -c 7)
endif

all: deps check-go-fmt lint test build

.PHONY: deps-init
deps-init:
	rm -f go.mod go.sum
	@GO111MODULE=on go mod init
	@GO111MODULE=on go mod tidy

.PHONY: update-deps
update-deps:
	@GO111MODULE=on go mod tidy

.PHONY: deps
deps:
	@GO111MODULE=on go mod download

modelgen:
	@mkdir -p pkg/apimodel/generated
	@swagger-codegen generate -Dmodels -DpackageName=generated -l go -o pkg/apimodel/generated -i specs/swagger_definition/API_1.0.0.yaml

test: unit_test integration_test

mocks:
	@GO111MODULE=on go generate ./...

unit_test:
	@GO111MODULE=on go test ./... -v -tags=unit -timeout 30s -count=1 -p=1 -race

integration_test:
	@GO111MODULE=on go test ./... -v -tags=integration -timeout 30s -count=1 -p=1 -race

clean:
	rm -rf ./bin
	set -e && for pkg in $$(ls pkg/lambda); do \
		rm -rf pkg/lambda/$$pkg/bin; \
	done

build: clean
	@echo "GitHash: ${GIT_COMMIT_SHORT}"
	mkdir -p bin
	@GO111MODULE=on
	set -e && for pkg in $$(ls pkg/lambda); do \
		echo "\nbuilding: $$pkg\n"; \
		mkdir -p pkg/lambda/$$pkg/bin; \
		GO111MODULE=on GOOS=linux CGO_ENABLED=0 go build -ldflags "-X main.GitCommit=${GIT_COMMIT_SHORT}" -o ./pkg/lambda/$$pkg/bin/$(notdir $$pkg) ./pkg/lambda/$$pkg/cmd; \
		zip -qj ./bin/$$pkg.zip ./pkg/lambda/$$pkg/bin/*; \
	done

lint:
	command -v golangci-lint || (cd /usr/local ; wget -O - -q https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s latest)
	@GO111MODULE=on golangci-lint run

check-go-fmt:
	@if [ -n "$$(gofmt -d $$(find . -name '*.go' -not -path './vendor/*'))" ]; then \
		>&2 echo "The .go sources aren't formatted. Please format them with 'make reformat-go-code'."; \
		exit 1; \
	fi

reformat-go-code:
	gofmt -s -w .

.PHONY: reset-db
reset-db:				## Recreate the postgres container and run
	docker-compose rm -fs postgresql
	docker-compose up -d postgresql
	docker-compose up flyway
