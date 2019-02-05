package config

import (
    "os"
    "os/user"
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

func TestConfigUserHomeDirectory (t *testing.T) {
    user, err := user.Current ()
    assert.NoError (t, err)

    assert.Equal (t, user.HomeDir, Uhd ())
}
