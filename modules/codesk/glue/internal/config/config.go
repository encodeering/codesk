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
