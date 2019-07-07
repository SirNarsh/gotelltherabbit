package readconf

import (
	"fmt"
	"log"
	"os"
)

// CheckAllRequiredFiles Run at the beggining of the app to make sure all needed files are there
func CheckAllRequiredFiles() {
	names := []string{
		"general",
		"rabbit2http",
		"http2rabbit",
	}

	log.Println("Checking required files")
	for _, name := range names {
		checkRequiredFile(name)
	}
	log.Println("Config check completed")
}

// checkRequiredFile check config file exist, and exit otherwise
func checkRequiredFile(name string) {
	if _, err := os.Stat(fmt.Sprintf("./config/%s.json", name)); err != nil {
		if os.IsNotExist(err) {
			log.Fatalf(
				"Fatal Err: Required config file './config/%s.json' doesn't exist please copy the template from './config/%s.example.json'",
				name,
				name,
			)
		}
	}
	log.Printf("Required config file '/config/%s.json' found", name)
}
