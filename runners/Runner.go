package runners

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"sync"
)

func Runner(TermInst *exec.Cmd) {
	inPipe, err := TermInst.StdinPipe()
	if err != nil {
		log.Print("Unable to get stdin: ", err, "\n")
	}

	outPipe, err := TermInst.StdoutPipe()
	if err != nil {
		log.Print("Unable to get stdout: ", err, "\n")
	}

	// errPipe, err := TermInst.StderrPipe()
	// if err != nil {
	// 	log.Print("Unable to get stderr: ", err, "\n")
	// }

	err = TermInst.Start()
	if err != nil {
		log.Print("Unable to run script: ", err, "\n")
	}
	var wg sync.WaitGroup

	go func() {
		runOutput(outPipe)
	}()

	go func() {
		runInput(inPipe)
	}()

	wg.Wait()
	TermInst.Wait()

}

func runOutput(rd io.ReadCloser) {
	reader := bufio.NewReader(rd)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			rd.Close()
			break
		} else if err != nil {
			log.Print("Error reading output: ", err, "\n")
			break
		}

		if line == "" {
			fmt.Println()
		} else {
			log.Println(line)
		}
	}
}

func runInput(wr io.WriteCloser) {
	reader := bufio.NewReader(os.Stdin)

	for {
		buf, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Print("Unable to read from stdin: ", err, "\n")
		}
		wr.Write([]byte(buf))
	}

}
