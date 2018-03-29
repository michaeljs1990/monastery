package config

import (
	"flag"
)

var (
	CQLServer   string
	CQLKeyspace string
)

func init() {
	flag.StringVar(&CQLServer, "CQLServer", "127.0.0.1", "Set server to connect to")
	flag.StringVar(&CQLKeyspace, "CQLKeyspace", "monastery", "Set keyspace to use")
}
