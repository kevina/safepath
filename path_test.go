package safepath

import (
	"testing"
)

func TestClean(t *testing.T) {
	test := func (orig string, expect string) {
		res := Clean(orig)
		if res != expect {
			t.Errorf("CleanPath failed on '%s'. Got '%s'. Expected '%s'.",
				orig, res, expect)
		}
	}
	
	test("/a/b/c/", "/a/b/c/")
	test("//a/b/c/", "/a/b/c/")
	test("///a/b/c/", "/a/b/c/")
	test("/a/b//c", "/a/b/c")
	test("/a/b/c//d", "/a/b/c/d")
	test("./a/b/c", "./a/b/c")
	test("/a/.b/.c", "/a/.b/.c")
	test("/a/b/.c", "/a/b/.c")
	test("/a/./b/c", "/a/b/c")
	test("/a/b/./c", "/a/b/c")
	test("/.a/b/c", "/.a/b/c")
	test("/a/////b", "/a/b")
	test("////a/b", "/a/b")
	test("foo/.///./././///bar", "foo/bar")
}
