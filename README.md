# Gravity archer

In this game you will try to kill the other archers in the space while taking into the consideration the force of gravity of surrounding planets!

## Developer guide

In order to quickly start developing you should do following things on your machine:

1. Install [Golang](https://golang.org/doc/install) on your machine (at least 1.15)
1. Install [Nodejs](https://nodejs.org/en/download/) (at least v12.18.3)
1. Install [Docker](https://docs.docker.com/get-docker/) (at least 17.05)
1. Install [Buffalo](https://gobuffalo.io/en/): `go get -u -v -tags sqlite github.com/gobuffalo/buffalo/buffalo`
1. Run `docker-compose up -d` to run all dependencies (like SQL server).
1. Run `buffalo setup` to prepare your environment
1. Run `buffalo test` to ensure that all tests are passing
1. Run `buffalo dev` to run a daemon that will watch your source code, rebuild and launch it automatically. The dev instance will be served under [http://127.0.0.1:3000](http://127.0.0.1:3000) address.

### Dependencies

#### MySQL server

Gravity archer depends on instance of MySQL server. There are two options of satisfying this dependency.

1. Your own instance

   In case you already have your own server instance set up, you can simply pass following information via environment variables:

   ```
   ARCHER_DATABASE=<DATABASE_NAME>
   ARCHER_DATABASE_USER=<USER>
   ARCHER_DATABASE_PASSWORD=<PASSWORD>
   ARCHER_DATABASE_HOST=<HOST_OF_MYSQL_SERVER>
   ARCHER_DATABASE_PORT=<PORT_OF_MYSQL_SERVER>
   ```

1. Instance launched by docker.

   You can easily launch MySQL server instance to be used for the development by running:

   ```
   docker-compose up -d
   ```

   At this moment it requires port 8135 to be available on your machine. In case you'd be facing any MySQL related issues you can inspect the logs of the container by running:

   ```
   docker logs archer_mysql_1
   ```

##### Creating required tables and databases

In case Archer gives you some errors related to (lack of) existence of some SQL tables - those will happen after you launch MySQL server for the first time, or do some changes in the schema. Try running:

```
buffalo pop migrate
```

More useful database related commands:

- `buffalo pop drop -a` - causes Buffalo to drop all tables for all environments
- `buffalo pop create -a` - causes Buffalo to create all tables for all environments
- `buffalo pop drop -e development` - causes Buffalo to drop all tables for environment `development`
- `buffalo pop create -e development` - causes Buffalo to create all tables for environment `development`
- `buffalo pop --help` - prints help on topics regarding database operations.
