package treeseers
//345678901234567890123456789012345678901234567890123456789012345678901234567890

import (
	"fmt"
	"io/ioutil"
)

type KVPair struct {
	dkey	string
	dobj	*DirObject
}

func ConLoadDirObject(dSpec string) (pdo *DirObject) {
	pdo = NewDirObject(dSpec)

	l, err :=	ioutil.ReadDir(dSpec)
	if err != nil {
		fmt.Printf("Error reading directory %s (%v)\n",dSpec,err)
		// But continue anyway
	} else {
		c := make(chan *KVPair, 5)
		tdc := 0
		for _, f := range l {
			fn := f.Name()
			fs := dSpec + "/" + fn
			if f.IsDir() {
				if pdo.Dtree == nil {
					pdo.Dtree = make(DirTree)
				}
				go func (kvpCh chan *KVPair, fN, fS string) {
						lkvp := KVPair{dkey: fN}
						lkvp.dobj = ConLoadDirObject(fS)
						kvpCh <- &lkvp
					}(c,fn,fs)
				tdc += 1
			} else {
				if pdo.Flist == nil {
					pdo.Flist = make(map[string]*FileObject)
				}
				if IsGoSrcFile(fn) {
					pdo.Flist[fn] = LoadFileObject(fs)
				}
			}
		}
		var kvp *KVPair
		for j := tdc;  j > 0;  j-- {
			kvp = <- c;
			pdo.Dtree[kvp.dkey] = kvp.dobj
		}
	}
	return
}
