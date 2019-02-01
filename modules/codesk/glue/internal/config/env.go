package config

import (
    "os"
    "strings"
)

func ObtainOS (prefix string) (config *Config, err error) {
    lookup :=  prefix + "_"

    config = DefaultConfig ()

    for _, e := range os.Environ () {
        if ! strings.HasPrefix (e, lookup) {
            continue
        }

        config.Command.Environment.Var = append (config.Command.Environment.Var, strings.Replace (e, lookup, "", 1))
    }

    if err = CheckEnvvars (config.Command.Environment.Var); err != nil {
        config = nil
    }

    return
}
