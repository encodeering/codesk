package config

import (
    "io/ioutil"
    "os"
    "gopkg.in/yaml.v2"
)

func (e* Environment) UnmarshalYAML (unmarshal func (interface{}) error) (err error) {
    type   typeclone Environment
    raw := typeclone {
        Resolution : e.Resolution,
        Var        : e.Var,
    }

    if err = unmarshal (& raw); err != nil {
        return
    }

    *e = Environment (raw)

    return
}

func ObtainYaml (file string) (config *Config, err error) {
    config = &Config {}
    config.Command.Environment.Resolution = "last"

    if _, err = os.Stat (file); err != nil {
        err   = nil
        return
    }

    var b []byte
    if  b, err = ioutil.ReadFile (file); err != nil {
        return
    }

    err = yaml.Unmarshal (b, & config)

    return
}
