package store

import (
  "fmt"
  "errors"
  "strings"
  "github.com/jinzhu/gorm"
  //_ "github.com/lib/pq"
  //_ "github.com/go-sql-driver/mysql"
  log "github.com/Sirupsen/logrus"
  "github.com/thiagoao/GoLang-Structure/shared/envconfig"
  "github.com/thiagoao/GoLang-Structure/store/migration"
)

// Load database base configuration and return a store.Store
func Load(env envconfig.Env) store.Store {
  var driver, config, err := getConfig(env)
  
  if err != nil {
    log.WithFields(log.Fields{
      "event": "Faild to connect",
      "topic": "Store",
      "code" : 14,
    }).Fatal(string(err))
  }
  
  return New(driver, config)
}

func New(driver, config string) store.Store {
  return 
}

func Open(driver, config string) gorm.DB {
  db, err := gorm.Open(driver, config)
  if err != nil {
    log.Errorln(err)
    log.Fatalln("database connection failed")
  }
  
  db.DB().Ping()
  db.DB().SetMaxIdleConns(10)
  db.DB().SetMaxOpenConns(100)
  
}

fun setupDatabase(driver string, db *gorm.DB) error {
  
}

// getConfig return the driver, confinguration and an error
func getConfig(env envconfig.Env) (d, c string, error) {

  d := env.DB.Driver
  switch d {

    default:
    d, c, error := nil, nil, errors.New("Database: Driver not found")

    case "mysql" : 
    c := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", 
                     env.DB.User, env.DB.Pass, env.DB.host, env.DB.Port, env.DB.DBName)
    case "postgresql" :
    c := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=verify-full",
                     env.DB.User, env.DB.Pass, env.DB.host, env.DB.Port, env.DB.DBName)
  }

  return
}
