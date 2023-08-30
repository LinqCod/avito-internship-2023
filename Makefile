dc:
	sudo docker-compose up --remove-orphans --build

swag-gen:
	swag init -g cmd/main.go --parseDependency --parseInternal

.PHONY: dc swag-gen