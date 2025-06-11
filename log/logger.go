package log

import (
	"fmt"
	"log"
	"os"
)

const errprefix = "ERROR:"
const infoprefix = "INFO:"
const fatalprefix = "FATAL:"

const app_prefix = "[BIKE_STORE] "

var logger *log.Logger = log.New(os.Stdout, app_prefix, log.LstdFlags)

func Error(format string, args ...any) {
	logger.Printf("%s %s\n", errprefix, fmt.Sprintf(format, args...))
}

func Info(format string, args ...any) {
	logger.Printf("%s %s\n", infoprefix, fmt.Sprintf(format, args...))
}

func Fatalf(format string, args ...any) {
	logger.Fatalf("%s %s\n", fatalprefix, fmt.Sprintf(format, args...))
}
