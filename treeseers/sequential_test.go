// Test sequential version of LoadDirObject Method
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


func TestSeqLoadDirObject(t *testing.T) {
	tpdo := SeqLoadDirObject(neffs)
	tst1LoadDirObject(t,"SeqLoadDirObject",tpdo)
	_, thisfs, _, _ := runtime.Caller(0)
	thisfiledir := filepath.Dir(thisfs)
	tpdo = SeqLoadDirObject(thisfiledir)
	tst2LoadDirObject(t,"SeqLoadDirObject",tpdo)
}

// End of sequential_test.go
