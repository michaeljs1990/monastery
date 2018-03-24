Design Doc
==========

Outline of high level design goals for this project as well as API outline.

## Goals High Level

 * Http interface that is abstract of backend storage solution
 * Support for the following storage solutions
   * Ceph
   * HDFS
   * S3
   * In Memory
   * Posix
 * The frontend will scale independently of the underlying storage solution
 * Configurable cache via groupcache
 * Support for multiple service discovery backends (etc, zookeeper, consul)
 * Builtin support for the following repository types
   * Debian
   * Pypi
   * Docker
 * Token and LDAP based auth
 * Optional ability to encrypt data before saving to backing file store

## High Level API

This is a basic high level API that I will flesh out more as I add support for
the repository types above which will likely require special APIs.

 * POST /relic/upload/${name}
 * GET /relic/${name}
