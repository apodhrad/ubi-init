package log

import (
	"fmt"
	"strings"
)

func Info(format string, a ...any) {
	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}
	fmt.Printf("[INFO] "+format, a...)
}

func Error(format string, a ...any) {
	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}
	fmt.Printf("[ERRO] "+format, a...)
}
