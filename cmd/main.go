package main

import (
	"fmt"
	"flag"
	"os"

	"github.com/spf13/viper"
)

type application struct {
	cfg config
}
type config struct {
	port int
	env string
	host string
}

func main() {
	var cfg config

	flag.StringVar(&cfg.env, "env", "dev", "Environment(dev|sit|prod)")
	flag.Parse()
	if !validEnv(cfg.env) {
		fmt.Println("Invalid environment: ", cfg.env)
		os.Exit(1)
	}
	viper.SetConfigName("dev")
	viper.AddConfigPath("internal/config")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	fmt.Println("Hello, World!", viper.Get("config.dbCreds.port"))
}
