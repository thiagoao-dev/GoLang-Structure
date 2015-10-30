package envconfig

import (
  "io/ioutil"
  "gopkg.in/yaml.v2"
  log "github.com/Sirupsen/logrus"
)

type Env struct {
  Dev struct {
    Server   struct {
      Host   string `yaml:"host"`
      Port   uint   `yaml:"port"`
      Prefix string `yaml:"prefix"`
    } `yaml:"server"`
    DB struct {
      Host   string `yaml:"host"`
      Driver string `yaml:"driver"`
      DBName string `yaml:"dbname"`
      Port   uint   `yaml:"port"`
      User   string `yaml:"user"`
      Pass   string `yaml:"pass"`
    } `yaml:"db"`
  } `yaml:"dev"`
  Prod struct {
    Server   struct {
      Host   string `yaml:"host"`
      Port   uint   `yaml:"port"`
      Prefix string `yaml:"prefix"`
    } `yaml:"server"`
    DB struct {
      Host   string `yaml:"host"`
      Driver string `yaml:"driver"`
      DBName string `yaml:"dbname"`
      Port   uint   `yaml:"port"`
      User   string `yaml:"user"`
      Pass   string `yaml:"pass"`
    } `yaml:"db"`
  } `yaml:"prod"`
}

// Load reads the environment file and reads variables in "key:value" yaml format.
// Then it read the system environment variables. It returns the combined
// results in a key value map.
func Load(filepath string) (Env, error) {
  
  cfgFile, err := ioutil.ReadFile(filepath)
  
  if err != nil {
    log.WithFields(log.Fields{
      "event": "Faild to read",
      "topic": "Config",
      "key"  : 16,
    }).Fatal("Faild to read the configuration file")
    return nil, err
  }
  
  appCfg := map[string]*Env{} 
  
  if err = yaml.Unmarshal(cfgFile, appCfg); err != nil {
    log.WithFields(log.Fields{
      "event": "Faild to set",
      "topic": "Config",
      "key"  : 27,
    }).Fatal("Faild to set the configuration file")
    return nil, err
  }
  
  return appCfg, nil
  
}