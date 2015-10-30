package main

import (
  "flag"
  "github.com/Sirupsen/logrus"
  "github.com/thiagoaugustus/shared/envconfig"
)

var (
  appenv = flag.String("config", ".yaml", "")
  debug	= flag.Bool("debug", false, "")
)

func main() {
  flag.Parse()
  
  // debug level if requested by user
  if *debug {
    logrus.SetLevel(logrus.DebugLevel)
  }
  
  // Load the configuration from yaml file
  env := envconfig.Load(*appenv)
}