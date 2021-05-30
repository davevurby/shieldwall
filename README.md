# Shield wall

[![Coverage Status](https://coveralls.io/repos/github/davevurby/shieldwall/badge.svg?branch=main)](https://coveralls.io/github/davevurby/shieldwall?branch=main)

**Shieldwall is still in progress, it is NOT READY to be used in production!!**

Shieldwall is GO package, that can handle policies, roles and identities in a single or multiple tenant/domain systems (called namespaces). Its goal is to handle the best performance with very simple queries onto your desired storage system.

As this is only a golang package, you have to create a server, that would handle requests as well. If you don't want to create your own server, you can use our own [open-source server solution](https://github.com/davevurby/shieldwall-server)

## Features

- [x] Create a role and restrict namespaces, to which the role could be assigned (with regex matching)
- [x] Manage identities (users, services, admins), their roles and namespaces
- [x] Create a policy for an identity or a role that is matching a namespace
- [x] Very SIMPLE way how to check if an identity is permitted to do an action (via single query)

## Multiple storage options

- [x] PostgreSQL
- [ ] MySQL / MariaDB
- [ ] CrockroachDB

## Caching options

- [ ] Redis caching
- [ ] Memcached caching

## Multiple interfaces options

- [x] Use as Golang dependency
- [x] [REST server](https://github.com/davevurby/shieldwall-server)
- [ ] gRPC server

## Getting started

Soon..
