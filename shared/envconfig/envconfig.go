package envconfig

import (
  "io/ioutil"
  "gopkg.in/yaml.v2"
  log "github.com/Sirupsen/logrus"
)

type Env struct {
  Dev struct {
    Server   struct {
      Host   string
      Port   uint
      Prefix string
    },
    DB struct {
      Host   string
      Driver string
      DBName string
      Port   uint
      User   string
      Pass   string      
    }
  }
}

// Load reads the environment file and reads variables in "key:value" yaml format.
// Then it read the system environment variables. It returns the combined
// results in a key value map.
func Load(filepath string) Env, err {
  
  cfgFile, err := ioutil.ReadFile(path)
  
  if err != nil {
    log.WithFields(log.Fields{
      "event": "Faild to read"
      "topic": "Config"
      "key"  : 16
    }).Fatal("Faild to read the configuration file")
    return nil, err
  }
  
  appCfg := map[string]*Env{} 
  
  if err = yaml.Unmarshal(cfgFile, appCfg); err != nil {
    log.WithFields(log.Fields{
      "event": "Faild to set"
      "topic": "Config"
      "key"  : 27
    }).Fatal("Faild to set the configuration file")
    return nil, err
  }
  
  return appCfg, nil
  
}