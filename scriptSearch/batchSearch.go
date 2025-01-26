package scriptsearch

func NewBatchSearch() *scriptList {
	return NewScriptList([]string{"bat", "cmd", "btm"})
}

func (sl *scriptList) FindBatch(dir string) {
	sl.FindScripts(dir)
}

func (sl *scriptList) GetBatch() []script {
	return sl.GetScripts()
}
