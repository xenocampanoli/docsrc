package reports

import (
	"docsrc/treeseers"
	"fmt"
	"regexp"
	"testing"
	)

func TestPrivate_genDirHdr(t *testing.T) {
	content := genDirHdr(neffs)
	if len(content) == 0 {
		t.Errorf("genDirHdr(%s) has no string content.",neffs)
	}
}

func TestPrivate_genFileListings(t *testing.T) {
	tpdo := treeseers.SeqLoadDirObject(thisfiledir)
	content := genFileListings(tpdo.Flist)
	if (len(content) == 0) {
		t.Errorf("genFileListings(dO.Flist) generated no string content.")
	}
}

func TestPrivate_genPkgList(t *testing.T) {
	tpdo := treeseers.SeqLoadDirObject(thisfiledir)
	content := genPkgList(tpdo)
	if (len(content) == 0) {
		t.Errorf("genPkgList(dO) generated no string content.")
	}
}

func TestPrivate_genTxtHR(t *testing.T) {
	char := "."
	content := genTxtHR(char)
	rs := fmt.Sprintf("\\%s{%d}",char,rptwidth)
	if !regexp.MustCompile(rs).MatchString(content) {
		t.Errorf("genLineNoPage(\"%c\") does not generate %d column horizontal line of character %c.",char,rptwidth,char)
	}
}

func TestReport0(t *testing.T) {
	tpdo := treeseers.SeqLoadDirObject(thisfiledir)
	content := Report0(tpdo,0)
	if (len(content) == 0) {
		t.Errorf("ReportO(%s) generated no string content.",thisfiledir)
	}
}
