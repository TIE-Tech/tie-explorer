## tie-explorer

`Disclaimer: only for technical exchange learning, not for commercial use`

This code is divided into frontend, backend. All features are for the Polygon-Edge blockchain,
Not exactly what Ethereum chains do.

backend code used `Beego` framework. 

frontend code used  `Vue` framework.

data is stored by `Postgresql`.

backend features:

- Parse block information
- Parse transaction information & logs
- Collecting addresses Information
- Collect daily transaction data
- Collect basic information of chain

frontend features:

- Homepage
- block list & block details
- transaction list & transaction details
- account details

### Run

#### run backend

```shell
go run main.go
或
go build -o indexer main.go
./indexer
```

#### run frontend
```shell
cd web
npm install
npm run serve
```


### Docker

#### Docker-compose (recommended)
```shell
docker-compose up
```

#### install manually
step 1
```shell
docker pull postgres:13.6

docker run --name postgres -e POSTGRES_PASSWORD='' -e POSTGRES_USER='postgres' -e POSTGRES_HOST_AUTH_METHOD='trust' -p 5432:5432 -d postgres:13.6
```

step 2 `(optional)`
```shell
docker exec -it postgres /bin/bash
psql -U postgres
CREATE DATABASE explorer;
```

step 3
```shell
# don't forget modify conf/app.conf

docker build . -f docker/server/Dockerfile -t server

docker run --name server -p 8081:8081 -d server
```

step 4
```shell
docker build . -f docker/web/Dockerfile -t web

docker run --name web -p 8080:8080 -d
```
