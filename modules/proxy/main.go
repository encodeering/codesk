package main

import (
    "github.com/encodeering/wsl/proxy/internal/cmd"
)

//go:generate tar -C glue -czvf ./internal/stick/res/glue.tar.gz .

func main() {
    cmd.Execute ()
}
