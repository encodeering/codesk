package config

import (
    "os"
    "os/user"
    "path/filepath"
    "strings"
    . "github.com/encodeering/wsl/glue/internal/util"
)

func Configuration () *Config {
    arg0 := os.Args[0]

    base := filepath.Base (arg0)
    ext  := filepath.Ext  (arg0)

    homeyml, err := ObtainYaml (filepath.Join (Uhd (), "codesk.yml"))
    Die (err)

    selfyml, err := ObtainYaml (filepath.Join (Pwd (), strings.TrimSuffix (base, ext) + ".yml"))
    Die (err)

    env, err := ObtainOS ("CODESK")
    Die (err)

    return Combine (homeyml, selfyml, env)
}

func Pwd () string {
    executable, err := os.Executable ()
    Die (err)

    path, err := filepath.EvalSymlinks (executable)
    Die (err)

    return filepath.Dir (path)
}

func Uhd () string {
    user, err := user.Current ()
    Die (err)

    return user.HomeDir
}
