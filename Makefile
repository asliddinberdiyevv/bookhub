.PHONY: build up down run stop logs

run:
	docker-compose up -d 

build:
	docker-compose build 

logs:
	docker-compose logs -f

stop:
	docker-compose down

start: stop run build logs



