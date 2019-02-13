package stick

import (
    "bytes"
    "os"
    "os/exec"
    "strings"
)

func StreamHandle (target string) Handle {
    return func (script []byte, fs []byte) (err error) {
        cmd := exec.Command ("wsl.exe", "bash", "-c", strings.Replace (string (script), "$", "\\$", -1), "--", target)
        cmd.Stdin = bytes.NewReader (fs)
        cmd.Stdout = os.Stdout
        cmd.Stderr = os.Stderr

        return cmd.Run ()
    }
}
