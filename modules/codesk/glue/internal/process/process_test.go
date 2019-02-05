package process

import (
    "io"
    "io/ioutil"
    "os"
    "strings"
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestExecOkay (t *testing.T) {
    capture (t, func (r io.Reader, w io.WriteCloser, _ io.Reader, _ io.WriteCloser) {
        assert.Equal   (t, 0, NewProxy ("echo").Exec ([]string{"-n", "hello"}))
        assert.NoError (t, w.Close ())

        content, err := ioutil.ReadAll (r)

        assert.NoError (t, err)
        assert.Equal   (t, "hello", string (content))
    })
}

func TestExecFail (t *testing.T) {
    capture (t, func (_ io.Reader, _ io.WriteCloser, r io.Reader, w io.WriteCloser) {
        assert.Equal   (t, 1, NewProxy ("echosounder").Exec ([]string{"-n", "hello"}))
        assert.NoError (t, w.Close ())

        content, err := ioutil.ReadAll (r)

        assert.NoError (t, err)
        assert.Equal   (t, "", string (content)) // stderr still empty ?
    })
}

func TestExecParentEnv (t *testing.T) {
    os.Setenv ("TESTVAR", "42")

    capture (t, func (r io.Reader, w io.WriteCloser, _ io.Reader, _ io.WriteCloser) {
        assert.Equal   (t, 0, NewProxy ("printenv").Exec ([]string{}))
        assert.NoError (t, w.Close ())

        content, err := ioutil.ReadAll (r)

        assert.NoError (t, err)
        assert.Equal   (t, strings.Join (os.Environ (), "\n") + "\n", string (content))
        assert.True    (t, strings.Contains (string (content), "TESTVAR=42"))
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
