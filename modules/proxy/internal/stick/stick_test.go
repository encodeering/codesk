package stick

import (
    "io/ioutil"
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestStickProcess (t * testing.T) {
    err := New (func (script []byte, fs []byte) error {
        gluesh, err := ioutil.ReadFile ("res/glue.sh")
        assert.NoError (t, err)
        assert.Equal (t, gluesh, script)

        gluetar, err := ioutil.ReadFile ("res/glue.tar.gz")
        assert.NoError (t, err)
        assert.Equal (t, gluetar, fs)

        return nil
    }).Process ()

    assert.NoError (t, err)
}
