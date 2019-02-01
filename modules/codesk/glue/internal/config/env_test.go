package config

import (
    "os"
    "strings"
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestOSPrefix (t *testing.T) {
    unsetenv ("CODESK")

    setenv ("CODESK_ANSWER", "", "42")
    setenv ("CODESK_LIFE", "pu", "42")

    config, err := ObtainOS ("CODESK")
    assert.NoError (t, err)

    assert.Equal (t, []WslVar{ DefaultWslVar ("ANSWER", "", "42"), DefaultWslVar ("LIFE", "pu", "42") }, config.Command.Environment.Var)
}

func TestOSEnvInvalid (t *testing.T) {
    unsetenv ("CODESK")

    setenv ("CODESK_LIFE", "x", "42")

    config, err := ObtainOS ("CODESK")

    assert.EqualError (t, err, "envvar value 'LIFE=x=42' doesn't match pattern NAME=[wulp]=VALUE")
    assert.Nil (t, config)
}

func setenv (key, flag, value string) {
    os.Setenv (key, flag + "=" + value)
}

func unsetenv (prefix string) {
    for _, e := range os.Environ () {
        if strings.HasPrefix (e, "CODESK") {
            os.Unsetenv (strings.Split (e, "=")[0])
        }
    }
}
