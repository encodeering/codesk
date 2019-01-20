package cmd

import (
    "os"
    "github.com/spf13/cobra"
)

var root = &cobra.Command {
    Use: `proxy`,
    Short: `a tool for wsl proxy generation`,
    Run: func (_ *cobra.Command, _ []string) {},
}

func Execute () {
    Die (root.Execute ())
}

func Die (err error) {
    if err != nil {
        os.Exit (1)
    }
}
