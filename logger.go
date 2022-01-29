package logging

import (
	"io"
	"os"
	"sync"

	"github.com/juicesix/rolling"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	*zap.SugaredLogger
	path         string
	dir          string
	rolling      rolling.RollingFormat
	rollingFiles []io.Writer
	loglevel     zap.AtomicLevel
	prefix       string
	encoderCfg   zapcore.EncoderConfig
	callSkip     int
}

var defaultEncoderConfig = zapcore.EncoderConfig{
	MessageKey:     "msg",
	LevelKey:       "level",
	TimeKey:        "time",
	NameKey:        "logger",
	CallerKey:      "caller",
	StacktraceKey:  "stack",
	LineEnding:     zapcore.DefaultLineEnding,
	EncodeLevel:    zapcore.CapitalColorLevelEncoder,
	EncodeTime:     MillSecondTimeEncoder,
	EncodeDuration: zapcore.StringDurationEncoder,
	EncodeCaller:   zapcore.ShortCallerEncoder,
	EncodeName:     zapcore.FullNameEncoder,
}

var (
	_defaultLogger  *Logger
	_jsonDataLogger *Logger
	hostIP          string
)

// Logger name for default loggers
const (
	DefaultLoggerName = "_default"
	SlowLoggerName    = "_slow"
	GenLoggerName     = "_gen"
	CrashLoggerName   = "_crash"
	BalanceLoggerName = "_balance"
)

func init() {
	_defaultLogger = New()
	logs[DefaultLoggerName] = _defaultLogger
	logs[SlowLoggerName] = slowlog
	logs[GenLoggerName] = genlog
	logs[CrashLoggerName] = crashlog
	logs[BalanceLoggerName] = balancelog
}

var logs = map[string]*Logger{}
var logsMtx sync.RWMutex

func Log(name string) *Logger {
	logsMtx.RLock()
	defer logsMtx.RUnlock()
	return logs[name]
}

func New() *Logger {
	cfg := defaultEncoderConfig
	lvl := zap.NewAtomicLevelAt(zap.DebugLevel)
	return &Logger{
		SugaredLogger: zap.New(zapcore.NewCore(zapcore.NewConsoleEncoder(cfg), zapcore.Lock(os.Stderr), lvl)).WithOptions(zap.AddCaller(), zap.AddCallerSkip(1)).Sugar(),
		path:          "",
		dir:           "",
		rolling:       rolling.DailyRolling,
		rollingFiles:  nil,
		loglevel:      lvl,
		prefix:        "",
		encoderCfg:    cfg,
	}
}
