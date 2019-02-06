package main

import (
    "os"
    . "github.com/encodeering/wsl/glue/internal/config"
    . "github.com/encodeering/wsl/glue/internal/process"
)

var target string

func main () {
    os.Exit (
        NewProxy ("wsl.exe", Configuration ()).Exec (append ([]string{target}, os.Args[1:]...)),
    )
}
