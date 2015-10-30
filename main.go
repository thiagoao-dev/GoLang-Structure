package main

import (
  "fmt"
  "flag"
  log "github.com/Sirupsen/logrus"
  "github.com/thiagoaugustus/GoLang-Structure/shared/envconfig"
)

var (
  appenv = flag.String("config", "ap.yaml", "")
  debug	= flag.Bool("debug", true, "")
)

func main() {
  flag.Parse()
  
  // debug level if requested by user
  if *debug {
    log.SetLevel(log.DebugLevel)
  }
  
  // Load the configuration from yaml file
  cfg, err := envconfig.Load(*appenv)
  
  if err != nil {
    log.WithFields(log.Fields{
      "event": "Faild to get",
      "topic": "Config",
      "key"  : 24,
    }).Fatal("Faild to get the configuration")    
  }
  
  fmt.Println(cfg.Dev.Server)
}