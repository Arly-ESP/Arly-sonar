package utilities

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/fatih/color" 
)

var logger *log.Logger

var (
	infoColor  = color.New(color.FgGreen).SprintFunc()
	errorColor = color.New(color.FgRed).SprintFunc()
	fatalColor = color.New(color.FgHiRed, color.Bold).SprintFunc()
	panicColor = color.New(color.FgMagenta, color.Bold).SprintFunc()
)

func currentEnv() string {
	env := os.Getenv("ENVIRONMENT")
	if env == "" {
		env = "development"
	}
	return env
}


func InitializeLogger() {
	prefix := "ArlyAPI: "
	env := currentEnv()

	if env != "production" {
		logger = log.New(os.Stdout, prefix, log.LstdFlags|log.Lshortfile)
	} else {
		file, err := os.OpenFile("error.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println("Failed to open error log file:", err)
			os.Exit(1)
		}
		logger = log.New(file, prefix, log.LstdFlags|log.Lshortfile)
	}
}

func LogInfo(message string) {
	if currentEnv() != "production" {
		logger.Println(infoColor(fmt.Sprintf("INFO: %s [%s]", message, currentTime())))
	}
}

func LogError(message string, err error) {
	logger.Println(errorColor(fmt.Sprintf("ERROR: %s: %v [%s]", message, err, currentTime())))
}

func LogFatal(message string, err error) {
	logger.Fatal(fatalColor(fmt.Sprintf("FATAL: %s: %v [%s]", message, err, currentTime())))
}

func LogPanic(message string, err error) {
	logger.Panic(panicColor(fmt.Sprintf("PANIC: %s: %v [%s]", message, err, currentTime())))
}

func currentTime() string {
	return time.Now().Format(time.RFC3339)
}



