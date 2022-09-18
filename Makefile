.PHONY: linters test validate docker-validate shell
linters:
	./entrypoint.sh linters

test:
	./entrypoint.sh tests

validate:
	./entrypoint.sh validate

docker-validate:
	docker-compose build api_dev && docker-compose run --rm --service-ports api_dev validate 

shell:
	docker-compose build api_dev && docker-compose run -v "$PWD:/app" --rm api_dev shell
