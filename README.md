# modpath

[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/nakabonne/modpath)
[![codecov](https://codecov.io/gh/nakabonne/modpath/branch/master/graph/badge.svg)](https://codecov.io/gh/nakabonne/modpath)

`modpath` provides an API to detect the module path from the `go.mod` file underneath the given directory or its root.

## Installation

```
go get github.com/nakabonne/modpath
```

## Usage

```go
package main

import (
	"fmt"

	"github.com/nakabonne/modpath"
)

func main() {
	modulePath, _ := modpath.Run("/path/to/foo/bar")
	fmt.Println(modulePath) // -> "example.com/foo/bar"
}
```

### Using as a replacement for `go list -m`

Unlike `go list -m`, you can inspect any directories except for the current directory.

```console
$ cat /path/to/foo/bar/go.mod
module example.com/foo/bar

go 1.13

$ modpath /path/to/foo/bar/
example.com/foo/bar
```

If no directory is explicitly given, the module containing the current directory is emitted as `go list -m` does.
