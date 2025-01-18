package configure

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func Setup() {
	configDir, _ := os.UserConfigDir()
	configDir = strings.ReplaceAll(configDir, "\\", "/")
	configDir += "/script_runner"

	if _, err := os.Stat(configDir); err != nil {
		log.Print("Config not found - attempting to create config\n")
		if err := os.Mkdir(configDir, 0755); err != nil {
			log.Print("Unable to create config directory: ", err, "\n")
			log.Fatal("Intended path: ", configDir)
		}
		create(configDir + "/directories.txt")
		log.Print("New \"directories\" file created\n")
	} else if _, err := os.Stat(configDir + "/directories.txt"); err != nil {
		log.Print("Unable to locate \"directories\" file\n")
		log.Print("Attempting to create \"directories\" file\n")
		create(configDir + "/directories.txt")
	} else {
		log.Print("Successfully located config!\n")
	}
}

func create(configPath string) {
	if _, err := os.Stat(configPath); err != nil {
		_, err := os.Create(configPath)
		if err != nil {
			log.Fatal("Unable to create directories file - exting runner: ", err, "\n")
		}

		scriptDir, _ := os.UserHomeDir()
		scriptDir = strings.ReplaceAll(scriptDir, "\\", "/")
		scriptDir += "/script_runner/scripts"

		buf := []byte(scriptDir)

		err = os.WriteFile(configPath, buf, 0644)
		if err != nil {
			log.Print("Unable to write to directories file: ", err, "\n")
			log.Print("File location: ", configPath, "\n")
			log.Fatalf("Exiting runner")
		}
	}
}

func ReadDirs() []string {
	configPath, _ := os.UserConfigDir()
	configPath = strings.ReplaceAll(configPath, "\\", "/")
	configPath += "/script_runner/directories.txt"

	file, _ := os.Open(configPath)
	scanner := bufio.NewScanner(file)

	var contents []string

	for scanner.Scan() {
		contents = append(contents, scanner.Text())
	}

	return contents
}
