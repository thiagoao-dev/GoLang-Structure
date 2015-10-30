package main

import (
  "flag"
  log "github.com/Sirupsen/logrus"
  "github.com/thiagoaugustus/GoLang-Structure/shared/envconfig"
)

var (
  appenv = flag.String("config", ".yaml", "")
  debug	= flag.Bool("debug", false, "")
)

func main() {
  flag.Parse()
  
  // debug level if requested by user
  if *debug {
    log.SetLevel(logrus.DebugLevel)
  }
  
  // Load the configuration from yaml file
  env := envconfig.Load(*appenv)
}