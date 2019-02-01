package config

import (
    "os"
    "strings"
)

func ObtainOS (prefix string) (config *Config, err error) {
    lookup :=  prefix + "_"

    config = DefaultConfig ()

    candidates := make ([]string, 0)

    for _, e := range os.Environ () {
        if ! strings.HasPrefix (e, lookup) {
            continue
        }

        candidates = append (candidates, strings.Replace (e, lookup, "", 1))
    }

    if  config.Command.Environment.Var, err = ConvertEnvvars (candidates); err != nil {
        config = nil
    }

    return
}
