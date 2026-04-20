package cmd

import (
	"log"
	"os"

	"github.com/fatih/color"
)

var (
	loggingLevel int
	logInfo      *log.Logger
	logWarning   *log.Logger
	logError     *log.Logger
)

func initLoggers() {
	hiCyan := color.New(color.FgHiCyan).SprintFunc()
	hiYellow := color.New(color.FgHiYellow).SprintFunc()
	hiRed := color.New(color.FgHiRed).SprintFunc()

	logInfo = log.New(os.Stderr, hiCyan("╭info\n╰"), 0)
	logWarning = log.New(os.Stderr, hiYellow("╭warning\n╰"), 0)
	logError = log.New(os.Stderr, hiRed("╭error\n╰"), 0)
}
