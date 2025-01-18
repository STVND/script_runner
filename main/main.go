package main

import (
	"bufio"
	"log"
	configure "main/main.go/configuration"
	"os"
	"strconv"
)

func main() {
	configure.Setup()

	scanner := bufio.NewScanner(os.Stdin)
	log.Print("What would you like to do?\n\n")
	log.Print("[1] - List scripts\t[2] - Edit script directories\t\t[0] - Exit\n")

	for scanner.Scan() {
		var userInt int
		var err error
		line := scanner.Text()
		println(line)

		if line == "" {
			continue
		} else {
			userInt, err = strconv.Atoi(line)
			if err != nil {
				log.Print("Unable to parse input: ", err, "\n")
				continue
			}
		}

		switch userInt {
		case 0:
			log.Fatal("Closing Runner")
		case 1:
			listOption()
		case 2:
			configDirectories()
		default:
			log.Print("Did not understand input - please try again")
		}
	}

}

func listOption() {
	log.Print("list scripts called\n")
	os.Executable()
}

func configDirectories() {
	log.Print("directorie editor called")
}
