# simple web app in golang to connect postgresql database

## purpose:
    > connect to postgresql database
    > work with migration / golang-migrate
    > improve our project structure

### how to run and build project
> first export your postgresql_url to env
> export POSTGRESQL_URL='postgres://postgres:yourpassword@localhost:5432/example?sslmode=disable'
> then run migration file
```
    migrate -database ${POSTGRESQL_URL} -path db/migration up
    make run
```