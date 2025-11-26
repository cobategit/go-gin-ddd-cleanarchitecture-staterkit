package logger

import (
	"log"
	"os"
)

var (
	infoLogger  = log.New(os.Stdout, "[Info] ", log.LstdFlags|log.Lshortfile)
	errorLogger = log.New(os.Stderr, "[Error] ", log.LstdFlags|log.Lshortfile)
)

func InitLogger(msg string, typ string, args ...any) {
	if typ == "info" {
		infoLogger.Printf(msg, args...)
	}
	if typ == "error" {
		errorLogger.Printf(msg, args...)
	}
}
