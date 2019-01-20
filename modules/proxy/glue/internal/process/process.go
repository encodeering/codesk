package process

import (
    "os"
    "os/exec"
    "syscall"
)

func Exec (binary string, args []string) (code int) {
    cmd := exec.Command (binary, args...)
    cmd.Stdin  = os.Stdin
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    code = 0

    if err := cmd.Run (); err != nil {
        code = 1

        if exiterr, ok := err.(*exec.ExitError); ok {
            code = exiterr.Sys ().(syscall.WaitStatus).ExitStatus ()
        }
    }

    return
}
