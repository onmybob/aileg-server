package utils

import (
	"io"
	"os"
	"time"

	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
)
 
var (
	Log  zerolog.Logger
)


func init() {
}
 
// Config 可用在配置文件中
type Config struct {
	Filename   string            // 日志文件
	MaxSize    int               // megabytes
	MaxBackups int               // MaxBackups
	MaxAge     int               // days
	Fields     map[string]string // slog的初始化字段(session)
	IsJSON     bool              // 默认是非json格式
}


func init(){
	c := Config{
		Filename: "./logs/"+time.Now().Format("2006-01-02")+".log",
		IsJSON:   false,
		Fields:map[string]string{
			"S": "aileg-web",
		},
	}
	multi:=zerolog.MultiLevelWriter(ConsoleWriter(),FileWriter(c))
	zc:= zerolog.New(multi).With().Caller().Timestamp()
	for k, v := range c.Fields {
		zc = zc.Str(k, v)
	}
	Log=zc.Logger()

}

// New 日常使用的log, 建议使用json效率更高
func FileWriter(c Config) io.Writer {
	// init zerolog format
	zerolog.TimeFieldFormat = "15:04:05.000"
	zerolog.DurationFieldInteger = true
	zerolog.TimestampFieldName = "timestamp"
	zerolog.DurationFieldUnit = time.Millisecond

	out := &lumberjack.Logger{
		Filename:   c.Filename,
		MaxSize:    c.MaxSize,    // megabytes
		MaxBackups: c.MaxBackups, // MaxBackups
		MaxAge:     c.MaxAge,     // days
		LocalTime:  true,         // 这个需要设置, 不然日志文件的名字就是UTC时间
	}

	return out
} 

func ConsoleWriter() io.Writer {
	consoleWriter:=zerolog.ConsoleWriter{Out:os.Stdout,TimeFormat: "15:04:05", NoColor: false}
	consoleWriter.Out = os.Stderr

	 
	return consoleWriter
}