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