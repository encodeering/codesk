package config

import (
    "io/ioutil"
    "os"
    "gopkg.in/yaml.v2"
)

func (e* Environment) UnmarshalYAML (unmarshal func (interface{}) error) (err error) {
    raw := struct {
        Resolution string `yaml:"resolution"`
        Var []string `yaml:"var"`
    }{
        Resolution : string (e.Resolution),
    }

    if err = unmarshal (& raw); err != nil {
        return
    }

    var resolution Resolution
    if  resolution, err = ConvertResolution (raw.Resolution); err != nil {
        return
    }

    var wslvars []WslVar
    if  wslvars, err = ConvertEnvvars (raw.Var); err != nil {
        return
    }

    *e = Environment {
        Resolution: resolution,
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
