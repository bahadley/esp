package log

import (
	"fmt"
	"log"
	"os"
)

const (
	appName = "esp"

	// Logging message prefixes
	InfoPrefix  = "INFO"
	WarnPrefix  = "WARN"
	ErrPrefix   = "ERROR"
	DebugPrefix = "DEBUG"
)

func init() {
	// Change the device for logging to stdout
	log.SetOutput(os.Stdout)
}

func Logoutput(level string, msg string) {
	log.Printf("%s: [%s] %s", appName, level, msg)
}

func Logfatalerror(err error) {
	log.Fatal(fmt.Sprintf("%s: [%s] Exiting ... (%v)",
		appName, ErrPrefix, err))
}
