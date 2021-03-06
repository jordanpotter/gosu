# Gosu
Gosu is a massively distributed voice communication platform that focuses on performance, simplicity, and high availability. By leveraging the latest technologies, Gosu proves resilience to server instability and transient connectivity issues. And with relay servers all around the world, there is no other faster medium to transfer voice communication reliably.

**NOTICE:** Gosu is a work in progress. The future is very much unknown. However the project has a very strong mission statement -- to provide a powerful user-friendly and performant voice communication experience.

# Development Setup

## Etcd
All discovery and configuration for the cluster is managed by [etcd](github.com/coreos/etcd). During development, make sure to have at least one instance running for the servers to communicate with.

**NOTE:** Setting up a high availability Etcd cluster is left out of this setup guide, however is crucial in production. For your production deployments please follow [CoreOS's guide](coreos.com/etcd/docs/2.0.12/clustering.html) on the subject.

For simplicity, etcd can be run locally via [Docker](docker.com) during development

    docker run --name etcd -d -p 4001:4001 -p 7001:7001 microbox/etcd:0.4.9 -name gosu

### Discovery

We use etcd to keep track of where our servers are located. Run these to set the development defaults

    curl http://127.0.0.1:4001/v2/keys/addrs/auth   -XPOST -d value='{"ip": "127.0.0.1", "httpPort": 8080}'
    curl http://127.0.0.1:4001/v2/keys/addrs/api    -XPOST -d value='{"ip": "127.0.0.1", "httpPort": 8081, "pubPort": 9001}'
    curl http://127.0.0.1:4001/v2/keys/addrs/events -XPOST -d value='{"ip": "127.0.0.1", "httpPort": 8082, "subPort": 9002}'
    curl http://127.0.0.1:4001/v2/keys/addrs/relay  -XPOST -d value='{"ip": "127.0.0.1", "httpPort": 8083, "commsPort": 1337}'
    curl http://127.0.0.1:4001/v2/keys/addrs/postgres  -XPOST -d value='{"ip": "127.0.0.1", "dbPort": 5432}'

 If you decide to run a server on a different ip address or port, be sure to remove the old entry and add the new one for that server. Thorough instructions can be found [here](github.com/coreos/etcd/blob/master/Documentation/api.md).

### Config

The `conf` directory includes some cluster-wide configuration files to insert into our etcd service

    curl -L http://127.0.0.1:4001/v2/keys/conf/auth/token -XPUT --data-urlencode value@conf/auth/token.json
    curl -L http://127.0.0.1:4001/v2/keys/conf/postgres -XPUT --data-urlencode value@conf/db/postgres.json

If you wish to modify some configuration parameters in the `conf` directory, be sure to update etcd by running the above commands again.

## Postgres
Currently Gosu uses [Postgres](postgresql.org) as its backing data store. Postgres provides performance and data guarantees that resonate well with the high reliability mission Gosu is trying to achieve.

**NOTE:** Setting up a high availability Postgres cluster is left out of this setup guide, however is crucial in production. In a production setting, either follow [Postgres's guide](wiki.postgresql.org/wiki/Replication,_Clustering,_and_Connection_Pooling) or use a trusted third-party solution.

Postgres can be run locally via [Docker](docker.com) during development

    docker run --name postgres -d -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=password postgres:9.4.1

Likewise, we can use [Docker](docker.com) to enter psql

    docker run -it --rm --link postgres:postgres postgres:9.4.1 sh -c 'exec psql gosu -h "$POSTGRES_PORT_5432_TCP_ADDR" -p "$POSTGRES_PORT_5432_TCP_PORT" -U postgres'

Be sure to create the `gosu` database. Leveraging the above command, we can do this easily

    docker run -it --rm --link postgres:postgres postgres:9.4.1 sh -c 'exec psql -h "$POSTGRES_PORT_5432_TCP_ADDR" -p "$POSTGRES_PORT_5432_TCP_PORT" -U postgres -c "create database gosu"'
