package share

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/GoldenDeals/StatusCoin/internal/share/shutdown"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type customLogger struct {
	defaultField string
	formatter    logrus.Formatter
}

func (l customLogger) Format(entry *logrus.Entry) ([]byte, error) {
	entry.Data["src"] = l.defaultField
	return l.formatter.Format(entry)
}

func NewLogger(src string, shut *shutdown.Shutdown) *logrus.Logger {
	log := logrus.New()
	log.SetLevel(logrus.Level(viper.GetUint32("log.level")))

	if viper.GetBool("log.json") {
		log.SetFormatter(customLogger{
			defaultField: src,
			formatter:    &logrus.JSONFormatter{PrettyPrint: true},
		})
	} else {
		log.SetFormatter(customLogger{
			defaultField: src,
			formatter:    &logrus.TextFormatter{},
		})
	}

	writers := make([]io.Writer, 0)
	for _, out := range viper.GetStringSlice("log.out") {
		switch out {
		case "stdout":
			writers = append(writers, os.Stdout)
		case "stderr":
			writers = append(writers, os.Stderr)
		default:
			writers = append(writers, openPathWithShutdown(out, shut))
		}
	}

	log.SetOutput(io.MultiWriter(writers...))
	return log
}

func openPathWithShutdown(out string, shut *shutdown.Shutdown) io.Writer {
	out = StringVarReplace(out)
	err := os.MkdirAll(filepath.Dir(out), 0755)
	if err != nil {
		panic(fmt.Errorf("unable to create all directories. error: %w", err))
	}

	file, err := os.Create(out)
	if err != nil {
		panic(fmt.Errorf("unable to open. error: %w", err))
	}

	if shut != nil {
		shut.Push(fmt.Sprintf("file: %s", file.Name()), func(_ string) error {
			return file.Close()
		})
	}

	return file
}
