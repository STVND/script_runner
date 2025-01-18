package runners

import (
	"os/exec"
)

func BatchRunner(scriptPath string) {
	exec.Command("cls")
	BatchInst := exec.Command("pwsh", "-File")
	BatchInst.Args = append(BatchInst.Args, scriptPath)
	Runner(BatchInst)
}
