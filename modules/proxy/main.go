package main

import (
    "bytes"
    "errors"
    "os"
    "os/exec"
    "strings"
    "github.com/encodeering/wsl/proxy/internal/stick"
)

//go:generate tar -C glue -czvf ./internal/stick/res/glue.tar.gz .

func main() {
    if len (os.Args) < 1 {
        die (errors.New ("too less arguments"))
    }

    proxy := stick.New (func (script []byte, fs []byte) (err error) {
        cmd := exec.Command ("wsl.exe", "bash", "-c", strings.Replace (string (script), "$", "\\$", -1), "--", os.Args[1])
        cmd.Stdin = bytes.NewReader (fs)
        cmd.Stdout = os.Stdout
        cmd.Stderr = os.Stderr

        return cmd.Run ()
    })

    die (proxy.Process ())
}

func die (err error) {
    if err != nil {
        os.Exit (1)
    }
}
