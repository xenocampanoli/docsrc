//  Common routines for use by members and clients of treeseers package.

// Tree Seer is simply a mild mannered parsing package for use by docsrc for
// gathering custom source tree items of interest into a common reference, which
// starts with the DirObject defined below.  This object holds a tree subset of
// the source tree data for passing to the selected report packages.
//
//RefNote:  Note that these resources are automatically seen by all other go
//programming in this package directory.
package treeseers
//345678901234567890123456789012345678901234567890123456789012345678901234567890

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

const noPkgNameName ="<No Package Name Found In Source File>"

// Common objects are things that will be used in reporting for both files and
// directories.
type CommonObject struct {
	Fspec		string
	Description	string
	Notes		string 
}

// DirObject represents the directory tree from the starting point.
type DirObject struct {
	CommonObject
	README	string 
	Dtree	DirTree
	Flist	FileList
}

// FileObject represents items of reporting interest for the specified source file.
type FileObject struct {
	CommonObject
	Pkgname	string
	Srccode	string
}

type DirTree map[string]*DirObject
type FileList map[string]*FileObject
type PkgCounts map[string]int

// Regular expression to match a go source file, which is to say something like:
// "common.go", "docsrc.go" or "common_test.go".
var goSrcFlPatt *regexp.Regexp = regexp.MustCompile(`^[A-Za-z0-9_]+\.go$`)

// Regular expression presuming single line mode works well over multi-line
// string.
var pkgNamePatt *regexp.Regexp = regexp.MustCompile(`(?m:^\s*package\s+(\w+)\s*$)`)

//345678901234567890123456789012345678901234567890123456789012345678901234567890
// Begin functions (I use 80 column terminals, so the rule is just for me: Xeno)

// GetContentIfExists simply provides a string with file content if that file
// is found where specified, and otherwise returns a zero length string.  This
// allows variable reporting on source files and directories as the maintainer
// sees fit or otherwise finds convenient.  It also accommodates incremental
// building of an application and the writing of it's documentation.
func GetContentIfExists(fSpec string) string {
	bytes, err := ioutil.ReadFile(fSpec)
	// Note that in the case of errors in this file, returning
	// nil is not a signal of error, but just legitimate data
	// for passing to the reporter.  The whole idea here is to
	// not care if you just don't get the file.
	if os.IsNotExist(err) {
		return ""
	}
	if err != nil {
		fmt.Errorf("Unexpected Error reading %v, %q.(%v)",err,fSpec,os.IsNotExist(err))
		return ""
	}
	return string(bytes)
}

// Count the total .go files referenced in the DirObject tree.
func (dO *DirObject) CountGoFiles() (count int) {
	count = len(dO.Flist)
	for _, ldo := range dO.Dtree {
		count += ldo.CountGoFiles()
	}
	return
}

// Count the total .go files for each package referenced in the DirObject tree.
func (dO *DirObject) CountPkgs(pkgCounts PkgCounts) {
	for _, v := range dO.Flist {
		if _, ok := pkgCounts[v.Pkgname]; ok {
			pkgCounts[v.Pkgname] += 1
		} else {
			pkgCounts[v.Pkgname] = 1
		}
	}
	for _, ldo := range dO.Dtree {
		ldo.CountPkgs(pkgCounts)
	}
}

// GetPkgNameFromStr takes a string from a presumed Golang source file and
// returns the package name listed first in the file.
func GetPkgNameFromStr(srcStr string) string {
	xa := pkgNamePatt.FindStringSubmatch(srcStr)
	if len(xa) > 1 {
		return xa[1]
	}
	return noPkgNameName
}

// IsGoSrcFile is a boolean (bool) returning function for use in places where
// reporting must decide to use files iff they are such source files.
func IsGoSrcFile(fSpec string) (bool) {
	if len(fSpec) < 4 { // Because the smallest size is like "a.go".
		return false
	}
	if goSrcFlPatt.MatchString(fSpec) {
		return true
	}
	ra := strings.Split(fSpec,"/") // For the case of a filespec or URL.
	lra := len(ra)
	if lra > 0 {
		if goSrcFlPatt.MatchString(ra[lra-1]) {
			return true
		} 
	}
	return false
}

// Load the FileObject data.
func LoadFileObject(fSpec string) (psfo *FileObject) {
	psfo = &FileObject{CommonObject: CommonObject{Fspec: fSpec}}
	scbytes, err := ioutil.ReadFile(fSpec)
	if err != nil {
		fmt.Printf("Error reading file %s %v(%v)\n",fSpec,err)
		// But continue anyway
	} else {
		psfo.Srccode = string(scbytes)
		psfo.Pkgname = GetPkgNameFromStr(psfo.Srccode)
	}

	fsprefix := strings.TrimSuffix(fSpec,".go")
	psfo.Description = GetContentIfExists(fsprefix + "_description.md")
	psfo.Notes = GetContentIfExists(fsprefix + "_notes.md")
	return
}

// Create a new DirObject struct.
func NewDirObject(dSpec string) (pdo *DirObject) {
	pdo = &DirObject{CommonObject: CommonObject{Fspec: dSpec}}

	pdo.Description = GetContentIfExists(dSpec + "/Directory_description.md")
	pdo.Notes = GetContentIfExists(dSpec + "/Directory_notes.md")
	pdo.README = GetContentIfExists(dSpec + "/README.md")
	return
}

// Generate a slice of keys for a Directory Tree map.
func (dTree *DirTree) NewKeySlice() (keyslice []string) {
	l := len(*dTree) 
	keyslice = make([]string,l,l)
	i := 0
	for k, _ := range *dTree {
		keyslice[i] = k
		i += 1
	}
	return
}

// Generate a slice of keys for a File List map.
func (fList *FileList) NewKeySlice() (keyslice []string) {
	l := len(*fList) 
	keyslice = make([]string,l,l)
	i := 0
	for k, _ := range *fList {
		keyslice[i] = k
		i += 1
	}
	return
}

// Generate a slice of keys for a Package Counts map.
func (pkgCounts PkgCounts) NewKeySlice() (keyslice []string) {
	l := len(pkgCounts) 
	keyslice = make([]string,l,l)
	i := 0
	for k, _ := range pkgCounts {
		keyslice[i] = k
		i += 1
	}
	return
}

// End of common.go (because Xeno likes to put that here anyway)5678901234567890
