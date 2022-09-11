linters:
	./entrypoint.sh linters

tests:
	./entrypoint.sh tests

validate:
	./entrypoint.sh validate

docker-validate:
	docker-compose build api_dev && docker-compose run --rm --service-ports api_dev validate 

shell:
	docker-compose build api_dev && docker-compose run --rm api_dev shell
