# Gosu
Gosu is a massively distributed voice communication platform that focuses on performance, simplicity, and high availability.

## Setup

### Etcd
All configuration for the cluster is managed by [etcd](github.com/coreos/etcd). During development, make sure to have at least one instance running for the servers to communicate with.

For simplicity, etcd can be run locally via [Docker](docker.com)

    docker run --name etcd -d -p 4001:4001 -p 7001:7001 microbox/etcd:0.4.9 -name gosu

For simplicity, the `conf` directory includes some cluster-wide configuration files to insert into our etcd service

    curl -L http://127.0.0.1:4001/v2/keys/authToken -XPUT --data-urlencode value@conf/authToken.json
    curl -L http://127.0.0.1:4001/v2/keys/mongo -XPUT --data-urlencode value@conf/mongo.json

If you wish to modify some configuration parameters in the `conf` directory, be sure to update etcd by running the above commands again.

### Mongo
Currently Gosu uses [MongoDB](mongodb.org) as its backing data store. MongoDB's data model and availability guarantees resonate well with what Gosu is trying to achieve, although this decision is by no means permanent.

MongoDB can be run locally via [Docker](docker.com)

    docker run --name mongo -d -p 27017:27017 mongo:3.0.2
