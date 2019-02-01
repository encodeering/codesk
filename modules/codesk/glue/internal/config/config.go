package config

import (
    "errors"
    "fmt"
    "regexp"
)

type Resolution string

type Distribution struct {
    Name string
}

type User struct {
    Name string
}

type Box struct {
    Distribution Distribution
    User User
}

type Environment struct {
    Resolution Resolution
    Var []string
}

type Command struct {
    Environment Environment
}

type Config struct {
    Box Box
    Command Command
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
