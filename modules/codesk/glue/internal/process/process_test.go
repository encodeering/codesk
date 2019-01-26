package process

import (
    "io"
    "io/ioutil"
    "os"
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestExecOkay (t *testing.T) {
    capture (t, func (r io.Reader, w io.WriteCloser, _ io.Reader, _ io.WriteCloser) {
        assert.Equal   (t, 0, Exec ("echo", []string{"-n", "hello"}))
        assert.NoError (t, w.Close ())

        content, err := ioutil.ReadAll (r)

        assert.NoError (t, err)
        assert.Equal   (t, "hello", string (content))
    })
}

func TestExecFail (t *testing.T) {
    capture (t, func (_ io.Reader, _ io.WriteCloser, r io.Reader, w io.WriteCloser) {
        assert.Equal   (t, 1, Exec ("echosounder", []string{"-n", "hello"}))
        assert.NoError (t, w.Close ())

        content, err := ioutil.ReadAll (r)

        assert.NoError (t, err)
        assert.Equal   (t, "", string (content)) // stderr still empty ?
    })
}

func capture (t* testing.T, f func (outR io.Reader, outW io.WriteCloser, errR io.Reader, errW io.WriteCloser)) {
    outR, outW, err := os.Pipe ()
    assert.NoError (t, err)
    defer outW.Close ()

    errR, errW, err := os.Pipe ()
    assert.NoError (t, err)
    defer errW.Close ()

    stdout := os.Stdout
    os.Stdout = outW
    defer func () { os.Stdout = stdout } ()

    stderr := os.Stderr
    os.Stderr = errW
    defer func () { os.Stderr = stderr } ()

    f (outR, outW, errR, errW)
}
