package config

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestYamlNotExist (t *testing.T) {
    config, err := ObtainYaml ("./test/fixture/yaml/_____.yml")
    assert.NoError (t, err)

    assert.Equal (t, DefaultConfig (), config)
}

func TestYamlReadFull (t *testing.T) {
    config, err := ObtainYaml ("./test/fixture/example-full.yml")
    assert.NoError (t, err)

    assert.Equal (t, "encodeering", config.Box.User.Name)
    assert.Equal (t, "debian", config.Box.Distribution.Name)

    assert.Equal (t, []string{"ANSWER==42", "LIFE=pu=42"}, config.Command.Environment.Var)
    assert.Equal (t, Resolution("parent"), config.Command.Environment.Resolution)
}

func TestYamlReadEnvNoResolution (t *testing.T) {
    config, err := ObtainYaml ("./test/fixture/example-env-no-resolution.yml")
    assert.NoError (t, err)

    assert.Equal (t, "encodeering", config.Box.User.Name)
    assert.Equal (t, "debian", config.Box.Distribution.Name)

    assert.Equal (t, []string{"ANSWER==42", "LIFE=pu=42"}, config.Command.Environment.Var)
    assert.Equal (t, Resolution("last"), config.Command.Environment.Resolution)
}

func TestYamlReadEnvResolutionInvalid (t *testing.T) {
    config, err := ObtainYaml ("./test/fixture/example-env-invalid-resolution.yml")
    assert.EqualError (t, err, "resolution value 'whatever' is unknown")
    assert.Nil (t, config)
}

func TestYamlReadEnvNoVar (t *testing.T) {
    config, err := ObtainYaml ("./test/fixture/example-env-no-var.yml")
    assert.NoError (t, err)

    assert.Equal (t, "encodeering", config.Box.User.Name)
    assert.Equal (t, "debian", config.Box.Distribution.Name)

    assert.Equal (t, []string(nil), config.Command.Environment.Var)
    assert.Equal (t, Resolution("parent"), config.Command.Environment.Resolution)
}

func TestYamlReadNothing (t *testing.T) {
    config, err := ObtainYaml ("./test/fixture/example-nothing.yml")
    assert.NoError (t, err)

    assert.Equal (t, "", config.Box.User.Name)
    assert.Equal (t, "", config.Box.Distribution.Name)

    assert.Equal (t, []string(nil), config.Command.Environment.Var)
    assert.Equal (t, Resolution("last"), config.Command.Environment.Resolution)
}
