# API example

## Development

### Linters and tests

```sh
make linters
make test

# or with one command
make validate

# or with one command inside docker
make docker-validate
```
### Database migrations:

Install migrate tool:
```sh
go get -tags 'sqlite3' -u github.com/golang-migrate/migrate/v4/cmd/migrate
go install -tags 'sqlite3' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

Create new migration
```sh
./bin/makemigrations.sh
```

Apply all migrations
```sh
./bin/migrate.sh up
```
