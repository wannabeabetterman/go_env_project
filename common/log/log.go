package log

type (
	Type string
)

const (
	Zap Type = "zap"
)

type Config struct {
	Type         string //日志类型 默认zap
	Level        string //日志等级
	FilePath     string //日志文件的位置
	FileName     string //日志名称
	MaxSize      int    //在进行切割之前，日志文件的最大大小（以MB为单位）
	MaxBackups   int    //保留旧文件的最大个数
	MaxAge       int    //保留旧文件的最大天数
	Compress     bool   //是否压缩/归档旧文件
	LocalTime    bool   //是否使用本地时间
	LogInConsole bool   //是否同时输出到控制台
	SentryEnable bool   //是否启用sentry
	SentryDsn    string //sentry-DSN
}

func Initialize(config Config) Log {

	switch Type(config.Type) {
	case Zap:
		return Create(config)
	default:
		return Create(config)
	}
}
