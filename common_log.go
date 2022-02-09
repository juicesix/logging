package logging

var slowlog *Logger = New()
var genlog *Logger = New()
var crashlog *Logger = New()
var balancelog *Logger = New()

var (
	DAY_ROTATE  = "day"
	HOUR_ROTATE = "hour"
)

type CommonLogConfig struct {
	Pathprefix      string
	Rotate          string
	GenLogLevel     string
	BalanceLogLevel string
}

var isInit bool = false
var trunon bool = true

func init() {
	isInit = false
}

func isHourRotate(rotate string) bool {
	if rotate == HOUR_ROTATE {
		return true
	}
	return false
}

func CloseCommonLog() {
	trunon = false
}

func OpenCommonLog() {
	trunon = true
}

func checkOpenStatus() bool {
	if trunon == false {
		return false
	}
	return true
}

func checkNeedLog() bool {
	if isInit == false {
		return false
	}

	return true
}

func GenLog(v ...interface{}) {
	if !checkOpenStatus() {
		return
	}

	if !checkNeedLog() {
		_defaultLogger.Debug(v...)
		return
	}

	genlog.Debug(v...)
}

func GenLogf(format string, v ...interface{}) {

	if !checkOpenStatus() {
		return
	}

	if !checkNeedLog() {
		_defaultLogger.Debugf(format, v...)
		return
	}

	genlog.Debugf(format, v...)
}

func SlowLog(v ...interface{}) {

	if !checkOpenStatus() {
		return
	}

	if !checkNeedLog() {
		_defaultLogger.Debug(v...)
		return
	}
	slowlog.Debug(v...)
}

func SlowLogf(format string, v ...interface{}) {

	if !checkOpenStatus() {
		return
	}

	if !checkNeedLog() {
		_defaultLogger.Debugf(format, v...)
		return
	}
	slowlog.Debugf(format, v...)
}

func CrashLog(v ...interface{}) {

	if !checkOpenStatus() {
		return
	}

	if !checkNeedLog() {
		_defaultLogger.Debug(v...)
		return
	}

	crashlog.Debug(v...)
}

func CrashLogf(format string, v ...interface{}) {

	if !checkOpenStatus() {
		return
	}

	if !checkNeedLog() {
		_defaultLogger.Debugf(format, v...)
		return
	}
	crashlog.Debugf(format, v...)
}

func BalanceLog(v ...interface{}) {

	if !checkOpenStatus() {
		return
	}
	if !checkNeedLog() {
		_defaultLogger.Debug(v...)
		return
	}
	balancelog.Debug(v...)
}

func BalanceLogf(format string, v ...interface{}) {

	if !checkOpenStatus() {
		return
	}
	if !checkNeedLog() {
		_defaultLogger.Debugf(format, v...)
		return
	}
	balancelog.Debugf(format, v...)
}
