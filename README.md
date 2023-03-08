# ORDER SERVICE

### Tools

 - Docker
 - Go
 - JetStream
 - Postgres

### Quick start

Before start make an .env file or build a sample one form .env-example:

```shell
make build-env
```

To run the Postgres and JetStream server:

```shell
make build
```

To run migrations:

```shell
make migrate-up
```

To run the application:

```shell
make start
```



### OPTIONS

If you want to make a producer that will send orders on your JetStream server you can run this:

```shell
make prod
```

To reset migrations:

```shell
make migrate-reset
```
