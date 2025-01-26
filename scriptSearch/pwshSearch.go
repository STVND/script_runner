package scriptsearch

func NewPwshSearch() *scriptList {
	return NewScriptList([]string{"ps1"})
}

func (sl *scriptList) FindPwsh(dir string) {
	sl.FindScripts(dir)
}

func (sl *scriptList) GetPwsh() []script {
	return sl.GetScripts()
}
