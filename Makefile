service-up:
	sudo docker-compose up --remove-orphans --build

service-down:
	sudo docker-compose down

swag-gen:
	swag init -g cmd/main.go --parseDependency --parseInternal

.PHONY: service-up service-down swag-gen