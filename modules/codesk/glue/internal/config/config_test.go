package config

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestCheckResolution (t *testing.T) {
    assert.NoError (t, CheckResolution ("first"))
    assert.NoError (t, CheckResolution ("parent"))
    assert.NoError (t, CheckResolution ("self"))
    assert.NoError (t, CheckResolution ("last"))

    assert.EqualError (t, CheckResolution (""), "resolution value '' is unknown")
    assert.EqualError (t, CheckResolution ("barent"), "resolution value 'barent' is unknown")
}

func TestCheckEnvvar (t *testing.T) {
    assert.NoError (t, CheckEnvvar ("ANSWER==42"))
    assert.NoError (t, CheckEnvvar ("ANSWER=="))
    assert.NoError (t, CheckEnvvar ("LIFE=pu=42"))
    assert.NoError (t, CheckEnvvar ("LIFE=pu="))

    // currently not restricting the number of flags or duplication
    assert.NoError (t, CheckEnvvar ("LIFE=pp=42"))
    assert.NoError (t, CheckEnvvar ("LIFE=ppppp=42"))

    assert.EqualError (t, CheckEnvvar (""),          "envvar value '' doesn't match pattern NAME=[wulp]=VALUE")
    assert.EqualError (t, CheckEnvvar ("LIFE=x=42"), "envvar value 'LIFE=x=42' doesn't match pattern NAME=[wulp]=VALUE")

    assert.EqualError (t, CheckEnvvar ("ANSWER"),    "envvar value 'ANSWER' doesn't match pattern NAME=[wulp]=VALUE")
    assert.EqualError (t, CheckEnvvar ("ANSWER="),   "envvar value 'ANSWER=' doesn't match pattern NAME=[wulp]=VALUE")
    assert.EqualError (t, CheckEnvvar ("ANSWER=42"), "envvar value 'ANSWER=42' doesn't match pattern NAME=[wulp]=VALUE")
}

func TestCheckEnvvars (t *testing.T) {
    assert.NoError (t, CheckEnvvars ([]string{}))
    assert.NoError (t, CheckEnvvars ([]string{"ANSWER==42", "LIFE=pu=42"}))

    assert.EqualError (t, CheckEnvvars ([]string{"LIFE=x=42"}), "envvar value 'LIFE=x=42' doesn't match pattern NAME=[wulp]=VALUE")
}

func TestConvertEnvvar (t *testing.T) {
    wslvar, err := ConvertEnvvar ("ANSWER==42")
    assert.NoError (t, err)
    assert.Equal (t, DefaultWslVar ("ANSWER", "", "42"), wslvar)

    wslvar, err = ConvertEnvvar ("LIFE=x=42")
    assert.EqualError (t, err, "envvar value 'LIFE=x=42' doesn't match pattern NAME=[wulp]=VALUE")
    assert.Equal (t, DefaultWslVar ("", "", ""), wslvar)
}

func TestConvertEnvvars (t *testing.T) {
    wslvars, err := ConvertEnvvars ([]string{ "ANSWER==42", "LIFE=pu=42" })
    assert.NoError (t, err)
    assert.Equal (t, []WslVar{ DefaultWslVar ("ANSWER", "", "42"), DefaultWslVar ("LIFE", "pu", "42") }, wslvars)

    wslvars, err = ConvertEnvvars ([]string{ "LIFE=x=42" })
    assert.EqualError (t, err, "envvar value 'LIFE=x=42' doesn't match pattern NAME=[wulp]=VALUE")
    assert.Equal (t, []WslVar(nil), wslvars)
}
