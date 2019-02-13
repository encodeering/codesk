package stick

import (
    "bytes"
    "io"
    "os"
    "os/exec"
    "strings"
)

func IoHandle (target string, writer io.WriteCloser) Handle {
    return func (script []byte, fs []byte) (err error) {
        defer writer.Close ()

        cmd := exec.Command ("wsl.exe", "bash", "-c", strings.Replace (string (script), "$", "\\$", -1), "--", target)
        cmd.Stdin = bytes.NewReader (fs)
        cmd.Stdout = writer
        cmd.Stderr = os.Stderr

        return cmd.Run ()
    }
}
