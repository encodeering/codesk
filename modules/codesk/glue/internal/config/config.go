package config

import (
    "errors"
    "fmt"
    "regexp"
)

type Resolution string

type Distribution struct {
    Name string `yaml:"name,omitempty"`
}

type User struct {
    Name string `yaml:"name,omitempty"`
}

type Box struct {
    Distribution Distribution `yaml:"distribution"`
    User User `yaml:"user"`
}

type Environment struct {
    Resolution Resolution `yaml:"resolution"`
    Var []string `yaml:"var"`
}

type Command struct {
    Environment Environment `yaml:"environment"`
}

type Config struct {
    Box Box `yaml:"box"`
    Command Command `yaml:"command"`
}

func CheckResolution (resolution Resolution) error {
    if resolution == "first"  ||
       resolution == "parent" ||
       resolution == "self"   ||
       resolution == "last" {

        return nil
    }

    return errors.New (fmt.Sprintf ("resolution value '%v' is unknown", resolution))
}

func CheckEnvvar (envvar string) (err error) {
    var matches bool
    if  matches, err = regexp.Match ("^([^=]+)=([wulp]*)=(.*?)$", []byte (envvar)); err != nil {
        return
    }

    if ! matches {
        return errors.New (fmt.Sprintf ("envvar value '%v' doesn't match pattern NAME=[wulp]=VALUE", envvar))
    }

    return
}

func CheckEnvvars (envvars []string) (err error) {
    for _, e := range envvars {
        if err = CheckEnvvar (e); err != nil {
            return
        }
    }

    return
}
