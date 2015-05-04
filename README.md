# Gosu
Gosu is a massively distributed voice communication platform that focuses on performance, simplicity, and high availability.

## Etcd
All configuration for the cluster is managed by etcd. During development, make sure to have at least one instance running for the servers to communicate with.

For simplicity, etcd can be run locally via Docker

    docker run --name etcd -d -p 4001:4001 -p 7001:7001 microbox/etcd:0.4.9 -name gosu

## Mongo
Currently Gosu uses MongoDB as its backing data store. MongoDB's data model and availability guarantees resonate well with what Gosu is trying to achieve, although this decision is by no means permanent.

MongoDB can be run locally via Docker

    docker run --name mongo -d -p 27017:27017 mongo:3.0.2
