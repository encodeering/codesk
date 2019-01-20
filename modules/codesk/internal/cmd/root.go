package cmd

import (
    "os"
    "github.com/spf13/cobra"
)

var root = &cobra.Command {
    Use: `codesk`,
    Short: `a tool for windows/linux desk operations using wsl`,
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
