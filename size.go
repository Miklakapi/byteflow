package byteflow

import (
	"fmt"
	"strings"
)

// Size represents a data size in bytes.
type Size int64

const (
	// Byte represents one byte.
	B    Size = 1
	Byte      = B

	// KiB, MiB, GiB and TiB represent binary byte units.
	KiB = 1024 * Byte
	MiB = 1024 * KiB
	GiB = 1024 * MiB
	TiB = 1024 * GiB

	// KB, MB, GB and TB represent decimal byte units.
	KB = 1000 * Byte
	MB = 1000 * KB
	GB = 1000 * MB
	TB = 1000 * GB
)

// Bytes returns a Size from a raw byte value.
func Bytes(value int64) Size {
	return Size(value)
}

// Bytes returns the raw byte value of the size.
func (size Size) Bytes() int64 {
	return int64(size)
}

// Kilobytes returns the size as decimal kilobytes.
func (size Size) Kilobytes() float64 {
	return float64(size) / float64(KB)
}

// Megabytes returns the size as decimal megabytes.
func (size Size) Megabytes() float64 {
	return float64(size) / float64(MB)
}

// Gigabytes returns the size as decimal gigabytes.
func (size Size) Gigabytes() float64 {
	return float64(size) / float64(GB)
}

// Terabytes returns the size as decimal terabytes.
func (size Size) Terabytes() float64 {
	return float64(size) / float64(TB)
}

// Kibibytes returns the size as binary kibibytes.
func (size Size) Kibibytes() float64 {
	return float64(size) / float64(KiB)
}

// Mebibytes returns the size as binary mebibytes.
func (size Size) Mebibytes() float64 {
	return float64(size) / float64(MiB)
}

// Gibibytes returns the size as binary gibibytes.
func (size Size) Gibibytes() float64 {
	return float64(size) / float64(GiB)
}

// Tebibytes returns the size as binary tebibytes.
func (size Size) Tebibytes() float64 {
	return float64(size) / float64(TiB)
}

// String returns the size formatted using binary byte units.
//
// The result uses B, KiB, MiB, GiB or TiB depending on the size value.
// Fractional values are rounded to two decimal places and trailing zeros are removed.
func (size Size) String() string {
	return size.BinaryString()
}

// BinaryString returns the size formatted using binary byte units.
//
// The result uses B, KiB, MiB, GiB or TiB depending on the absolute size value.
// Fractional values are rounded to two decimal places and trailing zeros are removed.
func (size Size) BinaryString() string {
	unitValue, unitName := size.binaryFormatUnit()

	return size.formatWithUnit(unitValue, unitName)
}

// DecimalString returns the size formatted using decimal byte units.
//
// The result uses B, KB, MB, GB or TB depending on the absolute size value.
// Fractional values are rounded to two decimal places and trailing zeros are removed.
func (size Size) DecimalString() string {
	unitValue, unitName := size.decimalFormatUnit()

	return size.formatWithUnit(unitValue, unitName)
}

func (size Size) binaryFormatUnit() (Size, string) {
	absoluteSize := size.absoluteFloat()

	switch {
	case absoluteSize >= float64(TiB):
		return TiB, "TiB"
	case absoluteSize >= float64(GiB):
		return GiB, "GiB"
	case absoluteSize >= float64(MiB):
		return MiB, "MiB"
	case absoluteSize >= float64(KiB):
		return KiB, "KiB"
	default:
		return B, "B"
	}
}

func (size Size) decimalFormatUnit() (Size, string) {
	absoluteSize := size.absoluteFloat()

	switch {
	case absoluteSize >= float64(TB):
		return TB, "TB"
	case absoluteSize >= float64(GB):
		return GB, "GB"
	case absoluteSize >= float64(MB):
		return MB, "MB"
	case absoluteSize >= float64(KB):
		return KB, "KB"
	default:
		return B, "B"
	}
}

func (size Size) formatWithUnit(unitValue Size, unitName string) string {
	if unitName == "B" {
		return fmt.Sprintf("%d %s", size, unitName)
	}

	formattedValue := formatFloat(float64(size) / float64(unitValue))

	return fmt.Sprintf("%s %s", formattedValue, unitName)
}

func (size Size) absoluteFloat() float64 {
	if size < 0 {
		return -float64(size)
	}

	return float64(size)
}

func formatFloat(value float64) string {
	formattedValue := fmt.Sprintf("%.2f", value)
	formattedValue = strings.TrimRight(formattedValue, "0")
	formattedValue = strings.TrimRight(formattedValue, ".")

	return formattedValue
}
