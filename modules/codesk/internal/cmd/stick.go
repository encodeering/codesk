package cmd

import (
    "github.com/encodeering/wsl/codesk/internal/stick"
    "github.com/spf13/cobra"
)

var Target string

var gluestick = &cobra.Command {
    Use: `stick`,
    Short: `glues a binary that acts as a proxy between windows and linux over wsl.exe`,
    Run: func (_ *cobra.Command, args []string) {
        Die (stick.New (stick.StreamHandle (Target)).Process ())
    },
}

func init () {
    root.AddCommand (gluestick)

    gluestick.Flags ().StringVarP (&Target, "target", "t", "", "targets this linux binary, either full-qualified or short")
    gluestick.MarkFlagRequired ("target")
}
