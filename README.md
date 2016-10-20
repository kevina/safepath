# Safe version of path/filepath for go

`filepath.Clean` in go mechanically removes ".." in paths.  This
behavior is incorrect on Unix when the path contains symbolic links.
However, this behavior is documented as such so the go developers are
unwilling to change `filepath.Clean`.  See
https://github.com/golang/go/issues/16255.

This package provides a version of path/filepath clean that is safe to
use on unix paths with symbolic links.

License: MIT
