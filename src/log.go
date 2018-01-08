package main

import (
	"fmt"
	"os"
	"path"
	"time"
)

// Log - Write Message into Console and/or Logfile
func Log(message string) {
	Console(message)
	writeLogfile(message)
}

// Console - Write Message into the Program Console
func Console(message string) {
	fmt.Println(getLogTimestamp(), message)
}

// Error - Handle a error object and log it correctly
func Error(err error) {
	Log("Fatal Error: " + err.Error())
}

func writeLogfile(message string) {
	if ConfigGetValue("log", "enable") == "true" {
		logpath := path.Join(ConfigGetValue("log", "logdir") + ConfigGetValue("log", "logfile"))
		if _, err := os.Stat(ConfigGetValue("log", "logdir")); os.IsNotExist(err) {
			os.MkdirAll(ConfigGetValue("log", "logdir"), os.ModePerm)
		}
		f, err := os.OpenFile(logpath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			Error(err)
		}
		defer f.Close()
		_, err = f.WriteString(getLogTimestamp() + " " + message)
		if err != nil {
			Error(err)
		}
	}
}

func getLogTimestamp() string {
	return time.Now().Format(time.Stamp)
}
