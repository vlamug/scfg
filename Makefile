.PHONY: docker
docker:
	@echo " >> up docker env..."
	@docker-compose stop
	@docker-compose rm -v -f
	@docker-compose build
	@docker-compose up -d
	@sleep 3
	@docker-compose ps

.PHONY: docker-rebuild
docker-rebuild:
	@echo " >> rebuild scfg container..."
	@docker-compose rm -f scfg
	@docker-compose build --no-cache scfg
	@docker-compose up -d --force-recreate scfg

.PHONY: run
run:
	@echo " >> run app..."
	@GOOS=$(GOOS) GOARCH=$(GOARCH) go run main.go

build:
	@echo " >> build app..."
	@GOOS=$(GOOS) GOARCH=$(GOARCH) go build main.go

.PHONY: test
test:
	@echo " >> run tests..."
	@go list ./... | grep -v /vendor/ | xargs go test

.PHONY: fmt
fmt:
	@echo " >> running fmt tool..."
	@git diff origin/master --name-only --diff-filter=ACMRTUXB | grep '.go' | xargs gofmt -w && git status

.PHONY: lint
lint:
	@echo " >> running lint tool..."
	@git diff origin/master --name-only --diff-filter=ACMRTUXB | grep '.go' | xargs golint

.PHONY: vet
vet:
	@echo " >> running vet tool..."
	@git diff origin/master --name-only --diff-filter=ACMRTUXB | grep '.go' | xargs go vet
