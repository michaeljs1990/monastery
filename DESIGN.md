Design Doc
==========

Outline of high level design goals for this project as well as API outline.

## Goals High Level

 * Http interface that is abstract of backend storage solution (In Progress)
 * Support for the following storage solutions
   * Ceph
   * HDFS
   * S3 (In Progress)
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
