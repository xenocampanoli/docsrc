// Test functions for use across treeseer package, or by other tests.
package treeseers
//345678901234567890123456789012345678901234567890123456789012345678901234567890

import ( 
		"path/filepath"
		"io/ioutil"
		"runtime"
		"testing"
	)

const ( 
	neffs = "/tmp/nothingeverthere"
	)

var thisfs string

// Subroutine Functions first in alphabetical order:

// Initialize runtime parameters for the tests.
func init() {
//RefNote:     Call to runtime.Caller(0) gets the name of this source file spec.
	_, thisfs, _, _ = runtime.Caller(0)
}

func tpField(t *testing.T, mUT string, fN string, fL int, zV bool) {
	if zV {
		if fL != 0 {
			t.Errorf("Fail %s length for %s(%s) = \"%v\", want zero.",fN,mUT,fL)
		}
	} else {
		if fL == 0 {
			t.Errorf("Fail %s length for %s(%s) = zero, want non-zero.",fN,mUT)
		}
	}
}

func tst1LoadDirObject(t *testing.T,tL string,tO *DirObject) {
	tpField(t,"LoadDirObject","Description",len(tO.Description),true)
	tpField(t,"LoadDirObject","Notes",len(tO.Notes),true)
	tpField(t,"LoadDirObject","README",len(tO.README),true)
	tpField(t,"LoadDirObject","Dtree",len(tO.Dtree),true)
	tpField(t,"LoadDirObject","Flist",len(tO.Flist),true)
}

func tst2LoadDirObject(t *testing.T,tL string,tO *DirObject) {
	tpField(t,"LoadDirObject","Description",len(tO.Description),false)
	tpField(t,"LoadDirObject","Notes",len(tO.Notes),false)
	tpField(t,"LoadDirObject","README",len(tO.README),false)
	tpField(t,"LoadDirObject","Dtree",len(tO.Dtree),true)
	tpField(t,"LoadDirObject","Flist",len(tO.Flist),false)
}

// Actual Tests Last, in alphabetical order:

func TestCountGoFiles(t *testing.T) {
	tpdo := NewDirObject(neffs)
	count := tpdo.CountGoFiles()
	if count != 0 {
		t.Errorf("TestCountGoFiles: FAIL count test should be 0, was %d.\n",count)
	}
	tpdo.Flist = make(FileList)
	tpdo.Flist["fakefile.go"] = new(FileObject)
	tpdo.Flist["fakefile.go"].Pkgname = "fakepackage"
	tpdo.Flist["fakefile.go"].Srccode = "fakesource"
	count = tpdo.CountGoFiles()
	if count != 1 {
		t.Errorf("TestCountGoFiles: FAIL count test should be 1, was %d.\n",count)
	}
}

func TestCountPkgs(t *testing.T) {
	tpdo := NewDirObject(neffs)
	pkgcounts := make(PkgCounts)
	tpdo.CountPkgs(pkgcounts)
	l := len(pkgcounts)
	if l != 0 {
		t.Errorf("TestCountPkgs: FAIL package count test should be 0, was %d.\n",l)
	}
	tpdo.Flist = make(FileList)
	tpdo.Flist["fakefile.go"] = new(FileObject)
	tpdo.Flist["fakefile.go"].Pkgname = "fakepackage"
	tpdo.Flist["fakefile.go"].Srccode = "fakesource"
	pkgcounts2 := make(PkgCounts)
	tpdo.CountPkgs(pkgcounts2)
	l = len(pkgcounts2)
	if l != 1 {
		t.Errorf("TestCountPkgs: FAIL package count test should be 1, was %d.\n",l)
	}
	for _, v := range pkgcounts {
		if v != 1 {
			t.Errorf("TestCountPkgs: FAIL package count test must be 1 or more, was %d.\n",v)
		}
	}
}

func TestGetContentIfExists(t *testing.T) {
	result := GetContentIfExists(neffs)
	if result != "" {
		t.Errorf("TestGetContentIfExists: FAIL reading nonexistent file %s (%v)\n",neffs,result)
	}
	result = GetContentIfExists(thisfs)
	if len(result) <= 0 {
		t.Errorf("TestGetContentIfExists: FAIL reading file %s (%v)\n",thisfs,result)
	}
}

func TestGetPkgNameFromStr(t *testing.T) {
	scbytes, err := ioutil.ReadFile(thisfs)
	if err != nil {
		t.Errorf("TestGetPkgNameFromStr: Error reading file %s (%v)\n",thisfs,err)
		// But continue anyway
	}
	sourcecode := string(scbytes)
	pkgname := GetPkgNameFromStr(sourcecode)
	if pkgname != "treeseers" {
		t.Errorf("GetPkgNameFromStr({this source file contents}) = \"%v\", want \"treeseers\"",pkgname)
	} 

	pkgname = GetPkgNameFromStr("No code here.")
	if pkgname != noPkgNameName {
		t.Errorf("GetPkgNameFromStr(\"No code here.\") = \"%v\", want \"%v\"",pkgname,noPkgNameName)
	} 

	pkgname = GetPkgNameFromStr("")
	if pkgname != noPkgNameName {
		t.Errorf("GetPkgNameFromStr(\"\") = \"%v\", want \"%v\"",pkgname,noPkgNameName)
	} 
}

func TestIsGoSrcFile(t *testing.T) {
	for _, v := range []string{"myreallylongsourcefilenamethatissolongyoudontwannareadit.go","b.go"} {
		if !IsGoSrcFile(v) {
			t.Errorf("IsGoSrcFile(\"%s\") == false, want true.",v)
		}
	}
	for _, v := range []string{"myreallylongsourcefilenamethatissolongyoudontwannareadit.goesaway","b.g",".go","bgo"} {
		if IsGoSrcFile(v) {
			t.Errorf("IsGoSrcFile(\"%s\") == true, want false.",v)
		}
	}
}

func TestLoadFileObject(t *testing.T) {
	tpfo := LoadFileObject(neffs)
	tpField(t,"LoadFileObject","Description",len(tpfo.Description),true)
	tpField(t,"LoadFileObject","Notes",len(tpfo.Notes),true)
	tpField(t,"LoadFileObject","Pkgname",len(tpfo.Pkgname),true)
	tpField(t,"LoadFileObject","Srccode",len(tpfo.Srccode),true)
	tpfo = LoadFileObject(thisfs)
	tpField(t,"LoadFileObject","Description",len(tpfo.Description),false)
	tpField(t,"LoadFileObject","Notes",len(tpfo.Notes),false)
	tpField(t,"LoadFileObject","Pkgname",len(tpfo.Pkgname),false)
	tpField(t,"LoadFileObject","Srccode",len(tpfo.Srccode),false)
}

func TestNewDirObject(t *testing.T) {
	tpdo := NewDirObject(neffs)
	tpField(t,"NewDirObject","Description",len(tpdo.Description),true)
	tpField(t,"NewDirObject","Notes",len(tpdo.Notes),true)
	tpField(t,"NewDirObject","README",len(tpdo.README),true)
	tpField(t,"NewDirObject","Dtree",len(tpdo.Dtree),true)
	tpField(t,"NewDirObject","Flist",len(tpdo.Flist),true)
	thisfiledir := filepath.Dir(thisfs)
	tpdo = NewDirObject(thisfiledir)
	tpField(t,"NewDirObject","Description",len(tpdo.Description),false)
	tpField(t,"NewDirObject","Notes",len(tpdo.Notes),false)
	tpField(t,"NewDirObject","README",len(tpdo.README),false)
	tpField(t,"NewDirObject","Dtree",len(tpdo.Dtree),true)
	tpField(t,"NewDirObject","Flist",len(tpdo.Flist),true)
}

func TestNewKeySlice(t *testing.T) {

	// on DirTree map type.

	dtree := make(DirTree)
	ks := dtree.NewKeySlice()
	l := len(ks)
	if l != 0 {
		t.Errorf("dtree.NewKeySlice() has slice length %d, want 0.",l)
	}
	dtree = make(DirTree)
	dtree["fakedir"] = new(DirObject)
	dtree["fakedir2"] = new(DirObject)
	ks = dtree.NewKeySlice()
	l = len(ks)
	if l != 2 {
		t.Errorf("dtree.NewKeySlice() has slice length %d, want 2.",l)
	}

	// on FileList map type.

	flist := make(FileList)
	ks = flist.NewKeySlice()
	l = len(ks)
	if l != 0 {
		t.Errorf("flist.NewKeySlice() has slice length %d, want 0.",l)
	}
	flist = make(FileList)
	flist["fakefile.go"] = new(FileObject)
	flist["fakefile.go"].Pkgname = "fakepackage"
	flist["fakefile.go"].Srccode = "fakesource"
	ks = flist.NewKeySlice()
	l = len(ks)
	if l != 1 {
		t.Errorf("flist.NewKeySlice() has slice length %d, want 1.",l)
	}

	// on PkgCounts map type.

	pkgcounts := make(PkgCounts)
	ks = pkgcounts.NewKeySlice()
	l = len(ks)
	if l != 0 {
		t.Errorf("pkgcounts.NewKeySlice() has slice length %d, want 0.",l)
	}
	pkgcounts = make(PkgCounts)
	pkgcounts["fakepkg1"] = 1
	pkgcounts["fakepkg2"] = 2
	pkgcounts["fakepkg3"] = 3
	ks = pkgcounts.NewKeySlice()
	l = len(ks)
	if l != 3 {
		t.Errorf("pkgcounts.NewKeySlice() has slice length %d, want 3.",l)
	}

}

// End of common_test.go
