stop:
	docker-compose stop

up:stop
	docker-compose up

.PHONY: stop up
