package byteflow

import (
	"fmt"
	"time"
)

// Rate represents a data transfer rate in bytes per second.
type Rate int64

const (
	// Bps represents one byte per second.
	Bps           Rate = 1
	BytePerSecond      = Bps

	// KiBps, MiBps, GiBps and TiBps represent binary byte rates per second.
	KiBps = Rate(KiB)
	MiBps = Rate(MiB)
	GiBps = Rate(GiB)
	TiBps = Rate(TiB)

	// KBps, MBps, GBps and TBps represent decimal byte rates per second.
	KBps = Rate(KB)
	MBps = Rate(MB)
	GBps = Rate(GB)
	TBps = Rate(TB)
)

const (
	bitPerByte = 8

	kilobit = 1000
	megabit = 1000 * kilobit
	gigabit = 1000 * megabit
	terabit = 1000 * gigabit
)

// BytesPerSecond returns a Rate from a raw bytes-per-second value.
func BytesPerSecond(value int64) Rate {
	return Rate(value)
}

// BytesPerSecond returns the raw bytes-per-second value of the rate.
func (rate Rate) BytesPerSecond() int64 {
	return int64(rate)
}

// KilobytesPerSecond returns the rate as decimal kilobytes per second.
func (rate Rate) KilobytesPerSecond() float64 {
	return float64(rate) / float64(KBps)
}

// MegabytesPerSecond returns the rate as decimal megabytes per second.
func (rate Rate) MegabytesPerSecond() float64 {
	return float64(rate) / float64(MBps)
}

// GigabytesPerSecond returns the rate as decimal gigabytes per second.
func (rate Rate) GigabytesPerSecond() float64 {
	return float64(rate) / float64(GBps)
}

// TerabytesPerSecond returns the rate as decimal terabytes per second.
func (rate Rate) TerabytesPerSecond() float64 {
	return float64(rate) / float64(TBps)
}

// KibibytesPerSecond returns the rate as binary kibibytes per second.
func (rate Rate) KibibytesPerSecond() float64 {
	return float64(rate) / float64(KiBps)
}

// MebibytesPerSecond returns the rate as binary mebibytes per second.
func (rate Rate) MebibytesPerSecond() float64 {
	return float64(rate) / float64(MiBps)
}

// GibibytesPerSecond returns the rate as binary gibibytes per second.
func (rate Rate) GibibytesPerSecond() float64 {
	return float64(rate) / float64(GiBps)
}

// TebibytesPerSecond returns the rate as binary tebibytes per second.
func (rate Rate) TebibytesPerSecond() float64 {
	return float64(rate) / float64(TiBps)
}

// BitsPerSecond returns the rate as bits per second.
func (rate Rate) BitsPerSecond() float64 {
	return float64(rate) * bitPerByte
}

// KilobitsPerSecond returns the rate as decimal kilobits per second.
func (rate Rate) KilobitsPerSecond() float64 {
	return rate.BitsPerSecond() / kilobit
}

// MegabitsPerSecond returns the rate as decimal megabits per second.
func (rate Rate) MegabitsPerSecond() float64 {
	return rate.BitsPerSecond() / megabit
}

// GigabitsPerSecond returns the rate as decimal gigabits per second.
func (rate Rate) GigabitsPerSecond() float64 {
	return rate.BitsPerSecond() / gigabit
}

// TerabitsPerSecond returns the rate as decimal terabits per second.
func (rate Rate) TerabitsPerSecond() float64 {
	return rate.BitsPerSecond() / terabit
}

// String returns the rate formatted using binary byte units per second.
//
// The result uses B/s, KiB/s, MiB/s, GiB/s or TiB/s depending on the rate value.
// Fractional values are rounded to two decimal places and trailing zeros are removed.
func (rate Rate) String() string {
	return rate.BinaryString()
}

// BinaryString returns the rate formatted using binary byte units per second.
//
// The result uses B/s, KiB/s, MiB/s, GiB/s or TiB/s depending on the absolute rate value.
// Fractional values are rounded to two decimal places and trailing zeros are removed.
func (rate Rate) BinaryString() string {
	size := Size(rate)

	return size.BinaryString() + "/s"
}

// DecimalString returns the rate formatted using decimal byte units per second.
//
// The result uses B/s, KB/s, MB/s, GB/s or TB/s depending on the absolute rate value.
// Fractional values are rounded to two decimal places and trailing zeros are removed.
func (rate Rate) DecimalString() string {
	size := Size(rate)

	return size.DecimalString() + "/s"
}

// BitString returns the rate formatted using decimal bit units per second.
//
// The result uses bps, Kbps, Mbps, Gbps or Tbps depending on the absolute rate value.
// Fractional values are rounded to two decimal places and trailing zeros are removed.
func (rate Rate) BitString() string {
	unitValue, unitName := rate.bitFormatUnit()

	if unitName == "bps" {
		return fmt.Sprintf("%s %s", formatFloat(rate.BitsPerSecond()), unitName)
	}

	formattedValue := formatFloat(rate.BitsPerSecond() / unitValue)

	return fmt.Sprintf("%s %s", formattedValue, unitName)
}

func (rate Rate) bitFormatUnit() (float64, string) {
	absoluteBitsPerSecond := rate.absoluteBitsPerSecond()

	switch {
	case absoluteBitsPerSecond >= terabit:
		return terabit, "Tbps"
	case absoluteBitsPerSecond >= gigabit:
		return gigabit, "Gbps"
	case absoluteBitsPerSecond >= megabit:
		return megabit, "Mbps"
	case absoluteBitsPerSecond >= kilobit:
		return kilobit, "Kbps"
	default:
		return 1, "bps"
	}
}

func (rate Rate) absoluteBitsPerSecond() float64 {
	bitsPerSecond := rate.BitsPerSecond()

	if bitsPerSecond < 0 {
		return -bitsPerSecond
	}

	return bitsPerSecond
}

// PerSecond calculates the transfer rate for a given size and duration.
//
// If duration is zero or negative, it returns 0.
func PerSecond(size Size, duration time.Duration) Rate {
	if duration <= 0 {
		return 0
	}

	bytesPerSecond := float64(size) / duration.Seconds()

	return Rate(bytesPerSecond)
}

// DurationFor calculates the duration needed to transfer a given size at a given rate.
//
// If rate is zero or negative, it returns 0.
func DurationFor(size Size, rate Rate) time.Duration {
	if rate <= 0 {
		return 0
	}

	seconds := float64(size) / float64(rate)

	return time.Duration(seconds * float64(time.Second))
}

// TransferredIn calculates how much data can be transferred at a given rate and duration.
//
// If duration is zero or negative, it returns 0.
func TransferredIn(rate Rate, duration time.Duration) Size {
	if duration <= 0 {
		return 0
	}

	bytesTransferred := float64(rate) * duration.Seconds()

	return Size(bytesTransferred)
}
