package config

import (
    "os"
    "os/user"
    "path/filepath"
    . "github.com/encodeering/wsl/glue/internal/util"
)

var Pwd string
var Uhd string

func init () {
    executable, err := os.Executable ()
    Die (err)

    path, err := filepath.EvalSymlinks (executable)
    Die (err)

    Pwd = filepath.Dir (path)

    user, err := user.Current ()
    Die (err)

    Uhd = user.HomeDir
}
