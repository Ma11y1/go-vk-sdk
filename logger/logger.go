package logger

import (
	"fmt"
	"io"
	"os"
)

const prefix = "GOVk-api-logger"

var out io.Writer = os.Stdout
var isLog bool = false

func Log(from, message string) {
	if isLog {
		_, _ = fmt.Fprintf(out, fmt.Sprintf("[%s] %s: %s\n", prefix, from, message))
	}
}

func LogMessage(message string) {
	if isLog {
		_, _ = fmt.Fprintf(out, message)
	}
}

func Enable() {
	isLog = true
}

func Disable() {
	isLog = false
}

func SetOutput(w io.Writer) {
	out = w
}
