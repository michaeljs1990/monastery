package config

import (
	"flag"
)

var (
	ServicePort string
)

func init() {
	flag.StringVar(&ServicePort, "ServicePort", "8100", "Set the port for the service to start on")
}
