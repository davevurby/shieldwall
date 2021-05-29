# Shield wall

[![codecov](https://codecov.io/gh/davevurby/shieldwall/branch/main/graph/badge.svg?token=SKF2DP364R)](https://codecov.io/gh/davevurby/shieldwall)

**Shield wall is still in progress, it is REALLY NOT READY to be used in production!!**

Shield wall is a package handling policies, roles, identities and tenants (domains). Its goal is to handle RBAC in multitenant systems out of the box with a very simple queries, that allows Shield wall remain lightweight even with millions of users

It can be used as a dependency for any Golang project or as a standalone server with several interfaces like REST or gRPC (soon).

## Features

- [x] Creating roles and restrict tenants, to which it could be assigned (with regex matching)
- [x] Managing identities (users, services, admins), their roles and domains
- [x] Creating policy for an identity or a role matching a namespace
- [x] Very SIMPLE way how to check if an identity is permitted to do an action (via single query)

## Multiple storage options

- [x] PostgreSQL
- [ ] MySQL / MariaDB
- [ ] CrockroachDB
- [ ] Redis

## Multiple interfaces options

- [x] Use as Golang dependency
- [ ] REST server
- [ ] gRPC server

## Getting started

Soon..
