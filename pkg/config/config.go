package config

import (
	"fmt"
	"log"
	"os"

	"github.com/vrischmann/envconfig"
)

// Env See: https://godoc.org/github.com/vrischmann/envconfig
var Env struct {
	Log struct {
		Level string `envconfig:"optional,default=info"`
	}
	HTTP struct {
		Port int `envconfig:"optional,default=3000"`
	}
	Key struct {
		Private string `envconfig:"optional,default=keys/private.pem"`
		Public  string `envconfig:"optional,default=keys/public.pem"`
	}
}

// Commit is the git commit set by ldflags
var Commit string

// Init will parse environment variables into the Env struct
func Init(commit string) {
	// this removes timestamp prefixes from logs
	log.SetFlags(0)

	Commit = commit

	if err := envconfig.Init(&Env); err != nil {
		fmt.Printf("%+v", err)
	}

	// check for Keys existence
	if _, err := os.Stat(Env.Key.Private); os.IsNotExist(err) {
		fmt.Printf("Error: Private key (%v) does not exist", Env.Key.Private)
		os.Exit(1)
	}
	if _, err := os.Stat(Env.Key.Public); os.IsNotExist(err) {
		fmt.Printf("Error: Private key (%v) does not exist", Env.Key.Public)
		os.Exit(1)
	}
}
