package scriptsearch

func NewBashSearch() *scriptList {
	return NewScriptList([]string{"sh"})
}

func (sl *scriptList) FindBash(dir string) {
	sl.FindScripts(dir)
}

func (sl *scriptList) GetBash() []script {
	return sl.GetScripts()
}
