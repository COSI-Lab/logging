package logging

import (
	"fmt"
	"sync"
	"time"
)

type threadSafeLogger struct {
	sync.Mutex
}

var logger = threadSafeLogger{}

type messageType int

const (
	// InfoT is used for logging informational [INFO] messages
	InfoT messageType = iota
	// WarningT is used for logging warning [WARN] messages
	WarningT
	// ErrorT is used for logging error [ERROR] messages
	ErrorT
	// PanicT is used for logging panic [PANIC] messages
	PanicT
	// SuccessT is used for logging success [SUCCESS] messages
	SuccessT
)

const tm = "2006/01/02 15:04:05"

func (mt messageType) String() string {
	switch mt {
	case InfoT:
		return "\033[1m[INFO]    \033[0m| "
	case WarningT:
		return "\033[1m\033[33m[WARN]    \033[0m| "
	case ErrorT:
		return "\033[1m\033[31m[ERROR]   \033[0m| "
	case PanicT:
		return "\033[1m\033[34m[PANIC]   \033[0m| "
	case SuccessT:
		return "\033[1m\033[32m[SUCCESS] \033[0m| "
	default:
		return ""
	}
}

// LogEntryT enables programmatic creation of log entries
type LogEntryT struct {
	Type    messageType
	Message string
}

func (le LogEntryT) String() string {
	if le.Message[len(le.Message)-1] != '\n' {
		return fmt.Sprintf("%s %s %s\n", time.Now().Format(tm), le.Type.String(), le.Message)
	}

	return fmt.Sprintf("%s %s %s", time.Now().Format(tm), le.Type.String(), le.Message)
}

func logf(mt messageType, format string, v ...interface{}) {
	logger.Lock()
	if format[len(format)-1] != '\n' {
		fmt.Printf("%s %s %s\n", time.Now().Format(tm), mt.String(), fmt.Sprintf(format, v...))
	} else {
		fmt.Printf("%s %s %s", time.Now().Format(tm), mt.String(), fmt.Sprintf(format, v...))
	}
	logger.Unlock()
}

func logln(mt messageType, v ...interface{}) {
	logger.Lock()
	fmt.Printf("%s %s %s\n", time.Now().Format(tm), mt.String(), fmt.Sprint(v...))
	logger.Unlock()
}

// Infof formats a message and logs it with [INFO] tag, it adds a newline if the message didn't end with one
func Infof(format string, v ...interface{}) {
	logf(InfoT, format, v...)
}

// Info logs a message with [INFO] tag and a newline
func Info(v ...interface{}) {
	logln(InfoT, v...)
}

// Warningf formats a message and logs it with [WARN] tag, it adds a newline if the message didn't end with one
func Warningf(format string, v ...interface{}) {
	logf(WarningT, format, v...)
}

// Warning logs a message with [WARN] tag and a newline
func Warning(v ...interface{}) {
	logln(WarningT, v...)
}

// Errorf formats a message and logs it with [ERROR] tag, it adds a newline if the message didn't end with one
func Errorf(format string, v ...interface{}) {
	logf(ErrorT, format, v...)
}

// Error logs a message with [ERROR] tag and a newline
func Error(v ...interface{}) {
	logln(ErrorT, v...)
}

// Panicf formats a message and logs it with [PANIC] tag, it adds a newline if the message didn't end with one
// Note: this function does not call panic() or otherwise stops the program
func Panicf(format string, v ...interface{}) {
	logf(PanicT, format, v...)
}

// Panic logs a message with [PANIC] tag and a newline
// Note: this function does not call panic() or otherwise stops the program
func Panic(v ...interface{}) {
	logln(PanicT, v...)
}

// Successf formats a message and logs it with [SUCCESS] tag, it adds a newline if the message didn't end with one
func Successf(format string, v ...interface{}) {
	logf(SuccessT, format, v...)
}

// Success logs a message with [SUCCESS] tag and a newline
func Success(v ...interface{}) {
	logln(SuccessT, v...)
}

// Logf formats a message and logs it with provided tag, it adds a newline if the message didn't end with one
func Logf(mt messageType, format string, v ...interface{}) {
	logf(mt, format, v...)
}

// Log logs a message with provided tag and a newline
func Log(mt messageType, v ...interface{}) {
	logln(mt, v...)
}

// LogEntry logs a LogEntryT
func LogEntry(le LogEntryT) {
	logger.Lock()
	fmt.Print(le.String())
	logger.Unlock()
}
