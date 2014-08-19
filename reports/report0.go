//	report0.go
package reports

import (
	"docsrc/treeseers"
	"fmt"
	"sort"
	"strings"
)

const rptwidth = 132

func genDirHdr(dSpec string) (content string) {
	content = genTxtHR("^")
	content += "\nSource Directory:  " + dSpec + "\n\n"
	return
}

func genFileListings(fList treeseers.FileList) (content string) {
	fks := fList.NewKeySlice()
	sort.Strings(fks)
	content = "\n"
	for _, k := range fks {
		content += "Go source file:  " + k + ":\n\n"
		content += genLineNoPage(fList[k].Srccode)
		content += "---snip---\n\n"
	}
	return
}

func genPkgList(dO *treeseers.DirObject) (content string) {
	filecount := dO.CountGoFiles()
	pkgcounts := make(treeseers.PkgCounts)
	dO.CountPkgs(pkgcounts)
	uniquepkgs := pkgcounts.NewKeySlice()
	sort.Strings(uniquepkgs)
	content += "Package counts:\n"
	for _, v := range uniquepkgs {
		content += fmt.Sprintf("%s: %d\n", v, pkgcounts[v])
	}
	content += fmt.Sprintf("Total Packages: %d\n", len(uniquepkgs))
	content += fmt.Sprintf("Total Go Source Files: %d\n", filecount)
	content += genTxtHR(".")
	return
}

func genTxtHR(charStr string) string {
	content := strings.Repeat(charStr, rptwidth)
	content += "\n"
	return content
}

func Report0(dO *treeseers.DirObject, dLevel int) (content string) {
//RefNote: Fspec is the directory spec for the Directory object:
	content = genDirHdr(dO.Fspec)
	content += genFileListings(dO.Flist)
	content += genTxtHR("~")
	dks := dO.Dtree.NewKeySlice()
	sort.Strings(dks)
	for _, k := range dks {
		content += Report0(dO.Dtree[k], dLevel+1)
	}
	if dLevel == 1 {
		content += genPkgList(dO)
	}
	return
}

// End of report0.go
