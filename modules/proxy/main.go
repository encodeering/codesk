package main

import (
    "os"
)

//go:generate tar -C glue -czvf ./internal/stick/res/glue.tar.gz .

func main() {
    die (nil)
}

func die (err error) {
    if err != nil {
        os.Exit (1)
    }
}
