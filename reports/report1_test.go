package reports

import (
	"docsrc/treeseers"
	"fmt"
	"regexp"
	"testing"
	)

func TestPrivate_genMDLevelHdrPrefix(t *testing.T) {
	for l := 1; l <= 6; l++ {
		content := genMDLevelHdrPrefix(l)
		rs := fmt.Sprintf("^#{%d} $",l)
		if !regexp.MustCompile(rs).MatchString(content) {
			t.Errorf("genMDLevelHdrPrefix(%d) generates incorrect markdown header prefix '%s'.",l,content)
		}
	}
}

func TestPrivate_genMDDirHdr(t *testing.T) {
	content := genMDDirHdr(2, pwd)
//var pkgNamePatt *regexp.Regexp = regexp.MustCompile(`(?m:^\s*package\s+(\w+)\s*$)`)
	x := `## Source Directory \(level 2\):`
	if !regexp.MustCompile(x).MatchString(content) {
		t.Errorf("genMDDirHdr(2,%s) generates incorrect primary markdown heading.\n",pwd)
		fmt.Print(content)
		fmt.Println(x)
	}
	if !regexp.MustCompile(pwd).MatchString(content) {
		t.Errorf("genMDDirHdr(2,%s) generates header missing directory spec:  '%s'.\n",pwd,pwd)
	}
	x = `## Directory Description`
	if !regexp.MustCompile(x).MatchString(content) {
		t.Errorf("genMDDirHdr(2,%s) generates incorrect description markdown heading.\n",pwd)
		fmt.Print(content)
		fmt.Println(x)
	}
}

func TestPrivate_genMDDirEndings(t *testing.T) {
	pdo := treeseers.SeqLoadDirObject(pwd)
	content := genMDDirEndings(1, pdo)
	x := `# Package Counts:`
	if !regexp.MustCompile(x).MatchString(content) {
		t.Errorf("genMDDirHdr(2,%s) generates incorrect primary markdown heading.\n",pwd)
		fmt.Print(content)
		fmt.Println(x)
	}
	if !regexp.MustCompile(pwd).MatchString(content) {
		t.Errorf("genMDDirHdr(2,%s) generates header missing directory spec:  '%s'.\n",pwd,pwd)
	}
	x = `# Ending for \S+, Level 1`
	if !regexp.MustCompile(x).MatchString(content) {
		t.Errorf("genMDDirHdr(2,%s) generates incorrect primary markdown heading.\n",pwd)
		fmt.Print(content)
		fmt.Println(x)
	}
}

func TestReport1(t *testing.T) {
	pdo := treeseers.SeqLoadDirObject(pwd)
	content := Report1(pdo, 1)
	if len(content) == 0 {
		t.Errorf("Report1(%s,pdo,1) yields no content.",pwd)
	}
}

// report1_test.go
