package test

import (
    "errors"
    "os"
    "path"
)

func init () {
    for {
        cwd, err := os.Getwd ()
        if err != nil {
            panic (err)
        }

        if _, err := os.Stat (path.Join (cwd, "main.go")); ! os.IsNotExist (err) {
            return
        }

        if cwd == "." {
            panic (errors.New ("couldn't find root package"))
        }

        if err := os.Chdir (".."); err != nil {
            panic (err)
        }
    }
}
