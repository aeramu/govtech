# How to Run
## Pre-requisite
### Golang
### Docker
### Goose DB Migration Tool
https://github.com/pressly/goose
## Unit Test
```bash
make test
```
## Run
```bash
make docker-sql //run mysql docker on port 6603
make migrate-up
make run
```