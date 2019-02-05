package config

import (
    "os"
    "os/user"
    "path/filepath"
    . "github.com/encodeering/wsl/glue/internal/util"
)

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
