package config

import (
    "os"
    "path/filepath"
    "testing"
    "github.com/stretchr/testify/assert"
    _ "github.com/encodeering/wsl/glue/test/cwd"
)

func TestConfigProcessWorkingDirectory (t *testing.T) {
    executable, err := os.Executable ()
    assert.NoError (t, err)

    resolved, err := filepath.EvalSymlinks (executable)
    assert.NoError (t, err)

    assert.Equal (t, filepath.Dir (resolved), Pwd ())
}
