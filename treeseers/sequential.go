package treeseers
//345678901234567890123456789012345678901234567890123456789012345678901234567890

import (
	"fmt"
	"io/ioutil"
)

func SeqLoadDirObject(dSpec string) (pdo *DirObject) {
	pdo = NewDirObject(dSpec)

	l, err :=	ioutil.ReadDir(dSpec)
	if err != nil {
		fmt.Printf("Error reading directory %s (%v)\n",dSpec,err)
		// But continue anyway
	} else {
		for _, f := range l {
			fn := f.Name()
			fs := dSpec + "/" + fn
			if f.IsDir() {
				if pdo.Dtree == nil {
					pdo.Dtree = make(map[string]*DirObject)
				}
				pdo.Dtree[fn] = SeqLoadDirObject(fs)
			} else {
				if pdo.Flist == nil {
					pdo.Flist = make(map[string]*FileObject)
				}
				if IsGoSrcFile(fn) {
					pdo.Flist[fn] = LoadFileObject(fs)
				}
			}
		}
	}
	return
}
