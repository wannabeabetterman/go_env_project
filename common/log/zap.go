package log

import (
	"fmt"
	"github.com/getsentry/sentry-go"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"runtime/debug"
)

type (
	Level string
)

type Tags struct {
	TenantCode string `json:"tenantCode"`
	BizType    string `json:"bizType"`
	DeviceCode string `json:"deviceCode"`
	OrderNo    string `json:"orderNo"`
}

const (
	Debug Level = "debug"
	Info        = "info"
	Error       = "error"
	Panic       = "panic"
	Fatal       = "fatal"
)

type ZapLog struct {
	Logger *zap.SugaredLogger
}

func (zl *ZapLog) With(tags *Tags) Log {
	return &ZapLog{
		Logger: zl.Logger.With("tenantCode", tags.TenantCode, "bizType", tags.BizType, "deviceCode", tags.DeviceCode, "orderNo", tags.OrderNo),
	}
}

func (zl *ZapLog) Debug(args ...interface{}) {
	zl.Logger.Debug(args)
}

func (zl *ZapLog) Debugf(template string, args ...interface{}) {
	zl.Logger.Debugf(template, args)
}

func (zl *ZapLog) Info(args ...interface{}) {
	zl.Logger.Info(args)
}

func (zl *ZapLog) Infof(template string, args ...interface{}) {
	zl.Logger.Infof(template, args)
}

func (zl *ZapLog) Warn(args ...interface{}) {
	zl.Logger.Warn(args)
}

func (zl *ZapLog) Warnf(template string, args ...interface{}) {
	zl.Logger.Warnf(template, args)
}

func (zl *ZapLog) Error(args ...interface{}) {
	zl.Logger.Error(fmt.Sprintf("%s\n%s", fmt.Sprint(args...), string(debug.Stack())))
}

func (zl *ZapLog) Errorf(template string, args ...interface{}) {
	zl.Logger.Errorf(template, fmt.Sprintf("%s\n%s", fmt.Sprint(args...), string(debug.Stack())))
}

func (zl *ZapLog) Flush() {
	zl.Logger.Sync()
}

func Create(config Config) Log {
	enc := getEncoder()
	ws := getLogWriter(config)
	core := zapcore.NewCore(enc, ws, getLevel(config.Level))

	//sentryEnable = false  不接入sentry 直接构建zap
	if !config.SentryEnable {
		logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1), zap.AddStacktrace(zapcore.ErrorLevel)).Sugar()
		return &ZapLog{Logger: logger}
	}

	//sentryEnable = true 接入sentry
	err := sentry.Init(sentry.ClientOptions{
		Dsn:   config.SentryDsn,
		Debug: true,
	})
	if err != nil {
		return nil
	}
	sentryOptions := zap.WrapCore(func(core zapcore.Core) zapcore.Core {
		return zapcore.RegisterHooks(core, func(entry zapcore.Entry) error {
			if entry.Level == zapcore.ErrorLevel {
				//defer sentry.Flush(2 * time.Second)
				sentry.CaptureMessage(fmt.Sprintf("%s, Line No: %d :: %s", entry.Caller.File, entry.Caller.Line, entry.Message))
			}
			return nil
		})
	})
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1), zap.AddStacktrace(zapcore.ErrorLevel), sentryOptions).Sugar()

	return &ZapLog{Logger: logger}
}

func getLevel(l string) (zapLevel zapcore.Level) {
	switch Level(l) {
	case Debug:
		zapLevel = zapcore.DebugLevel
	case Info:
		zapLevel = zapcore.InfoLevel
	case Error:
		zapLevel = zapcore.ErrorLevel
	case Panic:
		zapLevel = zapcore.PanicLevel
	case Fatal:
		zapLevel = zapcore.FatalLevel
	default:
		zapLevel = zapcore.InfoLevel
	}
	return
}

func getEncoder() zapcore.Encoder {
	return zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		TimeKey:        "timestamp",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	})
}

func getLogWriter(config Config) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   config.FilePath + config.FileName,
		MaxSize:    config.MaxSize,
		MaxBackups: config.MaxBackups,
		MaxAge:     config.MaxAge,
		Compress:   config.Compress,
		LocalTime:  config.LocalTime,
	}

	var syncer zapcore.WriteSyncer
	if config.LogInConsole {
		//只打印err 2 console
		syncer = zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stderr), zapcore.AddSync(lumberJackLogger))
	} else {
		syncer = zapcore.NewMultiWriteSyncer(zapcore.AddSync(lumberJackLogger))
	}
	return syncer
}
