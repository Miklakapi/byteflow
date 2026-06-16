package byteflow

import (
	"testing"
	"time"
)

func TestBytesPerSecond(t *testing.T) {
	rate := BytesPerSecond(1536)

	if rate != 1536*Bps {
		t.Fatalf("expected rate to be %d, got %d", 1536*Bps, rate)
	}

	if rate.BytesPerSecond() != 1536 {
		t.Fatalf("expected bytes per second to be %d, got %d", 1536, rate.BytesPerSecond())
	}
}

func TestRateDecimalConversions(t *testing.T) {
	tests := []struct {
		name     string
		rate     Rate
		expected float64
		actual   func(Rate) float64
	}{
		{
			name:     "kilobytes per second",
			rate:     1500 * Bps,
			expected: 1.5,
			actual:   Rate.KilobytesPerSecond,
		},
		{
			name:     "megabytes per second",
			rate:     1500 * KBps,
			expected: 1.5,
			actual:   Rate.MegabytesPerSecond,
		},
		{
			name:     "gigabytes per second",
			rate:     1500 * MBps,
			expected: 1.5,
			actual:   Rate.GigabytesPerSecond,
		},
		{
			name:     "terabytes per second",
			rate:     1500 * GBps,
			expected: 1.5,
			actual:   Rate.TerabytesPerSecond,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actualValue := test.actual(test.rate)

			if actualValue != test.expected {
				t.Fatalf("expected %v, got %v", test.expected, actualValue)
			}
		})
	}
}

func TestRateBinaryConversions(t *testing.T) {
	tests := []struct {
		name     string
		rate     Rate
		expected float64
		actual   func(Rate) float64
	}{
		{
			name:     "kibibytes per second",
			rate:     1536 * Bps,
			expected: 1.5,
			actual:   Rate.KibibytesPerSecond,
		},
		{
			name:     "mebibytes per second",
			rate:     1536 * KiBps,
			expected: 1.5,
			actual:   Rate.MebibytesPerSecond,
		},
		{
			name:     "gibibytes per second",
			rate:     1536 * MiBps,
			expected: 1.5,
			actual:   Rate.GibibytesPerSecond,
		},
		{
			name:     "tebibytes per second",
			rate:     1536 * GiBps,
			expected: 1.5,
			actual:   Rate.TebibytesPerSecond,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actualValue := test.actual(test.rate)

			if actualValue != test.expected {
				t.Fatalf("expected %v, got %v", test.expected, actualValue)
			}
		})
	}
}

func TestRateString(t *testing.T) {
	tests := []struct {
		name     string
		rate     Rate
		expected string
	}{
		{
			name:     "bytes per second",
			rate:     512 * Bps,
			expected: "512 B/s",
		},
		{
			name:     "kibibytes per second",
			rate:     1536 * Bps,
			expected: "1.5 KiB/s",
		},
		{
			name:     "mebibytes per second",
			rate:     1536 * KiBps,
			expected: "1.5 MiB/s",
		},
		{
			name:     "gibibytes per second",
			rate:     1536 * MiBps,
			expected: "1.5 GiB/s",
		},
		{
			name:     "tebibytes per second",
			rate:     1536 * GiBps,
			expected: "1.5 TiB/s",
		},
		{
			name:     "negative kibibytes per second",
			rate:     -1536 * Bps,
			expected: "-1.5 KiB/s",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actualValue := test.rate.String()

			if actualValue != test.expected {
				t.Fatalf("expected %q, got %q", test.expected, actualValue)
			}
		})
	}
}

func TestRateBinaryString(t *testing.T) {
	rate := 1536 * Bps

	actualValue := rate.BinaryString()
	expectedValue := "1.5 KiB/s"

	if actualValue != expectedValue {
		t.Fatalf("expected %q, got %q", expectedValue, actualValue)
	}
}

func TestRateDecimalString(t *testing.T) {
	tests := []struct {
		name     string
		rate     Rate
		expected string
	}{
		{
			name:     "bytes per second",
			rate:     512 * Bps,
			expected: "512 B/s",
		},
		{
			name:     "kilobytes per second",
			rate:     1500 * Bps,
			expected: "1.5 KB/s",
		},
		{
			name:     "megabytes per second",
			rate:     1500 * KBps,
			expected: "1.5 MB/s",
		},
		{
			name:     "gigabytes per second",
			rate:     1500 * MBps,
			expected: "1.5 GB/s",
		},
		{
			name:     "terabytes per second",
			rate:     1500 * GBps,
			expected: "1.5 TB/s",
		},
		{
			name:     "negative kilobytes per second",
			rate:     -1500 * Bps,
			expected: "-1.5 KB/s",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actualValue := test.rate.DecimalString()

			if actualValue != test.expected {
				t.Fatalf("expected %q, got %q", test.expected, actualValue)
			}
		})
	}
}

func TestPerSecond(t *testing.T) {
	tests := []struct {
		name     string
		size     Size
		duration time.Duration
		expected Rate
	}{
		{
			name:     "positive duration",
			size:     1500 * KB,
			duration: 10 * time.Second,
			expected: 150 * KBps,
		},
		{
			name:     "fractional rate is truncated",
			size:     10 * B,
			duration: 3 * time.Second,
			expected: 3 * Bps,
		},
		{
			name:     "zero duration",
			size:     1500 * KB,
			duration: 0,
			expected: 0,
		},
		{
			name:     "negative duration",
			size:     1500 * KB,
			duration: -1 * time.Second,
			expected: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actualValue := PerSecond(test.size, test.duration)

			if actualValue != test.expected {
				t.Fatalf("expected %d, got %d", test.expected, actualValue)
			}
		})
	}
}

func TestDurationFor(t *testing.T) {
	tests := []struct {
		name     string
		size     Size
		rate     Rate
		expected time.Duration
	}{
		{
			name:     "positive rate",
			size:     1500 * KB,
			rate:     150 * KBps,
			expected: 10 * time.Second,
		},
		{
			name:     "fractional duration",
			size:     1500 * KB,
			rate:     1 * MBps,
			expected: 1500 * time.Millisecond,
		},
		{
			name:     "zero rate",
			size:     1500 * KB,
			rate:     0,
			expected: 0,
		},
		{
			name:     "negative rate",
			size:     1500 * KB,
			rate:     -1 * KBps,
			expected: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actualValue := DurationFor(test.size, test.rate)

			if actualValue != test.expected {
				t.Fatalf("expected %s, got %s", test.expected, actualValue)
			}
		})
	}
}

func TestTransferredIn(t *testing.T) {
	tests := []struct {
		name     string
		rate     Rate
		duration time.Duration
		expected Size
	}{
		{
			name:     "positive duration",
			rate:     150 * KBps,
			duration: 10 * time.Second,
			expected: 1500 * KB,
		},
		{
			name:     "fractional size is truncated",
			rate:     1 * Bps,
			duration: 1500 * time.Millisecond,
			expected: 1 * B,
		},
		{
			name:     "zero duration",
			rate:     150 * KBps,
			duration: 0,
			expected: 0,
		},
		{
			name:     "negative duration",
			rate:     150 * KBps,
			duration: -1 * time.Second,
			expected: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actualValue := TransferredIn(test.rate, test.duration)

			if actualValue != test.expected {
				t.Fatalf("expected %d, got %d", test.expected, actualValue)
			}
		})
	}
}

func TestParseRate(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		expected Rate
	}{
		{
			name:     "bytes per second slash",
			value:    "512 B/s",
			expected: 512 * Bps,
		},
		{
			name:     "byte word per second slash",
			value:    "512 byte/s",
			expected: 512 * Bps,
		},
		{
			name:     "bytes word per second slash",
			value:    "512 bytes/s",
			expected: 512 * Bps,
		},
		{
			name:     "decimal uppercase slash",
			value:    "1.5 MB/s",
			expected: 1500 * KBps,
		},
		{
			name:     "decimal lowercase slash",
			value:    "1.5 mb/s",
			expected: 1500 * KBps,
		},
		{
			name:     "binary uppercase slash",
			value:    "1.5 MiB/s",
			expected: 1536 * KiBps,
		},
		{
			name:     "binary lowercase slash",
			value:    "1.5 mib/s",
			expected: 1536 * KiBps,
		},
		{
			name:     "decimal uppercase ps",
			value:    "1.5 MBps",
			expected: 1500 * KBps,
		},
		{
			name:     "binary uppercase ps",
			value:    "1.5 MiBps",
			expected: 1536 * KiBps,
		},
		{
			name:     "extra spaces slash",
			value:    "  1.5   MiB / s  ",
			expected: 1536 * KiBps,
		},
		{
			name:     "fractional bytes are truncated",
			value:    "1.5 B/s",
			expected: 1 * Bps,
		},
		{
			name:     "maximum int64 bytes per second",
			value:    "9223372036854775807 B/s",
			expected: Rate(9223372036854775807),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actualValue, err := ParseRate(test.value)
			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}

			if actualValue != test.expected {
				t.Fatalf("expected %d, got %d", test.expected, actualValue)
			}
		})
	}
}

func TestParseRateInvalidValues(t *testing.T) {
	tests := []struct {
		name  string
		value string
	}{
		{
			name:  "empty value",
			value: "",
		},
		{
			name:  "without per second suffix",
			value: "10 MB",
		},
		{
			name:  "invalid number",
			value: "abc MiB/s",
		},
		{
			name:  "unsupported mixed decimal bit-like unit slash",
			value: "10 Mb/s",
		},
		{
			name:  "unsupported mixed binary bit-like unit slash",
			value: "10 Mib/s",
		},
		{
			name:  "unsupported lowercase bit unit slash",
			value: "10 b/s",
		},
		{
			name:  "unsupported bit-like ps unit",
			value: "10 Mbps",
		},
		{
			name:  "unsupported lowercase bit-like ps unit",
			value: "10 mbps",
		},
		{
			name:  "unsupported mixed binary ps unit",
			value: "10 Mibps",
		},
		{
			name:  "unsupported unit",
			value: "10 XB/s",
		},
		{
			name:  "overflow bytes per second",
			value: "9223372036854775808 B/s",
		},
		{
			name:  "overflow after unit multiplication",
			value: "9223372036854775807 KiB/s",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := ParseRate(test.value)
			if err == nil {
				t.Fatalf("expected error, got nil")
			}
		})
	}
}

func TestMustParseRate(t *testing.T) {
	actualValue := MustParseRate("1.5 MiB/s")
	expectedValue := 1536 * KiBps

	if actualValue != expectedValue {
		t.Fatalf("expected %d, got %d", expectedValue, actualValue)
	}
}

func TestMustParseRateInvalidValue(t *testing.T) {
	defer func() {
		recoveredValue := recover()
		if recoveredValue == nil {
			t.Fatalf("expected panic, got nil")
		}
	}()

	_ = MustParseRate("10 Mbps")
}
