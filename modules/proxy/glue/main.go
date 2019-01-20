package main

import (
    "os"
    "github.com/encodeering/wsl/glue/internal/process"
)

var target string

func main () {
    os.Exit (
        process.Exec ("wsl.exe", append ([]string{target}, os.Args[1:]...)),
    )
}
