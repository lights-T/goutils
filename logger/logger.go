package logger

import (
	"fmt"
	"io"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/kardianos/service"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
	"gopkg.in/natefinch/lumberjack.v2"
)

var l *zerolog.Logger
var serviceSystemLog service.Logger

func newRollingFile(conf *Config) io.Writer {
	if err := os.MkdirAll(conf.Directory, 0744); err != nil {
		log.Error().Err(err).Str("path", conf.Directory).Msg("can't create log directory")
		return nil
	}

	fPath := path.Join(conf.Directory, conf.Filename)

	// 检查文件是否存在以及是否为空
	var needBOM bool
	if info, err := os.Stat(fPath); os.IsNotExist(err) {
		// 文件不存在 → 需要创建，并写入 BOM
		needBOM = true
	} else if err == nil && info.Size() == 0 {
		// 文件存在但为空 → 可以安全写入 BOM
		needBOM = true
	}

	// 如果需要 BOM，先创建文件并写入 BOM
	if needBOM {
		file, err := os.OpenFile(fPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
			log.Error().Err(err).Str("file", fPath).Msg("can't create log file for BOM")
			return nil
		}
		// 写入 UTF-8 BOM
		_, err = file.Write([]byte{0xEF, 0xBB, 0xBF})
		if err != nil {
			log.Error().Err(err).Str("file", fPath).Msg("failed to write BOM")
		}
		file.Close()
	}

	// 返回 lumberjack logger（它会追加写入）
	return &lumberjack.Logger{
		Filename:   fPath,
		MaxBackups: conf.MaxBackups, // files
		MaxSize:    conf.MaxSize,    // megabytes
		MaxAge:     conf.MaxAge,     // days
		Compress:   true,            // disabled by default
	}
}

type Config struct {
	// Enable console logging
	ConsoleLoggingEnabled bool

	// EncodeLogsAsJson makes the log framework log JSON
	EncodeLogsAsJson bool
	// FileLoggingEnabled makes the framework log to a file
	// the fields below can be skipped if this value is false!
	FileLoggingEnabled bool
	// Directory to log to to when filelogging is enabled
	Directory string
	// Filename is the name of the logfile which will be placed inside the directory
	Filename string
	// MaxSize the max size in MB of the logfile before it's rolled
	MaxSize int
	// MaxBackups the max number of rolled files to keep
	MaxBackups int
	// MaxAge the max age in days to keep a logfile
	MaxAge           int
	Debug            bool
	ServiceSystemLog service.Logger
}

func New(conf *Config) *zerolog.Logger {
	var writers []io.Writer
	if conf.ConsoleLoggingEnabled {
		writers = append(writers, zerolog.ConsoleWriter{Out: os.Stderr})
	}
	if conf.FileLoggingEnabled {
		writers = append(writers, newRollingFile(conf))
	}
	mw := io.MultiWriter(writers...)
	level := strings.ToLower(os.Getenv("MICRO_LOG_LEVEL"))
	if conf.Debug {
		level = strings.ToLower("debug")
	}
	if level == "debug" {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	nl := zerolog.New(mw).With().Timestamp().Stack().CallerWithSkipFrameCount(3).Logger()
	zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
		return WashPath(file) + ":" + strconv.Itoa(line)
	}
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	return &nl
}

// WashPath 路径脱敏
func WashPath(s string) string {
	lIndex := strings.LastIndex(s, "/")
	if lIndex < 0 {
		return s
	}
	if lIndex+1 > len(s)-1 {
		return s[lIndex:]
	}

	return s[lIndex+1:]
}

func InitLogger(conf *Config) *zerolog.Logger {
	l = New(conf)
	serviceSystemLog = conf.ServiceSystemLog
	return l
}

func Debugf(format string, args ...interface{}) {
	l.Debug().Msgf(format, args...)
	if serviceSystemLog != nil {
		_ = serviceSystemLog.Infof(format, args...)
	}
}

func Debug(args interface{}) {
	switch args.(type) {
	case string:
		l.Debug().Msg(args.(string))
		if serviceSystemLog != nil {
			_ = serviceSystemLog.Info(args.(string))
		}
	case error:
		l.Debug().Msg(args.(error).Error())
		if serviceSystemLog != nil {
			_ = serviceSystemLog.Info(args.(error).Error())
		}
	default:
		l.Debug().Msg(fmt.Sprintf("%v", args))
		if serviceSystemLog != nil {
			_ = serviceSystemLog.Info(fmt.Sprintf("%v", args))
		}
	}
}

func Infof(format string, args ...interface{}) {
	l.Info().Msgf(format, args...)
	if serviceSystemLog != nil {
		_ = serviceSystemLog.Infof(format, args...)
	}
}

func Info(args interface{}) {
	switch args.(type) {
	case string:
		l.Info().Msg(args.(string))
		if serviceSystemLog != nil {
			_ = serviceSystemLog.Info(args.(string))
		}
	case error:
		l.Info().Msg(args.(error).Error())
		if serviceSystemLog != nil {
			_ = serviceSystemLog.Info(args.(error).Error())
		}
	default:
		l.Info().Msg(fmt.Sprintf("%v", args))
		if serviceSystemLog != nil {
			_ = serviceSystemLog.Info(fmt.Sprintf("%v", args))
		}
	}
}

func Warnf(format string, args ...interface{}) {
	l.Warn().Msgf(format, args...)
	if serviceSystemLog != nil {
		_ = serviceSystemLog.Warningf(format, args...)
	}
}

func Warn(args interface{}) {
	switch args.(type) {
	case string:
		l.Warn().Msg(args.(string))
		if serviceSystemLog != nil {
			_ = serviceSystemLog.Warning(args.(string))
		}
	case error:
		l.Warn().Msg(args.(error).Error())
		if serviceSystemLog != nil {
			_ = serviceSystemLog.Warning(args.(error).Error())
		}
	default:
		l.Warn().Msg(fmt.Sprintf("%v", args))
		if serviceSystemLog != nil {
			_ = serviceSystemLog.Warning(fmt.Sprintf("%v", args))
		}
	}
}

func Errorf(format string, args ...interface{}) {
	l.Error().Msgf(format, args...)
	if serviceSystemLog != nil {
		_ = serviceSystemLog.Errorf(format, args...)
	}
}

func Error(err interface{}) {
	switch err.(type) {
	case string:
		l.Error().Msg(err.(string))
		if serviceSystemLog != nil {
			_ = serviceSystemLog.Error(err.(string))
		}
	case error:
		l.Err(err.(error)).Send()
		if serviceSystemLog != nil {
			_ = serviceSystemLog.Error(err.(error))
		}
	default:
		errStr := fmt.Sprintf("%v", err)
		l.Error().Msg(errStr)
		if serviceSystemLog != nil {
			_ = serviceSystemLog.Error(errStr)
		}
	}
}

func Fatalf(format string, args ...interface{}) {
	l.Fatal().Msgf(format, args...)
	if serviceSystemLog != nil {
		_ = serviceSystemLog.Errorf(format, args...)
	}
}

func Fatal(args interface{}) {
	switch args.(type) {
	case string:
		l.Fatal().Msg(args.(string))
		if serviceSystemLog != nil {
			_ = serviceSystemLog.Error(args.(string))
		}
	case error:
		l.Fatal().Msg(args.(error).Error())
		if serviceSystemLog != nil {
			_ = serviceSystemLog.Error(args.(error).Error())
		}
	default:
		l.Fatal().Msg(fmt.Sprintf("%v", args))
		if serviceSystemLog != nil {
			_ = serviceSystemLog.Error(fmt.Sprintf("%v", args))
		}
	}
}
