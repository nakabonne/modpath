# modpath

[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/nakabonne/modpath)

`modpath` detects the module path from the `go.mod` file underneath the given directory or its root.  

If no directory is explicitly given, current directory's module path is emitted as `go list -m` does.
As it also provides an API to retrieve the module path, you can handle it from anywhere with `modpath`.

## Installation

```
go get github.com/nakabonne/modpath
```

## Usage

```console
$ cat /path/to/foo/bar/go.mod
module example.com/foo/bar

go 1.13

$ modpath /path/to/foo/bar/
example.com/foo/bar
```

Combined with `goimports`, and then you can instructs it to sort the import paths with the module path into another group after 3rd-party packages, without any settings.

```diff
$ modpath | xargs -I{} goimports -local={} bar.go
diff -u bar.go.orig bar.go
--- bar.go.orig	2019-12-30 21:40:04.000000000 +0900
+++ bar.go	2019-12-30 21:40:04.000000000 +0900
@@ -3,8 +3,9 @@
 import (
 	"fmt"

-	"github.com/foo/bar"
 	"github.com/foo/baz"
+
+	"github.com/foo/bar"
 )
```


### Using in your project

```go
package main

import "github.com/nakabonne/modpath"

func main() {
	modulePath, _ := modpath.Run("/path/to/foo")
}
```

