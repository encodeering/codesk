package stick

import (
    "github.com/gobuffalo/packr/v2"
)

type Handle func ([]byte, []byte) error

type Stick interface {
    Process () error
}

type stick struct {
    box *packr.Box
    handle Handle
}

func New (handle Handle) Stick {
    return &stick{
        box: packr.New ("stick-res", "res"),
        handle: handle,
    }
}

func (s *stick) Process () (err error) {
    var script []byte
    var fs []byte

    if script, err = s.box.Find ("glue.sh"); err != nil {
        return
    }
    if fs, err = s.box.Find ("glue.tar.gz"); err != nil {
        return
    }

    return s.handle (script, fs)
}
