package loglite

import (
	"fmt"
	"log"
	"os"
)

const (
	INFO  = "[INFO]"
	DEBUG = "[DEBUG]"
)

var debug bool = true
var ilogger *log.Logger
var dlogger *log.Logger

func init() {
	ilogger = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile|log.LUTC)
	dlogger = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile|log.LUTC)
}

func appendHead(v []interface{}, newObj interface{}) []interface{} {
	newv := make([]interface{}, len(v)+1)
	newv[0] = newObj
	copy(newv[1:], v)
	return newv
}

func SetDebug(show bool) {
	debug = show
}

func Info(v ...interface{}) {

	v = appendHead(v, INFO)
	result := fmt.Sprintln(v...)
	ilogger.Output(2, result)

}

func Infof(format string, v ...interface{}) {
	result := fmt.Sprintf(INFO+" "+format, v...)
	ilogger.Output(2, result)
}

func Debug(v ...interface{}) {
	if debug {
		v = appendHead(v, DEBUG)
		result := fmt.Sprintln(v...)
		dlogger.Output(2, result)
	}
}

func Debugf(format string, v ...interface{}) {
	if debug {
		result := fmt.Sprintf(DEBUG+" "+format, v...)
		dlogger.Output(2, result)
	}
}
