package main

import (
    "os"
)

func main() {
    die (nil)
}

func die (err error) {
    if err != nil {
        os.Exit (1)
    }
}
