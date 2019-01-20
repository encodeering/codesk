package process

import (
    "io/ioutil"
    "os"
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestExecOkay (t *testing.T) {
    r, w, err := os.Pipe ()
    assert.NoError (t, err)

    stdout := os.Stdout
    os.Stdout = w
    defer func () { os.Stdout = stdout } ()

    assert.Equal   (t, 0, Exec ("echo", []string{"-n", "hello"}))
    assert.NoError (t, w.Close ())

    content, err := ioutil.ReadAll (r)

    assert.NoError (t, err)
    assert.Equal   (t, "hello", string (content))
}

func TestExecFail (t *testing.T) {
    r, w, err := os.Pipe ()
    assert.NoError (t, err)

    stderr := os.Stderr
    os.Stderr = w
    defer func () { os.Stderr = stderr } ()

    assert.Equal   (t, 1, Exec ("echosounder", []string{"-n", "hello"}))
    assert.NoError (t, w.Close ())

    content, err := ioutil.ReadAll (r)

    assert.NoError (t, err)
    assert.Equal   (t, "", string (content)) // stderr still empty ?
}
