package envconfig

import (
  //"gopkg.in/yaml.v2"
)

type Env map[string]string

// Load reads the environment file and reads variables in "key:value" yaml format.
// Then it read the system environment variables. It returns the combined
// results in a key value map.
func Load(filepath string) Env {
  return Env
}