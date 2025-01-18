package runners

import (
	"os/exec"
)

func PwshRunner(scriptPath string) {
	exec.Command("cls")
	PwshInst := exec.Command("pwsh", "-File")
	PwshInst.Args = append(PwshInst.Args, scriptPath)
	Runner(PwshInst)
}
