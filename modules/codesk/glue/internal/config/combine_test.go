package config

import (
    "testing"
    "github.com/stretchr/testify/assert"
    _ "github.com/encodeering/wsl/glue/test/cwd"
)

func TestCombineNothing (t *testing.T) {
    assert.Equal (t, DefaultConfig (), Combine ())
}

func TestCombineBox (t *testing.T) {
    actual := Combine (ReadCombineConfig (t, "box.yml"))

    assert.Equal (t, "debian",      actual.Box.Distribution.Name)
    assert.Equal (t, "encodeering", actual.Box.User.Name)
}

func TestCombineCommandParent (t *testing.T) {
    resolution := "parent"

    actual := Combine (WithResolution (ReadCombineConfig (t, "command-answer-life.yml"), resolution))

    assert.Equal (t, Resolution (resolution), actual.Command.Environment.Resolution)
    assert.Empty (t,                          actual.Command.Environment.Var)
}

func TestCombineCommandSelf (t *testing.T) {
    resolution := "self"

            actual := Combine (WithResolution (ReadCombineConfig (t, "command-answer-life.yml"), resolution))
    vars := actual.Command.Environment.Var

    assert.Equal    (t, Resolution (resolution), actual.Command.Environment.Resolution)
    assert.Contains (t, vars, DefaultWslVar ("ANSWER", "", "42"))
    assert.Contains (t, vars, DefaultWslVar ("LIFE", "pu", "42"))
}

func TestCombineCommandFirst (t *testing.T) {
    resolution := "first"

    life :=  WithResolution (ReadCombineConfig (t, "command-answer-life.yml"), "self")
    john :=  WithResolution (ReadCombineConfig (t, "command-answer-john.yml"), resolution)

            actual := Combine (life, john)
    vars := actual.Command.Environment.Var

    assert.Equal    (t, Resolution (resolution), actual.Command.Environment.Resolution)
    assert.Len      (t, vars, 3)
    assert.Contains (t, vars, DefaultWslVar ("ANSWER", "", "42"))
    assert.Contains (t, vars, DefaultWslVar ("LIFE", "pu", "42"))
    assert.Contains (t, vars, DefaultWslVar ("JOHN", "pu", "42"))
}

func TestCombineCommandLast (t *testing.T) {
    resolution := "last"

    life :=  WithResolution (ReadCombineConfig (t, "command-answer-life.yml"), "self")
    john :=  WithResolution (ReadCombineConfig (t, "command-answer-john.yml"), resolution)

            actual := Combine (life, john)
    vars := actual.Command.Environment.Var

    assert.Equal    (t, Resolution (resolution), actual.Command.Environment.Resolution)
    assert.Len      (t, vars, 3)
    assert.Contains (t, vars, DefaultWslVar ("ANSWER", "", "24"))
    assert.Contains (t, vars, DefaultWslVar ("LIFE", "pu", "42"))
    assert.Contains (t, vars, DefaultWslVar ("JOHN", "pu", "42"))
}

func TestCombineCommandLastEmpty (t *testing.T) {
    resolution := "last"

    life :=  WithResolution (ReadCombineConfig (t, "command-answer-life.yml"), "self")
    john :=  WithResolution (DefaultConfig (), resolution)

            actual := Combine (life, john)
    vars := actual.Command.Environment.Var

    assert.Equal    (t, Resolution (resolution), actual.Command.Environment.Resolution)
    assert.Len      (t, vars, 2)
    assert.Contains (t, vars, DefaultWslVar ("ANSWER", "", "42"))
    assert.Contains (t, vars, DefaultWslVar ("LIFE", "pu", "42"))
}

func ReadCombineConfig (t *testing.T, filename string) (config *Config) {
    config, err := ObtainYaml ("./test/fixture/combine/" + filename)
    assert.NoError (t, err)

    return
}

func WithResolution (config *Config, resolution string) *Config {
           config.Command.Environment.Resolution = Resolution (resolution)
    return config
}
