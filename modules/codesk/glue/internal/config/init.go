package config

import (
    "os"
    "path/filepath"
    "strings"
    . "github.com/encodeering/wsl/glue/internal/util"
)

func Configuration () *Config {
    arg0 := os.Args[0]

    base := filepath.Base (arg0)
    ext  := filepath.Ext  (arg0)

    codeskyml, err := ObtainYaml (filepath.Join (Pwd (), "codesk.yml"))
    Die (err)

    selfyml, err := ObtainYaml (filepath.Join (Pwd (), strings.TrimSuffix (base, ext) + ".yml"))
    Die (err)

    env, err := ObtainOS ("CODESK")
    Die (err)

    return Combine (codeskyml, selfyml, env)
}

func Pwd () string {
    executable, err := os.Executable ()
    Die (err)

    path, err := filepath.EvalSymlinks (executable)
    Die (err)

    return filepath.Dir (path)
}
