# gin-petclinic-api
[TOC]

## TODO
    godep

## Installation

    $ go get github.com/robfig/config
    $ go get github.com/Sirupsen/logrus
    $ go get github.com/jinzhu/gorm
    $ go get github.com/mattn/go-sqlite3
    $ go get github.com/gin-gonic/gin
    $ go get github.com/fvbock/endless

## Initial Setup

    $ go run ./tools/migrate.go
    $ go run app.go
    $ curl -X POST -d "name=cat&age=1" http://localhost:3001/pets
    $ curl -X POST -d "first_name=john&last_name=doe&address=chuo-ku&city=tokyo" http://localhost:3001/owners
    $ curl -X PUT -d "owner_id=1" http://localhost:3001/pets/1

    $ curl http://localhost:3001/pets | jq .
    $ curl http://localhost:3001/owners | jq .

## Auto Build

    $ go get github.com/codegangsta/gin
    $ gin -i run app.go
    $ curl -X GET http://localhost:3000/pets | jq .
    $ curl -X GET http://localhost:3000/owners | jq .

## License

MIT
