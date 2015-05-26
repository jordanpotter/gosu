# Gosu
Gosu is a massively distributed voice communication platform that focuses on performance, simplicity, and high availability.

# Setup

## Etcd
All discovery and configuration for the cluster is managed by [etcd](github.com/coreos/etcd). During development, make sure to have at least one instance running for the servers to communicate with.

For simplicity, etcd can be run locally via [Docker](docker.com)

    docker run --name etcd -d -p 4001:4001 -p 7001:7001 microbox/etcd:0.4.9 -name gosu

### Discovery

We use etcd to keep track of where our servers are located. Run these to set the development defaults

    curl http://127.0.0.1:4001/v2/keys/addrs/auth   -XPOST -d value=127.0.0.1:8080
    curl http://127.0.0.1:4001/v2/keys/addrs/api    -XPOST -d value=127.0.0.1:8081
    curl http://127.0.0.1:4001/v2/keys/addrs/events -XPOST -d value=127.0.0.1:8082
    curl http://127.0.0.1:4001/v2/keys/addrs/relay  -XPOST -d value=127.0.0.1:8083
    curl http://127.0.0.1:4001/v2/keys/addrs/mongo  -XPOST -d value=127.0.0.1:27017

 If you decide to run a server on a different ip address or port, be sure to remove the old entry and add the new one for that server. Thorough instructions can be found [here](https://github.com/coreos/etcd/blob/master/Documentation/api.md).

### Config

The `conf` directory includes some cluster-wide configuration files to insert into our etcd service

    curl -L http://127.0.0.1:4001/v2/keys/conf/mongo -XPUT --data-urlencode value@conf/mongo.json
    curl -L http://127.0.0.1:4001/v2/keys/conf/auth/token -XPUT --data-urlencode value@conf/authToken.json

If you wish to modify some configuration parameters in the `conf` directory, be sure to update etcd by running the above commands again.

## Mongo
Currently Gosu uses [MongoDB](mongodb.org) as its backing data store. MongoDB's data model and availability guarantees resonate well with what Gosu is trying to achieve, although this decision is by no means permanent.

MongoDB can be run locally via [Docker](docker.com)

    docker run --name mongo -d -p 27017:27017 mongo:3.0.2

Likewise, we can use [Docker](docker.com) to enter the mongo shell

    docker run -it --rm --link mongo:mongo mongo:3.0.2 sh -c 'exec mongo "$MONGO_PORT_27017_TCP_ADDR:$MONGO_PORT_27017_TCP_PORT/gosu"'

# TODO

## Postgres
Currently Gosu uses [Postgres](postgresql.org) as its backing data store.

Postgres can be run locally via [Docker](docker.com)

    docker run --name postgres -d -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=password postgres:9.4.1

Likewise, we can use [Docker](docker.com) to enter psql

    docker run -it --rm --link postgres:postgres postgres:9.4.1 sh -c 'exec psql -h "$POSTGRES_PORT_5432_TCP_ADDR" -p "$POSTGRES_PORT_5432_TCP_PORT" -U postgres'
