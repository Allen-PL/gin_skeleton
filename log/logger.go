package log

import (
	"fmt"
	"github.com/rs/zerolog"
	"os"
	"strings"
)

const timeFormat = "2006-01-02 15:04:05"

var (
	log zerolog.Logger
)

func SetUp() {
	// CallerSkipFrameCount is the number of stack frames to skip to find the caller.
	zerolog.CallerSkipFrameCount = 3
	// ConsoleWriter可输出便于我们阅读的，带颜色的日志，
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: timeFormat}
	// 进一步对ConsoleWriter进行配置，定制输出级别、信息、字段名、字段值的格式
	output.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf(" | %s", i))
	}
	output.FormatFieldName = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf(" | %s", i))
	}
	output.FormatFieldValue = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf(" | %s", i))
	}
	output.FormatCaller = func(i interface{}) string {
		var c string
		if cc, ok := i.(string); ok {
			c = cc
		}
		if len(c) > 0 {
			cwd, err := os.Getwd()
			if err == nil {
				c = strings.TrimPrefix(c, cwd)
				c = strings.TrimPrefix(c, "/")
			}
		}
		return "| " + c
	}
	log = zerolog.New(output).With().Timestamp().Logger()
}

func Debug(msg string) {
	log.Info().Caller().Msg(msg)
}

//Info : Level 1
func Info(msg string) {
	log.Info().Caller().Msg(msg)
}

//Warn : Level 2
func Warn(msg string) {
	log.Warn().Caller().Msg(msg)
}

//Error : Level 3
func Error(msg string) {
	log.Error().Caller().Msg(msg)
}

//Fatal : Level 4
func Fatal(msg string) {
	log.Fatal().Caller().Msg(msg)
}
