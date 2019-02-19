# Codesk [![Backer](https://img.shields.io/badge/backer-codesk-orange.svg?style=flat)](https://www.patreon.com/encodeering) [![Build Status](https://travis-ci.org/encodeering/codesk.svg?branch=master)](https://travis-ci.org/encodeering/codesk)

A desk to communicate with your wsl linux environment written in Go.

## Setup

You can build the binaries yourself with `make`.
Please make sure that your build environment includes `golang` and `build-essential` package or appropriate alternatives.

A successful compilation will leave *.exe* files within the directories of the modules.

### Binary Configuration

``` yaml
command:
  environment:
    resolution: [first,parent,self,last]
    var: [
      "LIFE=u=42"
    ]
```

Resolution strategy:

| value  | description                                                      |
|--------|------------------------------------------------------------------|
| first  | never overwrite any existing value and add unknown values        |
| parent | take all values from the parent and ignore the rest              |
| self   | take all values from this configuration and ignore the rest      |
| last   | overwrite any existing value and add unknown values              |

The signature has to match *two* equal signs and encompass special [env flags](https://blogs.msdn.microsoft.com/commandline/2017/12/22/share-environment-vars-between-wsl-and-windows/) from wsl
