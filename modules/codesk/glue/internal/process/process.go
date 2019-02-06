package process

import (
    "fmt"
    "os"
    "os/exec"
    "sort"
    "syscall"
    "github.com/encodeering/wsl/glue/internal/config"
)

type Proxy interface {
    Exec (args []string) int
}

func NewProxy (binary string, config *config.Config) Proxy {
    return &proxy {
        binary: binary,
        config: config,
    }
}

type proxy struct {
    binary string
    config *config.Config
}

func (p *proxy) Exec (args []string) (code int) {
    p.environmentify ()

    cmd := exec.Command (p.binary, args...)
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

func (p *proxy) environmentify ()  () {
                vars := append ([]config.WslVar (nil), p.config.Command.Environment.Var...)
    sort.Slice (vars, func (i, j int) bool {
        return  vars[i].Key < vars[j].Key
    })

    wslenv := ""

    for _, wslvar := range vars {
        os.Setenv                       (wslvar.Key, wslvar.Value)

        wslenv += fmt.Sprintf ("%s/%s:", wslvar.Key, wslvar.Spec)
    }

    os.Setenv ("WSLENV", wslenv)
}
