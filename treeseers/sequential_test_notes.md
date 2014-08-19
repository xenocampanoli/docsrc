// Test function group designed to load using ReadDir/ReadFile methods from
// ioutil.
//
//  For testing:  this is not package TreeSeers.
package treeseers
//345678901234567890123456789012345678901234567890123456789012345678901234567890

import ( 
		"io/ioutil"
		"runtime"
		"testing"
	)

func TestGetPkgNameFromStr(t *testing.T) {
//RefNote:     Call to runtime.Caller(0) gets the name of this source file spec.
	_, thisfs, _, _ := runtime.Caller(0)
	scbytes, err := ioutil.ReadFile(thisfs)
	if err != nil {
		t.Errorf("TestGetPkgNameFromStr: Error reading file %s (%v)\n",thisfs,err)
		// But continue anyway
	}
	sourcecode := string(scbytes)
	pkgname := GetPkgNameFromStr(sourcecode)
	if pkgname != "treeseers" {
		t.Errorf("GetContentIfExists({this source file contents}) = \"%v\", want \"treeseers\"",pkgname)
	} 

	pkgname = GetPkgNameFromStr("No code here.")
	if pkgname != noPkgNameName {
		t.Errorf("GetContentIfExists(\"No code here.\") = \"%v\", want \"%v\"",pkgname,noPkgNameName)
	} 

	pkgname = GetPkgNameFromStr("")
	if pkgname != noPkgNameName {
		t.Errorf("GetContentIfExists(\"No code here.\") = \"%v\", want \"%v\"",pkgname,noPkgNameName)
	} 
}

func TestInitDirObject(t *testing.T) {
	tpdo := InitDirObject("/tmp/nothingeverthere")
	if len(tpdo.description) != 0 {
		t.Errorf("Fail description length for InitDirObject(\"Nonexistant file specified\") = \"%v\", want zero.",tpdo.description)
	}
	if len(tpdo.notes) != 0 {
		t.Errorf("Fail notes length for InitDirObject(\"Nonexistant file specified\") = \"%v\", want zero.",tpdo.notes)
	}
	if len(tpdo.readme) != 0 {
		t.Errorf("Fail readme length for InitDirObject(\"Nonexistant file specified\") = \"%v\", want zero.",tpdo.readme)
	}
	if len(tpdo.dtree) != 0 {
		t.Errorf("Fail dtree length for InitDirObject(\"Nonexistant file specified\") = \"%v\", want zero.",tpdo.dtree)
	}
	if len(tpdo.flist) != 0 {
		t.Errorf("Fail flist length for InitDirObject(\"Nonexistant file specified\") = \"%v\", want zero.",tpdo.flist)
	}
	_, thisfs, _, _ := runtime.Caller(0)
	thisfiledir := filepath.Dir(thisfs)
	tpdo := InitDirObject(thisfiledir)
	if len(tpdo.description) == 0 {
		t.Errorf("Fail description length for InitDirObject(\"Nonexistant file specified\") = \"%v\", want non-zero.",tpdo.description)
	}
	if len(tpdo.notes) == 0 {
		t.Errorf("Fail notes length for InitDirObject(\"Nonexistant file specified\") = \"%v\", want non-zero.",tpdo.notes)
	}
	if len(tpdo.readme) != 0 {
		t.Errorf("Fail readme length for InitDirObject(\"Nonexistant file specified\") = \"%v\", want non-zero.",tpdo.readme)
	}
	if len(tpdo.dtree) != 0 {
		t.Errorf("Fail dtree length for InitDirObject(\"Nonexistant file specified\") = \"%v\", want zero.",tpdo.dtree)
	}
	if len(tpdo.flist) != 0 {
		t.Errorf("Fail flist length for InitDirObject(\"Nonexistant file specified\") = \"%v\", want non-zero.",tpdo.flist)
	}
}

func TestLoadGoFileData(t *testing.T) {
}

func TestLoadDirectory(t *testing.T) {
}

// End of byreaddir_test.go
