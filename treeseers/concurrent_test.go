// Test concurrent version of LoadDirObject Method
package treeseers

import ( 
		"path/filepath"
		"runtime"
		"testing"
	)

// Subroutine Functions first in alphabetical order:

// Initialize runtime parameters for the tests.
func init() {
//RefNote:     Call to runtime.Caller(0) gets the name of this source file spec.
	_, thisfs, _, _ = runtime.Caller(0)
}

func TestConLoadDirObject(t *testing.T) {
	tpdo := ConLoadDirObject(neffs)
	tst1LoadDirObject(t,"ConLoadDirObject",tpdo)
	_, thisfs, _, _ := runtime.Caller(0)
	thisfiledir := filepath.Dir(thisfs)
	tpdo = ConLoadDirObject(thisfiledir)
	tst2LoadDirObject(t,"ConLoadDirObject",tpdo)
}

// End of concurrent_test.go
