//	report1 - generate detailed markdown document for golang
//	source tree.
package reports

import (
	"docsrc/treeseers"
	"fmt"
	"sort"
	"strings"
)

func genMDLevelHdrPrefix(lNo int) string {
	mdhdrprefix := "######"
	if lNo < 6 {
		mdhdrprefix = strings.Repeat("#", lNo)
	}
	return mdhdrprefix + " "
}

func genMDDirHdr(dLevel int, dSpec string) (content string) {
	hdrprefix := genMDLevelHdrPrefix(dLevel)
	content = fmt.Sprintf("\n%sSource Directory (level %d):\n",hdrprefix,dLevel)
	content += "\n`" + dSpec + "`\n"
	content += "\n" + hdrprefix + "Directory Description\n\n"
	return
}

func genMDFileListings(dLevel int, fList treeseers.FileList,dSpec string) (content string) {
	hdrprefix := genMDLevelHdrPrefix(dLevel)
	fks := fList.NewKeySlice()
	sort.Strings(fks)
	content = "\n"
	for _, k := range fks {
		content += fmt.Sprintf("%sGo source file %s\n\n",hdrprefix,k)
		content += fmt.Sprintf("Full level %d filespec:  %s/%s)\n\n",dLevel,dSpec,k)
		content += "```\n"
		content += genLineNoPage(fList[k].Srccode)
		content += "---snip---\n"
		content += "```\n\n"
	}
	content += fmt.Sprintf("End of level %d source files in %s\n\n",dLevel,dSpec)
	return
}

func genMDDirEndings(dLevel int, dO *treeseers.DirObject) (content string) {
	content = ""
	if dLevel == 1 {
		filecount := dO.CountGoFiles()
		pkgcounts := make(treeseers.PkgCounts)
		dO.CountPkgs(pkgcounts)
		uniquepkgs := pkgcounts.NewKeySlice()
		sort.Strings(uniquepkgs)
		content += "# Package Counts:\n"
		for _, v := range uniquepkgs {
			content += fmt.Sprintf("- %s: %d\n", v, pkgcounts[v])
		}
		content += fmt.Sprintf("- Total Packages: %d\n", len(uniquepkgs))
		content += fmt.Sprintf("- Total Go Source Files: %d\n", filecount)
	}

	hdrprefix := genMDLevelHdrPrefix(dLevel)
	content += "\n" + hdrprefix + "Directory Notes for:\n"
	content += "`" + dO.Fspec + "`\n\n"
	content += dO.Notes
	content += fmt.Sprintf("\n%sEnding for %s, Level %d\n\n",hdrprefix,dO.Fspec,dLevel)
	return
}

func Report1(dO *treeseers.DirObject, dLevel int) (content string) {
//RefNote: Fspec is the directory spec for the Directory object:
	content = genMDDirHdr(dLevel, dO.Fspec)
	content += dO.Description
	content += genMDFileListings(dLevel,dO.Flist,dO.Fspec)
	dks := dO.Dtree.NewKeySlice()
	sort.Strings(dks)
	for _, k := range dks {
		content += Report1(dO.Dtree[k], dLevel+1)
	}
	content += genMDDirEndings(dLevel, dO)
	return
}

// End of report1.go
