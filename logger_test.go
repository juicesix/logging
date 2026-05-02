package logging

import (
	"testing"
	"time"
)

type stringArrayEncoder struct {
	values []string
}

func (e *stringArrayEncoder) AppendBool(bool)              {}
func (e *stringArrayEncoder) AppendByteString([]byte)      {}
func (e *stringArrayEncoder) AppendComplex128(complex128)  {}
func (e *stringArrayEncoder) AppendComplex64(complex64)    {}
func (e *stringArrayEncoder) AppendDuration(time.Duration) {}
func (e *stringArrayEncoder) AppendFloat64(float64)        {}
func (e *stringArrayEncoder) AppendFloat32(float32)        {}
func (e *stringArrayEncoder) AppendInt(int)                {}
func (e *stringArrayEncoder) AppendInt64(int64)            {}
func (e *stringArrayEncoder) AppendInt32(int32)            {}
func (e *stringArrayEncoder) AppendInt16(int16)            {}
func (e *stringArrayEncoder) AppendInt8(int8)              {}
func (e *stringArrayEncoder) AppendUint(uint)              {}
func (e *stringArrayEncoder) AppendUint64(uint64)          {}
func (e *stringArrayEncoder) AppendUint32(uint32)          {}
func (e *stringArrayEncoder) AppendUint16(uint16)          {}
func (e *stringArrayEncoder) AppendUint8(uint8)            {}
func (e *stringArrayEncoder) AppendUintptr(uintptr)        {}
func (e *stringArrayEncoder) AppendTime(time.Time)         {}

func (e *stringArrayEncoder) AppendString(value string) {
	e.values = append(e.values, value)
}

func TestDefaultLoggersAreRegistered(t *testing.T) {
	for _, name := range []string{
		DefaultLoggerName,
		SlowLoggerName,
		GenLoggerName,
		CrashLoggerName,
		BalanceLoggerName,
	} {
		if Log(name) == nil {
			t.Fatalf("expected logger %q to be registered", name)
		}
	}
}

func TestCommonLogOpenStatus(t *testing.T) {
	OpenCommonLog()
	if !checkOpenStatus() {
		t.Fatal("expected common log to be open")
	}

	CloseCommonLog()
	if checkOpenStatus() {
		t.Fatal("expected common log to be closed")
	}

	OpenCommonLog()
}

func TestIsHourRotate(t *testing.T) {
	if !isHourRotate(HOUR_ROTATE) {
		t.Fatal("expected hour rotate to be recognized")
	}
	if isHourRotate(DAY_ROTATE) {
		t.Fatal("did not expect day rotate to be treated as hour rotate")
	}
}

func TestMillSecondTimeEncoder(t *testing.T) {
	enc := &stringArrayEncoder{}
	when := time.Date(2026, 4, 30, 9, 32, 45, 123456000, time.UTC)

	MillSecondTimeEncoder(when, enc)

	if len(enc.values) != 1 {
		t.Fatalf("expected one encoded value, got %d", len(enc.values))
	}
	if enc.values[0] != "2026-04-30 09:32:45.123" {
		t.Fatalf("unexpected encoded time: %q", enc.values[0])
	}
}

// 针对秒级时间确认编码结果仍包含毫秒部分
func TestMillSecondTimeEncoder_Second(t *testing.T) {
	enc := &stringArrayEncoder{}
	when := time.Date(2026, 4, 30, 9, 32, 45, 123456000, time.UTC)

	MillSecondTimeEncoder(when, enc)

	if len(enc.values) != 1 {
		t.Fatalf("expected one encoded value, got %d", len(enc.values))
	}
	if enc.values[0] != "2026-04-30 09:32:45.123" {
		t.Fatalf("unexpected encoded time: %q", enc.values[0])
	}

}

// 增加一行注释

// 针对秒级时间确认编码结果仍包含毫秒部分
func TestMillSecondTimeEncoder_Second(t *testing.T) {
	enc := &stringArrayEncoder{}
	when := time.Date(2026, 4, 30, 9, 32, 45, 123456000, time.UTC)

	MillSecondTimeEncoder(when, enc)

	if len(enc.values) != 1 {
		t.Fatalf("expected one encoded value, got %d", len(enc.values))
	}
	if enc.values[0] != "2026-04-30 09:32:45.123" {
		t.Fatalf("unexpected encoded time: %q", enc.values[0])
	}

}

// 增加一行注释

// git stash test
