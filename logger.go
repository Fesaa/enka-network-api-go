package enkanetworkapigo

import (
	"github.com/rs/zerolog"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func logger() *zerolog.Logger {
	zerolog.DurationFieldUnit = time.Second
	zerolog.CallerMarshalFunc = callerMarshal
	l := zerolog.New(consoleWriter()).With().
		Caller().
		Timestamp().
		Logger().
		Level(zerolog.TraceLevel)
	return &l
}

func callerMarshal(pc uintptr, file string, line int) string {
	if strings.Contains(file, "go/pkg/mod") {
		file = filepath.Base(file)
	}
	return file + ":" + strconv.Itoa(line)
}

func consoleWriter() zerolog.ConsoleWriter {
	cw := zerolog.NewConsoleWriter()
	cw.TimeFormat = "2006-01-02 15:04:05"
	cw.PartsOrder = []string{
		zerolog.TimestampFieldName,
		zerolog.LevelFieldName,
		"handler",
		zerolog.CallerFieldName,
		zerolog.MessageFieldName,
	}
	cw.FieldsExclude = []string{
		"handler",
	}

	cw.NoColor = false
	cw.Out = os.Stdout
	return cw
}
