package config

import (
    "io/ioutil"
    "os"
    "gopkg.in/yaml.v2"
)

func (e* Environment) UnmarshalYAML (unmarshal func (interface{}) error) (err error) {
    raw := struct {
        Resolution Resolution `yaml:"resolution"`
        Var []string `yaml:"var"`
    }{
        Resolution : e.Resolution,
    }

    if err = unmarshal (& raw); err != nil {
        return
    }

    if err = CheckResolution (raw.Resolution); err != nil {
        return
    }

    var wslvars []WslVar
    if  wslvars, err = ConvertEnvvars (raw.Var); err != nil {
        return
    }

    *e = Environment {
        Resolution: raw.Resolution,
        Var: wslvars,
    }

    return
}

func ObtainYaml (file string) (config *Config, err error) {
    config = DefaultConfig ()

    if _, err = os.Stat (file); err != nil {
        err   = nil
        return
    }

    var b []byte
    if  b, err = ioutil.ReadFile (file); err != nil {
        return
    }

    if err = yaml.Unmarshal (b, & config); err != nil {
        config = nil
    }

    return
}
