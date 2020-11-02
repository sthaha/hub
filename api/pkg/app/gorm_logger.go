package app

import (
	"fmt"
	"strings"
	"time"

	"github.com/tektoncd/hub/api/gen/log"
	glog "gorm.io/gorm/logger"
)

//func New(writer Writer, config Config) Interface {
func gormLoggerForMode(mode EnvMode, l *log.Logger) (glog.Writer, glog.Config) {
	//glog := glog.New(gormLoggerForMode(mode, l))
	if mode == Production {
		return &prodWriter{l}, glog.Config{
			SlowThreshold: 100 * time.Millisecond,
			LogLevel:      glog.Info,
			Colorful:      false, // only Colorful returns data in an array
		}
	}

	return &devWriter{l}, glog.Config{
		SlowThreshold: 100 * time.Millisecond,
		LogLevel:      glog.Info,
		Colorful:      true,
	}
}

func newGormLogger(mode EnvMode, l *log.Logger) glog.Interface {
	if mode == Production {
		return glog.New(&prodWriter{l}, glog.Config{
			SlowThreshold: 100 * time.Millisecond,
			LogLevel:      glog.Info,
		})
	}

	return glog.New(&devWriter{l}, glog.Config{
		SlowThreshold: 100 * time.Millisecond,
		LogLevel:      glog.Info,
		Colorful:      true,
	})
}

// adaptor for gorm logger interface
type prodWriter struct {
	log *log.Logger
}

func (w *prodWriter) Printf(format string, data ...interface{}) {
	//config := zap.NewDevelopmentConfig()
	//config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	//l, _ := config.Build()

	//sl := w.log.SugaredLogger

	//l.Sugar().Info(msg + "\n\n" + spew.Sdump(data))
	//w.log.Infof(" db >>> "+msg, data...)

	//log := l.log.SugaredLogger

	fields := strings.Fields(strings.Replace(format, "\n", " ", -1))
	log := w.log.With("file", data[0])

	data = data[1:]
	fields = fields[1:]

	msg := ""

	for i, d := range data {
		//l.Sugar().Infof(" >>>>>> field: %v: %s -> %v | %v ", i, fields[i], reflect.TypeOf(d), d)
		switch d.(type) {
		case error:
			log = log.With("db-error", d)
		case float64:
			log = log.With("duration", fmt.Sprintf(fields[i], d))
		case int64:
			log = log.With("rows", d)
		case string:
			if i == len(data)-1 {
				msg = d.(string)
			} else {
				log = log.With("unknown", d)
			}
		default:
			log = log.With("unknown", d)
		}
	}

	log.Info(msg)
}

type devWriter struct {
	log *log.Logger
}

func (w *devWriter) Printf(msg string, data ...interface{}) {
	w.log.SugaredLogger.Infof(strings.Replace(msg, "%s ", "%s\n", 1), data...)
}
