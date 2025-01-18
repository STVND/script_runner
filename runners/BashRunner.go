package runners

import (
	"os/exec"
)

func BashRunner(scriptPath string) {
	exec.Command("clear")
	BashInst := exec.Command("pwsh", "-File")
	BashInst.Args = append(BashInst.Args, scriptPath)
	Runner(BashInst)
}
