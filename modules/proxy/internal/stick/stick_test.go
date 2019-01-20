package stick

import (
    "io/ioutil"
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestStickProcess (t * testing.T) {
    err := New (func (fs []byte) error {
        gluetar, err := ioutil.ReadFile ("res/glue.tar.gz")
        assert.NoError (t, err)
        assert.Equal (t, gluetar, fs)

        return nil
    }).Process ()

    assert.NoError (t, err)
}
