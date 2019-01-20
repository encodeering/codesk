package cmd

import (
    "bytes"
    "os"
    "os/exec"
    "strings"
    "github.com/encodeering/wsl/proxy/internal/stick"
    "github.com/spf13/cobra"
)

var Target string

var gluestick = &cobra.Command {
    Use: `stick`,
    Short: `glues a binary that acts as a proxy between windows and linux over wsl.exe`,
    Run: func (_ *cobra.Command, args []string) {
        proxy := stick.New (func (script []byte, fs []byte) (err error) {
            cmd := exec.Command ("wsl.exe", "bash", "-c", strings.Replace (string (script), "$", "\\$", -1), "--", Target)
            cmd.Stdin = bytes.NewReader (fs)
            cmd.Stdout = os.Stdout
            cmd.Stderr = os.Stderr

            return cmd.Run ()
        })

        Die (proxy.Process ())
    },
}

func init () {
    root.AddCommand (gluestick)

    gluestick.Flags ().StringVarP (&Target, "target", "t", "", "targets this linux binary, either full-qualified or short")
    gluestick.MarkFlagRequired ("target")
}
