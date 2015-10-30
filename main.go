package main

import (
  "flag"
  log "github.com/Sirupsen/logrus"
  "github.com/thiagoao/GoLang-Structure/shared/envconfig"
)

var (
  appenv = flag.String("config", "app.yaml", "")
  envmod	= flag.String("envmod", "dev", "")
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
  
  if err != nil {
    log.WithFields(log.Fields{
      "event": "Faild to get",
      "topic": "Config",
      "code" : 24,
    }).Fatal("Faild to get the configuration")    
  }
  
  
  //store_ := datastore.Load(env.EnvMode(*envmod))
}