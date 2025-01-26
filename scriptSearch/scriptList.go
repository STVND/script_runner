package scriptsearch

import (
	"os"
	"strings"
)

type script struct {
	directory string
	name      string
}
type scriptList struct {
	extensions []string
	scripts    []script
}

func NewScriptList(ext []string) *scriptList {
	return &scriptList{extensions: ext}
}

func newScript(dir, nme string) script {
	return script{directory: dir, name: nme}

}

func (sl *scriptList) FindScripts(dir string) {
	items, _ := os.ReadDir(dir)

	for _, ext := range sl.extensions {
		for _, item := range items {
			if item.IsDir() {
				sl.FindScripts(dir + item.Name())
			} else if strings.HasSuffix(item.Name(), ext) {
				script := newScript(dir, item.Name())
				sl.scripts = append(sl.scripts, script)
			}
		}
	}
}

func (sl *scriptList) GetScripts() []script {
	return sl.scripts
}
