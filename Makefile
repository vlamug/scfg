.PHONY: docker
docker:
	@echo " >> up docker env..."
	docker-compose stop
	docker-compose rm -v -f
	docker-compose build
	docker-compose up -d
	sleep 3
	docker-compose ps

docker-rebuild:
	@echo " >> rebuild scfg container..."
	docker-compose rm -f scfg
	docker-compose build --no-cache scfg
	docker-compose up -d --force-recreate scfg

run:
	@echo " >> run app..."
	@GOOS=$(GOOS) GOARCH=$(GOARCH) go run main.go
