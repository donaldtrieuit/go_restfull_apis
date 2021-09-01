package utils

import (
	"go_restfull_api/configs"
	"github.com/lestrrat-go/file-rotatelogs"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"time"
)

type Logger struct {

}

func (l *Logger) Trace(args ...interface{}) {
	log.Trace(args...)
}

func (l *Logger) Debug(args ...interface{}) {
	log.Debug(args...)
}

func (l *Logger) Print(args ...interface{}) {
	log.Print(args...)
}

func (l *Logger) Info(args ...interface{}) {
	log.Info(args...)
}

func (l *Logger) Warn(args ...interface{}) {
	log.Warn(args...)
}

func (l *Logger) Warning(args ...interface{}) {
	log.Warning(args...)
}

func (l *Logger) Error(args ...interface{}) {
	log.Error(args...)
}

func (l *Logger) Panic(args ...interface{}) {
	log.Panic(args...)
}

func (l *Logger) Fatal(args ...interface{}) {
	log.Fatal(args...)
}

func (l *Logger) Tracef(format string, args ...interface{}) {
	log.Tracef(format, args...)
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	log.Debugf(format, args...)
}

func (l *Logger) Printf(format string, args ...interface{}) {
	log.Printf(format, args...)
}

func (l *Logger) Infof(format string, args ...interface{}) {
	log.Infof(format, args...)
}

func (l *Logger) Warnf(format string, args ...interface{}) {
	log.Warnf(format, args...)
}

func (l *Logger) Warningf(format string, args ...interface{}) {
	log.Warningf(format, args...)
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	log.Errorf(format, args...)
}

func (l *Logger) Panicf(format string, args ...interface{}) {
	log.Panicf(format, args...)
}

func (l *Logger) Fatalf(format string, args ...interface{}) {
	log.Fatalf(format, args...)
}

func (l *Logger) Traceln(args ...interface{}) {
	log.Traceln(args...)
}

func (l *Logger) Debugln(args ...interface{}) {
	log.Debugln(args...)
}

func (l *Logger) Println(args ...interface{}) {
	log.Println(args...)
}

func (l *Logger) Infoln(args ...interface{}) {
	log.Infoln(args...)
}

func (l *Logger) Warnln(args ...interface{}) {
	log.Warnln(args...)
}

func (l *Logger) Warningln(args ...interface{}) {
	log.Warningln(args...)
}

func (l *Logger) Errorln(args ...interface{}) {
	log.Errorln(args...)
}

func (l *Logger) Panicln(args ...interface{}) {
	log.Panicln(args...)
}

func (l *Logger) Fatalln(args ...interface{}) {
	log.Fatalln(args...)
}

var logger *Logger

func InitLogger() {
	// TODO init log
	cfg := configs.GetConfig()
	log.SetLevel(log.Level(cfg.GetUint32("logging.level")))
	outs := []io.Writer{os.Stdout}
	if cfg.GetBool("logging.log_to_file") {
		file := cfg.GetString("logging.file")
		fileLogWriter, err := rotatelogs.New(
			file,
			rotatelogs.WithMaxAge(time.Duration(cfg.GetInt("logging.rotation_max_age")) * 24 * time.Hour),
			rotatelogs.WithRotationTime(time.Duration(cfg.GetInt("logging.rotation_time")) * time.Hour),
			rotatelogs.WithRotationSize(cfg.GetInt64("logging.rotation_size")),
		)
		if err != nil {
			panic(err)
		}
		outs = append(outs, fileLogWriter)
	}
	log.SetOutput(io.MultiWriter(outs...))

	logger = &Logger{}
}

func GetLogger() *Logger {
	return logger
}