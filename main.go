package main

import (
  "flag"
  log "github.com/Sirupsen/logrus"
  "github.com/thiagoaugustus/GoLang-Structure/shared/envconfig"
)

var (
  appenv = flag.String("config", "app.yaml", "")
  debug	= flag.Bool("debug", false, "")
)

func main() {
  flag.Parse()
  
  // debug level if requested by user
  if *debug {
    log.SetLevel(log.DebugLevel)
  }
  
  // Load the configuration from yaml file
  env, err := envconfig.Load(*appenv)
}