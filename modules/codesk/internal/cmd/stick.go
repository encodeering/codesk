package cmd

import (
    "os"
    "github.com/encodeering/wsl/codesk/internal/stick"
    "github.com/spf13/cobra"
)

var Target string
var Out string

var gluestick = &cobra.Command {
    Use: `stick`,
    Short: `glues a binary that acts as a proxy between windows and linux over wsl.exe`,
    Run: func (_ *cobra.Command, args []string) {
        Die (stick.New (handle ()).Process ())
    },
}

func handle () stick.Handle {
    if Out == "-" {
        return stick.IoHandle (Target, os.Stdout)
    }

    if Out == "" {
       Out = Target + ".exe"
    }

    out, err := os.Create (Out)
    Die (err)

    return stick.IoHandle (Target, out)
}

func init () {
    root.AddCommand (gluestick)

    gluestick.Flags ().StringVarP (&Target, "target", "t", "", "targets this linux binary, either full-qualified or short")
    gluestick.Flags ().StringVarP (&Out, "out", "o", "", "writes binary out to stdout or a file")
    gluestick.MarkFlagRequired ("target")
}
