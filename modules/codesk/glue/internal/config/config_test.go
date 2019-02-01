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
