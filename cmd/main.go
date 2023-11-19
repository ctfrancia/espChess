package main

import (
	"fmt"
	"flag"
	"log/slog"
	"os"
	"net/http"
	"time"

	"github.com/spf13/viper"
)

type application struct {
	config config
	logger *slog.Logger
}
type config struct {
	env string
	port int
	host string
}

func main() {
	var cfg config
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// get the environment from command line
	flag.StringVar(&cfg.env, "env", "dev", "Environment(dev|sit|prod)")
	flag.Parse()
	if !validEnv(cfg.env) {
		fmt.Println("Invalid environment: ", cfg.env)
		os.Exit(1)
	}
	// set config from file
	viper.SetConfigName(cfg.env)
	viper.AddConfigPath("internal/config")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	// set application configuration based on environment
	cfg.port = viper.GetInt("config.application.port")
	cfg.host = viper.GetString("config.application.host")
	cfg.env = viper.GetString("config.environment")
	app := &application{
		config: cfg,
		logger: logger,
	}

	srv := &http.Server{
		Addr: fmt.Sprintf(":%d", app.config.port),
		Handler: app.routes(),
		IdleTimeout: time.Minute,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	// start server with logger
	msg := fmt.Sprintf("Starting %s server on %s", app.config.env, srv.Addr) 
	logger.Info(msg)
	err = srv.ListenAndServe()
	logger.Error(err.Error())
	os.Exit(1)
}
