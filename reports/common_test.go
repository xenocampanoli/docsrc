package reports

import (
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"testing"
	)

const ( 
	neffs = "/tmp/nothingeverthere"
	)

var pwd string
var thisfs string
var thisfiledir string

// Subroutine Functions first in alphabetical order:

// Initialize runtime parameters for the tests.
func init() {
	pwd = os.Getenv("PWD")
//RefNote:     Call to runtime.Caller(0) gets the name of this source file spec.
	_, thisfs, _, _ = runtime.Caller(0)
	thisfiledir = filepath.Dir(thisfs)
}

func TestPrivate_genLineNoPage(t *testing.T) {
	content := genDirHdr(neffs)
	content = genLineNoPage(content)
	if len(content) == 0 {
		t.Errorf("genLineNoPage(genDirHdr(%s)) has no string content.",neffs)
	}
	if !regexp.MustCompile(`^\d+`).MatchString(content) {
		t.Errorf("genLineNoPage(genDirHdr(%s)) does not generate number prefixed lines.",neffs)
	}
}

func TestThisfile(t *testing.T) {
	if len(pwd) == 0 {
		t.Errorf("pwd has no string content.")
	}
	if len(thisfs) == 0 {
		t.Errorf("thisfs has no string content.")
	}
	if len(thisfiledir) == 0 {
		t.Errorf("thisfiledir has no string content.")
	}
}

