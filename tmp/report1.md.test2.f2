
# Source Directory (level 1):

`/home/xeno/shop/go/src/docsrc`

# Directory Description

# docsrc - a golang source tree document generator
## To help form documents directly from a Golang Source Tree.

People usually do not write well.  There are loads of reasons for
this, from social desires to hold resources to cognitive difficulties.
Not all tools that allow for more efficiency in writing a doc may
also provide encouragement to do so.  This tool is designed to help
write clear explanations around code without a lot of extra thought
of structure, thus leaving your personal neural networks more
available for the detail of what is actually helpful to explain.

When I first started as a programmer/consultant at the University
of Washington Health Sciences Building in 1985, early they let me
do one hour lectures teaching users how to use some of the
mainframe's facilities.  I had decided after attending the
sessions on these given by others that for most, a very simple
organization might be optimal for teaching and reinforcing the
concepts:

1.  Explain the following list of three additional items, so the
users don't get upset while going through the summary.  Explain as
quickly and clearly as possible that while listening to the summary,
they are to relax and just let it soak in, as during the detail part
all would be reinforced, and that last there would be time devoted
to their questions.  During this explanation, hand out any hardcopy
documents to be used with the session.  We had 50 minute sessions.
2.  Provide a 5 minute summary, concise and clear, but quick, with no
questions allowed, on the program, what it does, and how to use it.  
3.  Walk through the concepts to be taught during the session
exhaustively, and as clearly and concisely as possible, taking any
urgent questions, but at each point of such a question, reinforcing
that they would be allowed question time.  This section
typically should take 30 to 40 minutes of the session.  
4.  Answer questions as courteously, patiently, reponsively and
helpfully as possible, focusing on the need of the person asking,
and then on other group interests that may come to mind.  If the
session does not get taken up by questions, add explanations of
other functions to try to stimulate them, or review areas where
users had questions in the past to see if that brings up some.

I got this format from other kinds of things I'd learned in
school, from reinforcements I knew helped with me, and from other
teachers I had good conversations with about such matters.

This may not seem to fit perfectly with the present mechanism, but
it is intended primarly to accommodate this kind of pattern.  The
five minute introduction should end up in the Directory_description.md
files, followed by the README.md with the more full set of 
explanations for a new user.  Then, the source code itself is included
as a complement, and finally notes are provided at the end to
accommodate questions.  This application also will utilize sourcefile
specific documents:  1) sourcefile_description.md and sourcefile_notes.md,
however no sourcefile_README.md is provided for, as I presume that level
of detail should be documented in the source file itself, which by design
should be included in the reports in their entirety.

It may be that much solid but simple instruction can be weilded with such
a format in ways that reduce the cost of labor on both sides.  It is
certainly not rocket science, and there are likely other formats that can
be similarly helpful.  Anyway, that's my guess.  The report0
format does NOT accommodate this idea, as it is just for the source
code alone, but I hope all the others will.

Regarding usage then, to run this locally:

go run docsrc.go

from the top level directory where docsrc.go is natively found.

But of course, where it should become useful will be when it is
installed for use in arbitrary source trees for making application
documents.

# Go source file docsrc.go

Full level 1 filespec:  /home/xeno/shop/go/src/docsrc/docsrc.go)

```
1 // docsrc - Program to generate source tree documentation.
2 package main
3 
4 import (
5 	"docsrc/reports"
6 	"docsrc/treeseers"
7 	"bytes"
8 	"flag"
9 	"fmt"
10 	"os"
11 	"os/exec"
12 	"strings"
13 )
14 
15 // For now structures section (later reorganize)
16 
17 // CLConfiguration - Command Line Configurations for use by flag package.
18 // Values are parsed into this struct for passing to main for use in
19 // top level conditionals.
20 type CLConfiguration struct {
21 	report0    bool
22 	report1    bool
23 	report2    bool
24 	concurrent bool
25 	stdout     bool
26 }
27 
28 // getCLAC() - Get the command line switch values.
29 func getCLAC() *CLConfiguration {
30 	var clc CLConfiguration
31 	flag.BoolVar(&clc.report0, "report0", false, "Generate report0.txt, simple uncommented organized source listing.")
32 	flag.BoolVar(&clc.report1, "report1", false, "Generate report1.md, a relatively simple document but using all supporting files.")
33 	flag.BoolVar(&clc.report2, "report2", false, "Generate report2.html, report1 is generated, and report2 from report1.")
34 	flag.BoolVar(&clc.concurrent, "concurrent", false, "Use the concurrent version of the directory tree reader (for large trees).")
35 	flag.BoolVar(&clc.stdout, "stdout", false, "Write reports to stdout instead of report named files.")
36 	flag.Parse()
37 	// Apparently no err returned on zero value Flagset, which I think means this form.
38 	// I really did look around, and the examples are like this.
39 	return &clc
40 }
41 
42 // noReportSpecified() -  boolean function to indicate if a report is
43 // indicated on the command line so a default alternative may be chosen.
44 func (pCLC *CLConfiguration) noReportSpecified() bool {
45 	if pCLC.report0 {
46 		return false
47 	}
48 	if pCLC.report1 {
49 		return false
50 	}
51 	if pCLC.report2 {
52 		return false
53 	}
54 	return true
55 }
56 
57 // printUsage() - display USAGE statement.
58 func printUsage() {
59 	fmt.Printf("USAGE:  ./%s [argswtchs]", os.Args[0])
60 	flag.PrintDefaults()
61 }
62 
63 // Generates output to either named file, or stdout, according
64 // to the user's command line specification.
65 func outputReport(stdOut bool, cStr string, rName string) {
66 	if stdOut {
67 		fmt.Print(cStr)
68 	} else {
69 		fpo, err := os.Create(rName)
70 		if err != nil {
71 			panic(err)
72 		}
73 		// close fpo on exit and check for its returned error
74 		defer func() {
75 			if err := fpo.Close(); err != nil {
76 				panic(err)
77 			}
78 		}()
79 		_, err = fmt.Fprint(fpo, cStr)
80 		if err != nil {
81 			panic(err)
82 		}
83 	}
84 }
85 
86 // Generate report2.html from report1.md using shell markdown command.
87 func genHTMLfromMarkdown() (content string) {
88 	var out1, out2 bytes.Buffer
89 	cmd1 := exec.Command("which","markdown")
90 	cmd1.Stdout = &out1
91 	err := cmd1.Run()
92 	if err != nil {
93 		panic(err)
94 	}
95 	cmd2fs := strings.TrimSuffix(out1.String(),"\n")
96 	// The shortest path is /usr/bin/markdown, which is 17...
97 	if len(cmd2fs) > 16 {
98 		fmt.Println(cmd2fs)
99 		cmd2 := exec.Command(cmd2fs,"report1.md")
100 		cmd2.Stdout = &out2
101 		err = cmd2.Run()
102 		if err != nil {
103 			panic(err)
104 		}
105 		content = out2.String()
106 		return
107 	} else {
108 //345678901234567890123456789012345678901234567890123456789012345678901234567890
109 		nomdmsg := `
110 Cannot generate HTML output report2.html because markdown utility was not found
111 on your computer.  report1.md is at least available.  You may generate your HTML
112 file by hand after installing markdown by using the command:
113 
114 		markdown report1.md>out.html
115 
116 `
117 		fmt.Print(nomdmsg)
118 	}
119 	return
120 }
121 
122 // Generate the output report for the directory tree specified.
123 // Note that for now there is no way to specify a directory except
124 // to run the program in that directory.  This simplifies a lot of
125 // things, and reduces likelihood of messy mistakes.
126 func main() {
127 	pclc := getCLAC()
128 	pwd := os.Getenv("PWD")
129 	var pdo *treeseers.DirObject
130 	if pclc.concurrent {
131 		pdo = treeseers.ConLoadDirObject(pwd)
132 	} else {
133 		pdo = treeseers.SeqLoadDirObject(pwd)
134 	}
135 	if pclc.report0 || pclc.noReportSpecified() {
136 		content := reports.Report0(pdo, 1)
137 		outputReport(pclc.stdout, content, "report0.txt")
138 	}
139 	if pclc.report1 || pclc.report2 {
140 		content := reports.Report1(pdo, 1)
141 		outputReport(pclc.stdout, content, "report1.md")
142 		if pclc.report2 {
143 			content = genHTMLfromMarkdown()
144 			outputReport(pclc.stdout, content, "report2.html")
145 		}
146 	}
147 }
148 
149 //345678901234567890123456789012345678901234567890123456789012345678901234567890
150 // End of docsrc.go
151 
---snip---
```

End of level 1 source files in /home/xeno/shop/go/src/docsrc


## Source Directory (level 2):

`/home/xeno/shop/go/src/docsrc/reports`

## Directory Description


## Go source file common.go

Full level 2 filespec:  /home/xeno/shop/go/src/docsrc/reports/common.go)

```
1 //	report2.go
2 package reports
3 
4 import (
5 	"fmt"
6 	"strings"
7 )
8 
9 func genLineNoPage(cStr string) (oStr string) {
10 	la := strings.Split(cStr, "\n")
11 	l := 1
12 	oStr = ""
13 	for _, s := range la {
14 		oStr += fmt.Sprintf("%d %s\n", l, s)
15 		l += 1
16 	}
17 	return
18 }
19 
20 // End of common.go
21 
---snip---
```

## Go source file common_test.go

Full level 2 filespec:  /home/xeno/shop/go/src/docsrc/reports/common_test.go)

```
1 package reports
2 
3 import (
4 	"os"
5 	"path/filepath"
6 	"regexp"
7 	"runtime"
8 	"testing"
9 	)
10 
11 const ( 
12 	neffs = "/tmp/nothingeverthere"
13 	)
14 
15 var pwd string
16 var thisfs string
17 var thisfiledir string
18 
19 // Subroutine Functions first in alphabetical order:
20 
21 // Initialize runtime parameters for the tests.
22 func init() {
23 	pwd = os.Getenv("PWD")
24 //RefNote:     Call to runtime.Caller(0) gets the name of this source file spec.
25 	_, thisfs, _, _ = runtime.Caller(0)
26 	thisfiledir = filepath.Dir(thisfs)
27 }
28 
29 func TestPrivate_genLineNoPage(t *testing.T) {
30 	content := genDirHdr(neffs)
31 	content = genLineNoPage(content)
32 	if len(content) == 0 {
33 		t.Errorf("genLineNoPage(genDirHdr(%s)) has no string content.",neffs)
34 	}
35 	if !regexp.MustCompile(`^\d+`).MatchString(content) {
36 		t.Errorf("genLineNoPage(genDirHdr(%s)) does not generate number prefixed lines.",neffs)
37 	}
38 }
39 
40 func TestThisfile(t *testing.T) {
41 	if len(pwd) == 0 {
42 		t.Errorf("pwd has no string content.")
43 	}
44 	if len(thisfs) == 0 {
45 		t.Errorf("thisfs has no string content.")
46 	}
47 	if len(thisfiledir) == 0 {
48 		t.Errorf("thisfiledir has no string content.")
49 	}
50 }
51 
52 
---snip---
```

## Go source file report0.go

Full level 2 filespec:  /home/xeno/shop/go/src/docsrc/reports/report0.go)

```
1 //	report0.go
2 package reports
3 
4 import (
5 	"docsrc/treeseers"
6 	"fmt"
7 	"sort"
8 	"strings"
9 )
10 
11 const rptwidth = 132
12 
13 func genDirHdr(dSpec string) (content string) {
14 	content = genTxtHR("^")
15 	content += "\nSource Directory:  " + dSpec + "\n\n"
16 	return
17 }
18 
19 func genFileListings(fList treeseers.FileList) (content string) {
20 	fks := fList.NewKeySlice()
21 	sort.Strings(fks)
22 	content = "\n"
23 	for _, k := range fks {
24 		content += "Go source file:  " + k + ":\n\n"
25 		content += genLineNoPage(fList[k].Srccode)
26 		content += "---snip---\n\n"
27 	}
28 	return
29 }
30 
31 func genPkgList(dO *treeseers.DirObject) (content string) {
32 	filecount := dO.CountGoFiles()
33 	pkgcounts := make(treeseers.PkgCounts)
34 	dO.CountPkgs(pkgcounts)
35 	uniquepkgs := pkgcounts.NewKeySlice()
36 	sort.Strings(uniquepkgs)
37 	content += "Package counts:\n"
38 	for _, v := range uniquepkgs {
39 		content += fmt.Sprintf("%s: %d\n", v, pkgcounts[v])
40 	}
41 	content += fmt.Sprintf("Total Packages: %d\n", len(uniquepkgs))
42 	content += fmt.Sprintf("Total Go Source Files: %d\n", filecount)
43 	content += genTxtHR(".")
44 	return
45 }
46 
47 func genTxtHR(charStr string) string {
48 	content := strings.Repeat(charStr, rptwidth)
49 	content += "\n"
50 	return content
51 }
52 
53 func Report0(dO *treeseers.DirObject, dLevel int) (content string) {
54 //RefNote: Fspec is the directory spec for the Directory object:
55 	content = genDirHdr(dO.Fspec)
56 	content += genFileListings(dO.Flist)
57 	content += genTxtHR("~")
58 	dks := dO.Dtree.NewKeySlice()
59 	sort.Strings(dks)
60 	for _, k := range dks {
61 		content += Report0(dO.Dtree[k], dLevel+1)
62 	}
63 	if dLevel == 1 {
64 		content += genPkgList(dO)
65 	}
66 	return
67 }
68 
69 // End of report0.go
70 
---snip---
```

## Go source file report0_test.go

Full level 2 filespec:  /home/xeno/shop/go/src/docsrc/reports/report0_test.go)

```
1 package reports
2 
3 import (
4 	"docsrc/treeseers"
5 	"fmt"
6 	"regexp"
7 	"testing"
8 	)
9 
10 func TestPrivate_genDirHdr(t *testing.T) {
11 	content := genDirHdr(neffs)
12 	if len(content) == 0 {
13 		t.Errorf("genDirHdr(%s) has no string content.",neffs)
14 	}
15 }
16 
17 func TestPrivate_genFileListings(t *testing.T) {
18 	tpdo := treeseers.SeqLoadDirObject(thisfiledir)
19 	content := genFileListings(tpdo.Flist)
20 	if (len(content) == 0) {
21 		t.Errorf("genFileListings(dO.Flist) generated no string content.")
22 	}
23 }
24 
25 func TestPrivate_genPkgList(t *testing.T) {
26 	tpdo := treeseers.SeqLoadDirObject(thisfiledir)
27 	content := genPkgList(tpdo)
28 	if (len(content) == 0) {
29 		t.Errorf("genPkgList(dO) generated no string content.")
30 	}
31 }
32 
33 func TestPrivate_genTxtHR(t *testing.T) {
34 	char := "."
35 	content := genTxtHR(char)
36 	rs := fmt.Sprintf("\\%s{%d}",char,rptwidth)
37 	if !regexp.MustCompile(rs).MatchString(content) {
38 		t.Errorf("genLineNoPage(\"%c\") does not generate %d column horizontal line of character %c.",char,rptwidth,char)
39 	}
40 }
41 
42 func TestReport0(t *testing.T) {
43 	tpdo := treeseers.SeqLoadDirObject(thisfiledir)
44 	content := Report0(tpdo,0)
45 	if (len(content) == 0) {
46 		t.Errorf("ReportO(%s) generated no string content.",thisfiledir)
47 	}
48 }
49 
---snip---
```

## Go source file report1.go

Full level 2 filespec:  /home/xeno/shop/go/src/docsrc/reports/report1.go)

```
1 //	report1 - generate detailed markdown document for golang
2 //	source tree.
3 package reports
4 
5 import (
6 	"docsrc/treeseers"
7 	"fmt"
8 	"sort"
9 	"strings"
10 )
11 
12 func genMDLevelHdrPrefix(lNo int) string {
13 	mdhdrprefix := "######"
14 	if lNo < 6 {
15 		mdhdrprefix = strings.Repeat("#", lNo)
16 	}
17 	return mdhdrprefix + " "
18 }
19 
20 func genMDDirHdr(dLevel int, dSpec string) (content string) {
21 	hdrprefix := genMDLevelHdrPrefix(dLevel)
22 	content = fmt.Sprintf("\n%sSource Directory (level %d):\n",hdrprefix,dLevel)
23 	content += "\n`" + dSpec + "`\n"
24 	content += "\n" + hdrprefix + "Directory Description\n\n"
25 	return
26 }
27 
28 func genMDFileListings(dLevel int, fList treeseers.FileList,dSpec string) (content string) {
29 	hdrprefix := genMDLevelHdrPrefix(dLevel)
30 	fks := fList.NewKeySlice()
31 	sort.Strings(fks)
32 	content = "\n"
33 	for _, k := range fks {
34 		content += fmt.Sprintf("%sGo source file %s\n\n",hdrprefix,k)
35 		content += fmt.Sprintf("Full level %d filespec:  %s/%s)\n\n",dLevel,dSpec,k)
36 		content += "```\n"
37 		content += genLineNoPage(fList[k].Srccode)
38 		content += "---snip---\n"
39 		content += "```\n\n"
40 	}
41 	content += fmt.Sprintf("End of level %d source files in %s\n\n",dLevel,dSpec)
42 	return
43 }
44 
45 func genMDDirEndings(dLevel int, dO *treeseers.DirObject) (content string) {
46 	content = ""
47 	if dLevel == 1 {
48 		filecount := dO.CountGoFiles()
49 		pkgcounts := make(treeseers.PkgCounts)
50 		dO.CountPkgs(pkgcounts)
51 		uniquepkgs := pkgcounts.NewKeySlice()
52 		sort.Strings(uniquepkgs)
53 		content += "# Package Counts:\n"
54 		for _, v := range uniquepkgs {
55 			content += fmt.Sprintf("- %s: %d\n", v, pkgcounts[v])
56 		}
57 		content += fmt.Sprintf("- Total Packages: %d\n", len(uniquepkgs))
58 		content += fmt.Sprintf("- Total Go Source Files: %d\n", filecount)
59 	}
60 
61 	hdrprefix := genMDLevelHdrPrefix(dLevel)
62 	content += "\n" + hdrprefix + "Directory Notes for:\n"
63 	content += "`" + dO.Fspec + "`\n\n"
64 	content += dO.Notes
65 	content += fmt.Sprintf("\n%sEnding for %s, Level %d\n\n",hdrprefix,dO.Fspec,dLevel)
66 	return
67 }
68 
69 func Report1(dO *treeseers.DirObject, dLevel int) (content string) {
70 //RefNote: Fspec is the directory spec for the Directory object:
71 	content = genMDDirHdr(dLevel, dO.Fspec)
72 	content += dO.Description
73 	content += genMDFileListings(dLevel,dO.Flist,dO.Fspec)
74 	dks := dO.Dtree.NewKeySlice()
75 	sort.Strings(dks)
76 	for _, k := range dks {
77 		content += Report1(dO.Dtree[k], dLevel+1)
78 	}
79 	content += genMDDirEndings(dLevel, dO)
80 	return
81 }
82 
83 // End of report1.go
84 
---snip---
```

## Go source file report1_test.go

Full level 2 filespec:  /home/xeno/shop/go/src/docsrc/reports/report1_test.go)

```
1 package reports
2 
3 import (
4 	"docsrc/treeseers"
5 	"fmt"
6 	"regexp"
7 	"testing"
8 	)
9 
10 func TestPrivate_genMDLevelHdrPrefix(t *testing.T) {
11 	for l := 1; l <= 6; l++ {
12 		content := genMDLevelHdrPrefix(l)
13 		rs := fmt.Sprintf("^#{%d} $",l)
14 		if !regexp.MustCompile(rs).MatchString(content) {
15 			t.Errorf("genMDLevelHdrPrefix(%d) generates incorrect markdown header prefix '%s'.",l,content)
16 		}
17 	}
18 }
19 
20 func TestPrivate_genMDDirHdr(t *testing.T) {
21 	content := genMDDirHdr(2, pwd)
22 //var pkgNamePatt *regexp.Regexp = regexp.MustCompile(`(?m:^\s*package\s+(\w+)\s*$)`)
23 	x := `## Source Directory \(level 2\):`
24 	if !regexp.MustCompile(x).MatchString(content) {
25 		t.Errorf("genMDDirHdr(2,%s) generates incorrect primary markdown heading.\n",pwd)
26 		fmt.Print(content)
27 		fmt.Println(x)
28 	}
29 	if !regexp.MustCompile(pwd).MatchString(content) {
30 		t.Errorf("genMDDirHdr(2,%s) generates header missing directory spec:  '%s'.\n",pwd,pwd)
31 	}
32 	x = `## Directory Description`
33 	if !regexp.MustCompile(x).MatchString(content) {
34 		t.Errorf("genMDDirHdr(2,%s) generates incorrect description markdown heading.\n",pwd)
35 		fmt.Print(content)
36 		fmt.Println(x)
37 	}
38 }
39 
40 func TestPrivate_genMDDirEndings(t *testing.T) {
41 	pdo := treeseers.SeqLoadDirObject(pwd)
42 	content := genMDDirEndings(1, pdo)
43 	x := `# Package Counts:`
44 	if !regexp.MustCompile(x).MatchString(content) {
45 		t.Errorf("genMDDirHdr(2,%s) generates incorrect primary markdown heading.\n",pwd)
46 		fmt.Print(content)
47 		fmt.Println(x)
48 	}
49 	if !regexp.MustCompile(pwd).MatchString(content) {
50 		t.Errorf("genMDDirHdr(2,%s) generates header missing directory spec:  '%s'.\n",pwd,pwd)
51 	}
52 	x = `# Ending for \S+, Level 1`
53 	if !regexp.MustCompile(x).MatchString(content) {
54 		t.Errorf("genMDDirHdr(2,%s) generates incorrect primary markdown heading.\n",pwd)
55 		fmt.Print(content)
56 		fmt.Println(x)
57 	}
58 }
59 
60 func TestReport1(t *testing.T) {
61 	pdo := treeseers.SeqLoadDirObject(pwd)
62 	content := Report1(pdo, 1)
63 	if len(content) == 0 {
64 		t.Errorf("Report1(%s,pdo,1) yields no content.",pwd)
65 	}
66 }
67 
68 // report1_test.go
69 
---snip---
```

End of level 2 source files in /home/xeno/shop/go/src/docsrc/reports


## Directory Notes for:
`/home/xeno/shop/go/src/docsrc/reports`


## Ending for /home/xeno/shop/go/src/docsrc/reports, Level 2


## Source Directory (level 2):

`/home/xeno/shop/go/src/docsrc/tmp`

## Directory Description


End of level 2 source files in /home/xeno/shop/go/src/docsrc/tmp


## Directory Notes for:
`/home/xeno/shop/go/src/docsrc/tmp`


## Ending for /home/xeno/shop/go/src/docsrc/tmp, Level 2


## Source Directory (level 2):

`/home/xeno/shop/go/src/docsrc/treeseers`

## Directory Description

# treeseer package

## Function:

If it isn't clear, the primary calling program is docsrc.go, in the directory above
this one.

This package should simply read files of interest (*.go, *.md only for now) for use
in a quick example documenting scheme for a golang source tree.  I would like tests
to be reasonably supportive, but not thickly complete, but I am uneasy about the
test needs in this new language, so please I am open to suggestions, especially
from seasoned TDD/BDD afficionados in golang.

The data gets put into a tree of DirObject and FileObject structs, which are defined
in common.go.  This tree is then handed to an arbitrary reporter in the reports
package for use in generating output.

## Intents and Purposes:

I am calling this tree seer as a deprecative nom-de-guerre for software that is not
intended to be a whirlwind of lexing/parsing efficiency, but is only to get a simple
job done, to serve as a learning exercise for me, and to perhaps an example for
others interested in Go.

Tree Seer loads go source files and my own idea of documentation (which is optional) into a
struct tree for use by presentation reports.  My intention is to be able to helpfully
document an arbitrary go programming project tree with simple markdown, html or other
familiar format files for efficient understanding, AND instructional purposes.  Since this
is my first golang project I will humbly admit I am not sure this is ultimately helpful to
anyone but me as an exercise, but that it will merely do what I intended.  My primary
intentions with the activity were to:  1) learn go, 2) make something that others could
learn from that is clear, and well explained, and 3) possibly make something useful with an
unlikely bit of luck, but which is so simple that at least it could be expanded upon, if
realistic, to something more real anyway, if it were at least in some ballpark of long term
usefulness.

Originally I was to make a byreaddir.go and a byfpwalk.go, but filepath walk does a lexical
scan rather than a structural one of the directory tree, so it breaks the original idea I had
in mind for the structs/objects trees that I wanted to drive reports off of.  Given you end up
with an ordered list, (but out of order from the tree) it does make sense for some reporting purposes, but I want a structured
data representation, and then have the reports decide what to do with this, so I want only
availability created at load time, and not formatting or re-organization.  Given there is not
an issue with the amount of items likely to be seen for documenting, this just makes sense.

Tree seer therefore is divided between byreaddir.go, and common.go, the latter still which could
conceiveably be useful in the future if I add other load methods.  However, it is admittedly
counter to practice as typically golang philosophy is DRY and otherwise minimalist.  Hence,
should the possibility go away of future multiple loaders, or perhaps even not, the files and
tests should be merged into one.

## Go source file common.go

Full level 2 filespec:  /home/xeno/shop/go/src/docsrc/treeseers/common.go)

```
1 //  Common routines for use by members and clients of treeseers package.
2 
3 // Tree Seer is simply a mild mannered parsing package for use by docsrc for
4 // gathering custom source tree items of interest into a common reference, which
5 // starts with the DirObject defined below.  This object holds a tree subset of
6 // the source tree data for passing to the selected report packages.
7 //
8 //RefNote:  Note that these resources are automatically seen by all other go
9 //programming in this package directory.
10 package treeseers
11 //345678901234567890123456789012345678901234567890123456789012345678901234567890
12 
13 import (
14 	"fmt"
15 	"io/ioutil"
16 	"os"
17 	"regexp"
18 	"strings"
19 )
20 
21 const noPkgNameName ="<No Package Name Found In Source File>"
22 
23 // Common objects are things that will be used in reporting for both files and
24 // directories.
25 type CommonObject struct {
26 	Fspec		string
27 	Description	string
28 	Notes		string 
29 }
30 
31 // DirObject represents the directory tree from the starting point.
32 type DirObject struct {
33 	CommonObject
34 	README	string 
35 	Dtree	DirTree
36 	Flist	FileList
37 }
38 
39 // FileObject represents items of reporting interest for the specified source file.
40 type FileObject struct {
41 	CommonObject
42 	Pkgname	string
43 	Srccode	string
44 }
45 
46 type DirTree map[string]*DirObject
47 type FileList map[string]*FileObject
48 type PkgCounts map[string]int
49 
50 // Regular expression to match a go source file, which is to say something like:
51 // "common.go", "docsrc.go" or "common_test.go".
52 var goSrcFlPatt *regexp.Regexp = regexp.MustCompile(`^[A-Za-z0-9_]+\.go$`)
53 
54 // Regular expression presuming single line mode works well over multi-line
55 // string.
56 var pkgNamePatt *regexp.Regexp = regexp.MustCompile(`(?m:^\s*package\s+(\w+)\s*$)`)
57 
58 //345678901234567890123456789012345678901234567890123456789012345678901234567890
59 // Begin functions (I use 80 column terminals, so the rule is just for me: Xeno)
60 
61 // GetContentIfExists simply provides a string with file content if that file
62 // is found where specified, and otherwise returns a zero length string.  This
63 // allows variable reporting on source files and directories as the maintainer
64 // sees fit or otherwise finds convenient.  It also accommodates incremental
65 // building of an application and the writing of it's documentation.
66 func GetContentIfExists(fSpec string) string {
67 	bytes, err := ioutil.ReadFile(fSpec)
68 	// Note that in the case of errors in this file, returning
69 	// nil is not a signal of error, but just legitimate data
70 	// for passing to the reporter.  The whole idea here is to
71 	// not care if you just don't get the file.
72 	if os.IsNotExist(err) {
73 		return ""
74 	}
75 	if err != nil {
76 		fmt.Errorf("Unexpected Error reading %v, %q.(%v)",err,fSpec,os.IsNotExist(err))
77 		return ""
78 	}
79 	return string(bytes)
80 }
81 
82 // Count the total .go files referenced in the DirObject tree.
83 func (dO *DirObject) CountGoFiles() (count int) {
84 	count = len(dO.Flist)
85 	for _, ldo := range dO.Dtree {
86 		count += ldo.CountGoFiles()
87 	}
88 	return
89 }
90 
91 // Count the total .go files for each package referenced in the DirObject tree.
92 func (dO *DirObject) CountPkgs(pkgCounts PkgCounts) {
93 	for _, v := range dO.Flist {
94 		if _, ok := pkgCounts[v.Pkgname]; ok {
95 			pkgCounts[v.Pkgname] += 1
96 		} else {
97 			pkgCounts[v.Pkgname] = 1
98 		}
99 	}
100 	for _, ldo := range dO.Dtree {
101 		ldo.CountPkgs(pkgCounts)
102 	}
103 }
104 
105 // GetPkgNameFromStr takes a string from a presumed Golang source file and
106 // returns the package name listed first in the file.
107 func GetPkgNameFromStr(srcStr string) string {
108 	xa := pkgNamePatt.FindStringSubmatch(srcStr)
109 	if len(xa) > 1 {
110 		return xa[1]
111 	}
112 	return noPkgNameName
113 }
114 
115 // IsGoSrcFile is a boolean (bool) returning function for use in places where
116 // reporting must decide to use files iff they are such source files.
117 func IsGoSrcFile(fSpec string) (bool) {
118 	if len(fSpec) < 4 { // Because the smallest size is like "a.go".
119 		return false
120 	}
121 	if goSrcFlPatt.MatchString(fSpec) {
122 		return true
123 	}
124 	ra := strings.Split(fSpec,"/") // For the case of a filespec or URL.
125 	lra := len(ra)
126 	if lra > 0 {
127 		if goSrcFlPatt.MatchString(ra[lra-1]) {
128 			return true
129 		} 
130 	}
131 	return false
132 }
133 
134 // Load the FileObject data.
135 func LoadFileObject(fSpec string) (psfo *FileObject) {
136 	psfo = &FileObject{CommonObject: CommonObject{Fspec: fSpec}}
137 	scbytes, err := ioutil.ReadFile(fSpec)
138 	if err != nil {
139 		fmt.Printf("Error reading file %s %v(%v)\n",fSpec,err)
140 		// But continue anyway
141 	} else {
142 		psfo.Srccode = string(scbytes)
143 		psfo.Pkgname = GetPkgNameFromStr(psfo.Srccode)
144 	}
145 
146 	fsprefix := strings.TrimSuffix(fSpec,".go")
147 	psfo.Description = GetContentIfExists(fsprefix + "_description.md")
148 	psfo.Notes = GetContentIfExists(fsprefix + "_notes.md")
149 	return
150 }
151 
152 // Create a new DirObject struct.
153 func NewDirObject(dSpec string) (pdo *DirObject) {
154 	pdo = &DirObject{CommonObject: CommonObject{Fspec: dSpec}}
155 
156 	pdo.Description = GetContentIfExists(dSpec + "/Directory_description.md")
157 	pdo.Notes = GetContentIfExists(dSpec + "/Directory_notes.md")
158 	pdo.README = GetContentIfExists(dSpec + "/README.md")
159 	return
160 }
161 
162 // Generate a slice of keys for a Directory Tree map.
163 func (dTree *DirTree) NewKeySlice() (keyslice []string) {
164 	l := len(*dTree) 
165 	keyslice = make([]string,l,l)
166 	i := 0
167 	for k, _ := range *dTree {
168 		keyslice[i] = k
169 		i += 1
170 	}
171 	return
172 }
173 
174 // Generate a slice of keys for a File List map.
175 func (fList *FileList) NewKeySlice() (keyslice []string) {
176 	l := len(*fList) 
177 	keyslice = make([]string,l,l)
178 	i := 0
179 	for k, _ := range *fList {
180 		keyslice[i] = k
181 		i += 1
182 	}
183 	return
184 }
185 
186 // Generate a slice of keys for a Package Counts map.
187 func (pkgCounts PkgCounts) NewKeySlice() (keyslice []string) {
188 	l := len(pkgCounts) 
189 	keyslice = make([]string,l,l)
190 	i := 0
191 	for k, _ := range pkgCounts {
192 		keyslice[i] = k
193 		i += 1
194 	}
195 	return
196 }
197 
198 // End of common.go (because Xeno likes to put that here anyway)5678901234567890
199 
---snip---
```

## Go source file common_test.go

Full level 2 filespec:  /home/xeno/shop/go/src/docsrc/treeseers/common_test.go)

```
1 // Test functions for use across treeseer package, or by other tests.
2 package treeseers
3 //345678901234567890123456789012345678901234567890123456789012345678901234567890
4 
5 import ( 
6 		"path/filepath"
7 		"io/ioutil"
8 		"runtime"
9 		"testing"
10 	)
11 
12 const ( 
13 	neffs = "/tmp/nothingeverthere"
14 	)
15 
16 var thisfs string
17 
18 // Subroutine Functions first in alphabetical order:
19 
20 // Initialize runtime parameters for the tests.
21 func init() {
22 //RefNote:     Call to runtime.Caller(0) gets the name of this source file spec.
23 	_, thisfs, _, _ = runtime.Caller(0)
24 }
25 
26 func tpField(t *testing.T, mUT string, fN string, fL int, zV bool) {
27 	if zV {
28 		if fL != 0 {
29 			t.Errorf("Fail %s length for %s(%s) = \"%v\", want zero.",fN,mUT,fL)
30 		}
31 	} else {
32 		if fL == 0 {
33 			t.Errorf("Fail %s length for %s(%s) = zero, want non-zero.",fN,mUT)
34 		}
35 	}
36 }
37 
38 func tst1LoadDirObject(t *testing.T,tL string,tO *DirObject) {
39 	tpField(t,"LoadDirObject","Description",len(tO.Description),true)
40 	tpField(t,"LoadDirObject","Notes",len(tO.Notes),true)
41 	tpField(t,"LoadDirObject","README",len(tO.README),true)
42 	tpField(t,"LoadDirObject","Dtree",len(tO.Dtree),true)
43 	tpField(t,"LoadDirObject","Flist",len(tO.Flist),true)
44 }
45 
46 func tst2LoadDirObject(t *testing.T,tL string,tO *DirObject) {
47 	tpField(t,"LoadDirObject","Description",len(tO.Description),false)
48 	tpField(t,"LoadDirObject","Notes",len(tO.Notes),false)
49 	tpField(t,"LoadDirObject","README",len(tO.README),false)
50 	tpField(t,"LoadDirObject","Dtree",len(tO.Dtree),true)
51 	tpField(t,"LoadDirObject","Flist",len(tO.Flist),false)
52 }
53 
54 // Actual Tests Last, in alphabetical order:
55 
56 func TestCountGoFiles(t *testing.T) {
57 	tpdo := NewDirObject(neffs)
58 	count := tpdo.CountGoFiles()
59 	if count != 0 {
60 		t.Errorf("TestCountGoFiles: FAIL count test should be 0, was %d.\n",count)
61 	}
62 	tpdo.Flist = make(FileList)
63 	tpdo.Flist["fakefile.go"] = new(FileObject)
64 	tpdo.Flist["fakefile.go"].Pkgname = "fakepackage"
65 	tpdo.Flist["fakefile.go"].Srccode = "fakesource"
66 	count = tpdo.CountGoFiles()
67 	if count != 1 {
68 		t.Errorf("TestCountGoFiles: FAIL count test should be 1, was %d.\n",count)
69 	}
70 }
71 
72 func TestCountPkgs(t *testing.T) {
73 	tpdo := NewDirObject(neffs)
74 	pkgcounts := make(PkgCounts)
75 	tpdo.CountPkgs(pkgcounts)
76 	l := len(pkgcounts)
77 	if l != 0 {
78 		t.Errorf("TestCountPkgs: FAIL package count test should be 0, was %d.\n",l)
79 	}
80 	tpdo.Flist = make(FileList)
81 	tpdo.Flist["fakefile.go"] = new(FileObject)
82 	tpdo.Flist["fakefile.go"].Pkgname = "fakepackage"
83 	tpdo.Flist["fakefile.go"].Srccode = "fakesource"
84 	pkgcounts2 := make(PkgCounts)
85 	tpdo.CountPkgs(pkgcounts2)
86 	l = len(pkgcounts2)
87 	if l != 1 {
88 		t.Errorf("TestCountPkgs: FAIL package count test should be 1, was %d.\n",l)
89 	}
90 	for _, v := range pkgcounts {
91 		if v != 1 {
92 			t.Errorf("TestCountPkgs: FAIL package count test must be 1 or more, was %d.\n",v)
93 		}
94 	}
95 }
96 
97 func TestGetContentIfExists(t *testing.T) {
98 	result := GetContentIfExists(neffs)
99 	if result != "" {
100 		t.Errorf("TestGetContentIfExists: FAIL reading nonexistent file %s (%v)\n",neffs,result)
101 	}
102 	result = GetContentIfExists(thisfs)
103 	if len(result) <= 0 {
104 		t.Errorf("TestGetContentIfExists: FAIL reading file %s (%v)\n",thisfs,result)
105 	}
106 }
107 
108 func TestGetPkgNameFromStr(t *testing.T) {
109 	scbytes, err := ioutil.ReadFile(thisfs)
110 	if err != nil {
111 		t.Errorf("TestGetPkgNameFromStr: Error reading file %s (%v)\n",thisfs,err)
112 		// But continue anyway
113 	}
114 	sourcecode := string(scbytes)
115 	pkgname := GetPkgNameFromStr(sourcecode)
116 	if pkgname != "treeseers" {
117 		t.Errorf("GetPkgNameFromStr({this source file contents}) = \"%v\", want \"treeseers\"",pkgname)
118 	} 
119 
120 	pkgname = GetPkgNameFromStr("No code here.")
121 	if pkgname != noPkgNameName {
122 		t.Errorf("GetPkgNameFromStr(\"No code here.\") = \"%v\", want \"%v\"",pkgname,noPkgNameName)
123 	} 
124 
125 	pkgname = GetPkgNameFromStr("")
126 	if pkgname != noPkgNameName {
127 		t.Errorf("GetPkgNameFromStr(\"\") = \"%v\", want \"%v\"",pkgname,noPkgNameName)
128 	} 
129 }
130 
131 func TestIsGoSrcFile(t *testing.T) {
132 	for _, v := range []string{"myreallylongsourcefilenamethatissolongyoudontwannareadit.go","b.go"} {
133 		if !IsGoSrcFile(v) {
134 			t.Errorf("IsGoSrcFile(\"%s\") == false, want true.",v)
135 		}
136 	}
137 	for _, v := range []string{"myreallylongsourcefilenamethatissolongyoudontwannareadit.goesaway","b.g",".go","bgo"} {
138 		if IsGoSrcFile(v) {
139 			t.Errorf("IsGoSrcFile(\"%s\") == true, want false.",v)
140 		}
141 	}
142 }
143 
144 func TestLoadFileObject(t *testing.T) {
145 	tpfo := LoadFileObject(neffs)
146 	tpField(t,"LoadFileObject","Description",len(tpfo.Description),true)
147 	tpField(t,"LoadFileObject","Notes",len(tpfo.Notes),true)
148 	tpField(t,"LoadFileObject","Pkgname",len(tpfo.Pkgname),true)
149 	tpField(t,"LoadFileObject","Srccode",len(tpfo.Srccode),true)
150 	tpfo = LoadFileObject(thisfs)
151 	tpField(t,"LoadFileObject","Description",len(tpfo.Description),false)
152 	tpField(t,"LoadFileObject","Notes",len(tpfo.Notes),false)
153 	tpField(t,"LoadFileObject","Pkgname",len(tpfo.Pkgname),false)
154 	tpField(t,"LoadFileObject","Srccode",len(tpfo.Srccode),false)
155 }
156 
157 func TestNewDirObject(t *testing.T) {
158 	tpdo := NewDirObject(neffs)
159 	tpField(t,"NewDirObject","Description",len(tpdo.Description),true)
160 	tpField(t,"NewDirObject","Notes",len(tpdo.Notes),true)
161 	tpField(t,"NewDirObject","README",len(tpdo.README),true)
162 	tpField(t,"NewDirObject","Dtree",len(tpdo.Dtree),true)
163 	tpField(t,"NewDirObject","Flist",len(tpdo.Flist),true)
164 	thisfiledir := filepath.Dir(thisfs)
165 	tpdo = NewDirObject(thisfiledir)
166 	tpField(t,"NewDirObject","Description",len(tpdo.Description),false)
167 	tpField(t,"NewDirObject","Notes",len(tpdo.Notes),false)
168 	tpField(t,"NewDirObject","README",len(tpdo.README),false)
169 	tpField(t,"NewDirObject","Dtree",len(tpdo.Dtree),true)
170 	tpField(t,"NewDirObject","Flist",len(tpdo.Flist),true)
171 }
172 
173 func TestNewKeySlice(t *testing.T) {
174 
175 	// on DirTree map type.
176 
177 	dtree := make(DirTree)
178 	ks := dtree.NewKeySlice()
179 	l := len(ks)
180 	if l != 0 {
181 		t.Errorf("dtree.NewKeySlice() has slice length %d, want 0.",l)
182 	}
183 	dtree = make(DirTree)
184 	dtree["fakedir"] = new(DirObject)
185 	dtree["fakedir2"] = new(DirObject)
186 	ks = dtree.NewKeySlice()
187 	l = len(ks)
188 	if l != 2 {
189 		t.Errorf("dtree.NewKeySlice() has slice length %d, want 2.",l)
190 	}
191 
192 	// on FileList map type.
193 
194 	flist := make(FileList)
195 	ks = flist.NewKeySlice()
196 	l = len(ks)
197 	if l != 0 {
198 		t.Errorf("flist.NewKeySlice() has slice length %d, want 0.",l)
199 	}
200 	flist = make(FileList)
201 	flist["fakefile.go"] = new(FileObject)
202 	flist["fakefile.go"].Pkgname = "fakepackage"
203 	flist["fakefile.go"].Srccode = "fakesource"
204 	ks = flist.NewKeySlice()
205 	l = len(ks)
206 	if l != 1 {
207 		t.Errorf("flist.NewKeySlice() has slice length %d, want 1.",l)
208 	}
209 
210 	// on PkgCounts map type.
211 
212 	pkgcounts := make(PkgCounts)
213 	ks = pkgcounts.NewKeySlice()
214 	l = len(ks)
215 	if l != 0 {
216 		t.Errorf("pkgcounts.NewKeySlice() has slice length %d, want 0.",l)
217 	}
218 	pkgcounts = make(PkgCounts)
219 	pkgcounts["fakepkg1"] = 1
220 	pkgcounts["fakepkg2"] = 2
221 	pkgcounts["fakepkg3"] = 3
222 	ks = pkgcounts.NewKeySlice()
223 	l = len(ks)
224 	if l != 3 {
225 		t.Errorf("pkgcounts.NewKeySlice() has slice length %d, want 3.",l)
226 	}
227 
228 }
229 
230 // End of common_test.go
231 
---snip---
```

## Go source file concurrent.go

Full level 2 filespec:  /home/xeno/shop/go/src/docsrc/treeseers/concurrent.go)

```
1 package treeseers
2 //345678901234567890123456789012345678901234567890123456789012345678901234567890
3 
4 import (
5 	"fmt"
6 	"io/ioutil"
7 )
8 
9 type KVPair struct {
10 	dkey	string
11 	dobj	*DirObject
12 }
13 
14 func ConLoadDirObject(dSpec string) (pdo *DirObject) {
15 	pdo = NewDirObject(dSpec)
16 
17 	l, err :=	ioutil.ReadDir(dSpec)
18 	if err != nil {
19 		fmt.Printf("Error reading directory %s (%v)\n",dSpec,err)
20 		// But continue anyway
21 	} else {
22 		c := make(chan *KVPair, 5)
23 		tdc := 0
24 		for _, f := range l {
25 			fn := f.Name()
26 			fs := dSpec + "/" + fn
27 			if f.IsDir() {
28 				if pdo.Dtree == nil {
29 					pdo.Dtree = make(DirTree)
30 				}
31 				go func (kvpCh chan *KVPair, fN, fS string) {
32 						lkvp := KVPair{dkey: fN}
33 						lkvp.dobj = ConLoadDirObject(fS)
34 						kvpCh <- &lkvp
35 					}(c,fn,fs)
36 				tdc += 1
37 			} else {
38 				if pdo.Flist == nil {
39 					pdo.Flist = make(map[string]*FileObject)
40 				}
41 				if IsGoSrcFile(fn) {
42 					pdo.Flist[fn] = LoadFileObject(fs)
43 				}
44 			}
45 		}
46 		var kvp *KVPair
47 		for j := tdc;  j > 0;  j-- {
48 			kvp = <- c;
49 			pdo.Dtree[kvp.dkey] = kvp.dobj
50 		}
51 	}
52 	return
53 }
54 
---snip---
```

## Go source file concurrent_test.go

Full level 2 filespec:  /home/xeno/shop/go/src/docsrc/treeseers/concurrent_test.go)

```
1 // Test concurrent version of LoadDirObject Method
2 package treeseers
3 
4 import ( 
5 		"path/filepath"
6 		"runtime"
7 		"testing"
8 	)
9 
10 // Subroutine Functions first in alphabetical order:
11 
12 // Initialize runtime parameters for the tests.
13 func init() {
14 //RefNote:     Call to runtime.Caller(0) gets the name of this source file spec.
15 	_, thisfs, _, _ = runtime.Caller(0)
16 }
17 
18 func TestConLoadDirObject(t *testing.T) {
19 	tpdo := ConLoadDirObject(neffs)
20 	tst1LoadDirObject(t,"ConLoadDirObject",tpdo)
21 	_, thisfs, _, _ := runtime.Caller(0)
22 	thisfiledir := filepath.Dir(thisfs)
23 	tpdo = ConLoadDirObject(thisfiledir)
24 	tst2LoadDirObject(t,"ConLoadDirObject",tpdo)
25 }
26 
27 // End of concurrent_test.go
28 
---snip---
```

## Go source file sequential.go

Full level 2 filespec:  /home/xeno/shop/go/src/docsrc/treeseers/sequential.go)

```
1 package treeseers
2 //345678901234567890123456789012345678901234567890123456789012345678901234567890
3 
4 import (
5 	"fmt"
6 	"io/ioutil"
7 )
8 
9 func SeqLoadDirObject(dSpec string) (pdo *DirObject) {
10 	pdo = NewDirObject(dSpec)
11 
12 	l, err :=	ioutil.ReadDir(dSpec)
13 	if err != nil {
14 		fmt.Printf("Error reading directory %s (%v)\n",dSpec,err)
15 		// But continue anyway
16 	} else {
17 		for _, f := range l {
18 			fn := f.Name()
19 			fs := dSpec + "/" + fn
20 			if f.IsDir() {
21 				if pdo.Dtree == nil {
22 					pdo.Dtree = make(map[string]*DirObject)
23 				}
24 				pdo.Dtree[fn] = SeqLoadDirObject(fs)
25 			} else {
26 				if pdo.Flist == nil {
27 					pdo.Flist = make(map[string]*FileObject)
28 				}
29 				if IsGoSrcFile(fn) {
30 					pdo.Flist[fn] = LoadFileObject(fs)
31 				}
32 			}
33 		}
34 	}
35 	return
36 }
37 
---snip---
```

## Go source file sequential_test.go

Full level 2 filespec:  /home/xeno/shop/go/src/docsrc/treeseers/sequential_test.go)

```
1 // Test sequential version of LoadDirObject Method
2 package treeseers
3 
4 import ( 
5 		"path/filepath"
6 		"runtime"
7 		"testing"
8 	)
9 
10 // Subroutine Functions first in alphabetical order:
11 
12 // Initialize runtime parameters for the tests.
13 func init() {
14 //RefNote:     Call to runtime.Caller(0) gets the name of this source file spec.
15 	_, thisfs, _, _ = runtime.Caller(0)
16 }
17 
18 
19 func TestSeqLoadDirObject(t *testing.T) {
20 	tpdo := SeqLoadDirObject(neffs)
21 	tst1LoadDirObject(t,"SeqLoadDirObject",tpdo)
22 	_, thisfs, _, _ := runtime.Caller(0)
23 	thisfiledir := filepath.Dir(thisfs)
24 	tpdo = SeqLoadDirObject(thisfiledir)
25 	tst2LoadDirObject(t,"SeqLoadDirObject",tpdo)
26 }
27 
28 // End of sequential_test.go
29 
---snip---
```

End of level 2 source files in /home/xeno/shop/go/src/docsrc/treeseers


## Directory Notes for:
`/home/xeno/shop/go/src/docsrc/treeseers`

# TODO notes for treeseer

## Ending for /home/xeno/shop/go/src/docsrc/treeseers, Level 2

# Package Counts:
- main: 1
- reports: 6
- treeseers: 6
- Total Packages: 3
- Total Go Source Files: 13

# Directory Notes for:
`/home/xeno/shop/go/src/docsrc`

# Notes regarding the docsrc *app*:

## Maintenance Objectives

Fixes and Additions to docsrc should accommodate all existing tests,
possibly adding new if that is a responsible step, and complying with
go fmt, or equivalent.

Tests in packages use the native test system as per recommendations.
Tests I write on docsrc.go use a bash script.  My rationale with
tests is they should not only help prevent maintenance problems, but
can also serve well as a learning tool, so for something that
executes from the shell, I want to test from the shell.

## My Web sites are:

1.  [Eskimo.com](http://www.eskimo.com/~xeno)
2.  [github.com](https://github.com/xenocampanoli)

## Documents I found efficiently helpful:

1.  [Effective Go](http://golang.org/doc/effective_go.html)
2.  [Command line switch parsing](http://golang.org/pkg/flag/)
3.  [Documenting Go Code](http://blog.golang.org/godoc-documenting-go-code)
4.  [Generally go to golang.org/doc](http://golang.org/doc)
5.  [I also used group #go-nuts on IRC a lot](https://webchat.freenode.net)
6.  [I also used group #go-nuts on IRC a lot](https://webchat.freenode.net)

## Links about Markdown:

1.	[Reference Document](http://daringfireball.net/projects/markdown/)
2.	[Cheatsheet](https://github.com/adam-p/markdown-here/wiki/Markdown-Cheatsheet)
3.	[Github flavored markdown specified.](https://help.github.com/articles/github-flavored-markdown)
4.	[Wiki Page](http://en.wikipedia.org/wiki/Markdown)

## Acknowledgements:

Many people in IRC, mostly in the #go-nuts group, helped me learn Go.  I could not have gotten as far as I have in the time
available to me without this help.  As such, it is certainly my secular karma, as it were, to aid others correspondingly,
and I will try to do so, especially if concerned with the specifics of this application / exercise.

## [Terry Gilliam](http://en.wikipedia.org/wiki/Terry_Gilliam)

1.	[Time Bandits](http://en.wikipedia.org/wiki/Time_Bandits)
2.	[Brazil](http://en.wikipedia.org/wiki/Brazil_\(1985_film\))

And no, despite all this, I really don't know Python very well.  I do know Ruby, Bash, C, and now I know Golang, at least
a little.

# Ending for /home/xeno/shop/go/src/docsrc, Level 1

