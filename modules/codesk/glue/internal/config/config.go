package config

import (
    "errors"
    "fmt"
    "regexp"
)

var splitter = struct {
   envvar *regexp.Regexp
}{
   envvar: regexp.MustCompile ("^([^=]+)=([wulp]*)=(.*?)$"),
}

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

type WslVar struct {
    Key   string
    Value string
    Spec  string
}

type Environment struct {
    Resolution Resolution `yaml:"resolution"`
    Var []WslVar
}

type Command struct {
    Environment Environment `yaml:"environment"`
}

type Config struct {
    Box Box `yaml:"box"`
    Command Command `yaml:"command"`
}

func DefaultWslVar (key, spec, value string) WslVar {
    return WslVar {
        Key   : key,
        Spec  : spec,
        Value : value,
    }
}

func DefaultConfig () (config *Config) {
    config = &Config {}
    config.Command.Environment.Resolution = "last"

    return
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

func ConvertResolution (item string) (resolution Resolution, err error) {
    if err = CheckResolution (Resolution (item)); err != nil {
        return
    }

    resolution = Resolution (item)

    return
}

func CheckEnvvar (envvar string) (err error) {
    if matches := splitter.envvar.MatchString (envvar); ! matches {
        return errors.New (fmt.Sprintf ("envvar value '%v' doesn't match pattern NAME=[wulp]=VALUE", envvar))
    }

    return
}

func ConvertEnvvar (envvar string) (wslvar WslVar, err error) {
    if err = CheckEnvvar (envvar); err != nil {
        return
    }

                            match := splitter.envvar.FindStringSubmatch (envvar)
    wslvar = DefaultWslVar (match[1], match[2], match[3])

    return
}

func ConvertEnvvars (envvars []string) (wslvars []WslVar, err error) {
    for _, e := range envvars {
        var wslvar WslVar
        if  wslvar, err = ConvertEnvvar (e); err != nil {
            return
        }

        wslvars = append (wslvars, wslvar)
    }

    return
}
