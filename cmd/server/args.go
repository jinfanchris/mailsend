package main

import (
	"github.com/spf13/pflag"
)

type Args struct {
	Version bool
	ConfigF string
	CertF   string
	KeyF    string
}

func ParseArgs() Args {
	args := Args{}

	pflag.BoolVarP(&args.Version, "version", "v", false, "Print version")
	pflag.StringVarP(&args.ConfigF, "config", "c", "runtime/config.toml", "Path to config file")
	pflag.StringVarP(&args.CertF, "cert", "C", "runtime/cert.pem", "Path to TLS certificate")
	pflag.StringVarP(&args.KeyF, "key", "K", "runtime/key.pem", "Path to TLS key")

	pflag.Parse()
	return args
}
