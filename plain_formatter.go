package logrus

import (
	"bytes"
	"fmt"
	"runtime"
	"strings"
)

const (
	_nocolor = "\033[0m"
	_red     = "\033[91m"
	_green   = "\033[92m"
	_yellow  = "\033[93m"
	_magenta = "\033[95m"
	_cyan    = "\033[96m"
)

const (
	_time_format = "2006-01-02 15:04:05"
)

func init() {
	isTerminal = IsTerminal()
}

type PlainFormatter struct {
	ForceColors bool
}

func (f *PlainFormatter) Format(entry *Entry) ([]byte, error) {
	levelText := strings.ToUpper(entry.Level.String())[0:1]
	buf := bytes.NewBufferString("")
	if (f.ForceColors || isTerminal) && runtime.GOOS != "windows" {
		color := _nocolor
		switch entry.Level {
		case DebugLevel:
			color = _cyan
		case InfoLevel:
			color = _green
		case WarnLevel:
			color = _yellow
		case ErrorLevel:
			color = _magenta
		case PanicLevel, FatalLevel:
			color = _red
		}
		buf.WriteString(color)
	}
	buf.WriteString(fmt.Sprintf("[%s] ", entry.Time.Format(_time_format)))
	buf.WriteString(fmt.Sprintf("[%s] ", levelText))
	if entry.Logger.EnableCallFunc {
		buf.WriteString(fmt.Sprintf("[%s:%d] ", entry.File, entry.line))
	}
	for k, v := range entry.Data {
		buf.WriteString(fmt.Sprintf("[%s=%v] ", k, v))
	}
	buf.WriteString(entry.Message)
	if (f.ForceColors || isTerminal) && runtime.GOOS != "windows" {
		buf.WriteString(_nocolor)
	}
	buf.WriteString("\n")
	return buf.Bytes(), nil
}
