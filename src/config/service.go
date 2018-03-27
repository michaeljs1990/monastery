package config

import (
	"flag"
	"os"
)

var (
	ServicePort string
)

func init() {
	flag.StringVar(&ServicePort, "ServicePort", os.Getenv("SERVICE_PORT"), "Set the port for the service to start on")
}
