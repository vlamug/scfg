# SCFG

**SCFG** is a simple config storage. It provides api to load config by cache.
Service uses PostgreSQL to store configurations and memory cache to optimize the speed of loading.

## API endpoints

It provides only one api endpoint.

### Get config 

**/get** - this endpoint loads config by cache.

*The example of request:*

```json
{
    "type": "Deploy.robot",
    "data": "stages"
}
```

*The example of response:*

```json
{
    "hosts": [
        "hdth.prod",
        "hdph.prod",
        "zhru.prod"
    ],
    "repo": "http://github.com/zigi/st",
    "user": "deployer",
    "migrations": "http://st.zigi.st.migrations.prod"
}
```

## Build and run

It is possible to up service using two ways:

### Locally

1. Install PosgreSQL
2. Download *goose* and run migrations specifying correct data to PosgreSQL:
```bash
$ go get -u github.com/pressly/goose/cmd/goose
$ cd migrations/
$ goose postgres "user=postgres password=root dbname=scfg sslmode=disable" up
```
3. Run command:
```bash
make run
```

### In docker environment
Run command:
```bash
$ make docker
```

It will up docker environment with PosgreSQL container 


Using these ways it will run on *9002* port and with database connections parameters, which are stored in *etc/database.json* file.

You can use own parameters. For this run application as follow for example:

```bash
go run main.go --api.listen-addr="http://localhost:9005" --database.config-path="db.json"
```
