package byteflow

import "testing"

func TestBytes(t *testing.T) {
	size := Bytes(1536)

	if size != 1536*B {
		t.Fatalf("expected size to be %d, got %d", 1536*B, size)
	}

	if size.Bytes() != 1536 {
		t.Fatalf("expected bytes to be %d, got %d", 1536, size.Bytes())
	}
}

func TestSizeDecimalConversions(t *testing.T) {
	tests := []struct {
		name     string
		size     Size
		expected float64
		actual   func(Size) float64
	}{
		{
			name:     "kilobytes",
			size:     1500 * B,
			expected: 1.5,
			actual:   Size.Kilobytes,
		},
		{
			name:     "megabytes",
			size:     1500 * KB,
			expected: 1.5,
			actual:   Size.Megabytes,
		},
		{
			name:     "gigabytes",
			size:     1500 * MB,
			expected: 1.5,
			actual:   Size.Gigabytes,
		},
		{
			name:     "terabytes",
			size:     1500 * GB,
			expected: 1.5,
			actual:   Size.Terabytes,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actualValue := test.actual(test.size)

			if actualValue != test.expected {
				t.Fatalf("expected %v, got %v", test.expected, actualValue)
			}
		})
	}
}

func TestSizeBinaryConversions(t *testing.T) {
	tests := []struct {
		name     string
		size     Size
		expected float64
		actual   func(Size) float64
	}{
		{
			name:     "kibibytes",
			size:     1536 * B,
			expected: 1.5,
			actual:   Size.Kibibytes,
		},
		{
			name:     "mebibytes",
			size:     1536 * KiB,
			expected: 1.5,
			actual:   Size.Mebibytes,
		},
		{
			name:     "gibibytes",
			size:     1536 * MiB,
			expected: 1.5,
			actual:   Size.Gibibytes,
		},
		{
			name:     "tebibytes",
			size:     1536 * GiB,
			expected: 1.5,
			actual:   Size.Tebibytes,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actualValue := test.actual(test.size)

			if actualValue != test.expected {
				t.Fatalf("expected %v, got %v", test.expected, actualValue)
			}
		})
	}
}

func TestSizeString(t *testing.T) {
	tests := []struct {
		name     string
		size     Size
		expected string
	}{
		{
			name:     "bytes",
			size:     512 * B,
			expected: "512 B",
		},
		{
			name:     "kibibytes",
			size:     1536 * B,
			expected: "1.5 KiB",
		},
		{
			name:     "mebibytes",
			size:     1536 * KiB,
			expected: "1.5 MiB",
		},
		{
			name:     "gibibytes",
			size:     1536 * MiB,
			expected: "1.5 GiB",
		},
		{
			name:     "tebibytes",
			size:     1536 * GiB,
			expected: "1.5 TiB",
		},
		{
			name:     "two decimal places",
			size:     123456789 * B,
			expected: "117.74 MiB",
		},
		{
			name:     "negative bytes",
			size:     -512 * B,
			expected: "-512 B",
		},
		{
			name:     "negative kibibytes",
			size:     -1536 * B,
			expected: "-1.5 KiB",
		},
		{
			name:     "negative mebibytes",
			size:     -1536 * KiB,
			expected: "-1.5 MiB",
		},
		{
			name:     "minimum int64 size",
			size:     Size(-1 << 63),
			expected: "-8388608 TiB",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actualValue := test.size.String()

			if actualValue != test.expected {
				t.Fatalf("expected %q, got %q", test.expected, actualValue)
			}
		})
	}
}

func TestSizeBinaryString(t *testing.T) {
	size := 1536 * B

	actualValue := size.BinaryString()
	expectedValue := "1.5 KiB"

	if actualValue != expectedValue {
		t.Fatalf("expected %q, got %q", expectedValue, actualValue)
	}
}

func TestSizeDecimalString(t *testing.T) {
	tests := []struct {
		name     string
		size     Size
		expected string
	}{
		{
			name:     "bytes",
			size:     512 * B,
			expected: "512 B",
		},
		{
			name:     "kilobytes",
			size:     1500 * B,
			expected: "1.5 KB",
		},
		{
			name:     "megabytes",
			size:     1500 * KB,
			expected: "1.5 MB",
		},
		{
			name:     "gigabytes",
			size:     1500 * MB,
			expected: "1.5 GB",
		},
		{
			name:     "terabytes",
			size:     1500 * GB,
			expected: "1.5 TB",
		},
		{
			name:     "two decimal places",
			size:     123456789 * B,
			expected: "123.46 MB",
		},
		{
			name:     "negative bytes",
			size:     -512 * B,
			expected: "-512 B",
		},
		{
			name:     "negative kilobytes",
			size:     -1500 * B,
			expected: "-1.5 KB",
		},
		{
			name:     "negative megabytes",
			size:     -1500 * KB,
			expected: "-1.5 MB",
		},
		{
			name:     "minimum int64 size",
			size:     Size(-1 << 63),
			expected: "-9223372.04 TB",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actualValue := test.size.DecimalString()

			if actualValue != test.expected {
				t.Fatalf("expected %q, got %q", test.expected, actualValue)
			}
		})
	}
}

func TestParseSize(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		expected Size
	}{
		{
			name:     "without unit",
			value:    "512",
			expected: 512 * B,
		},
		{
			name:     "bytes",
			value:    "512 B",
			expected: 512 * B,
		},
		{
			name:     "byte word",
			value:    "512 byte",
			expected: 512 * B,
		},
		{
			name:     "bytes word",
			value:    "512 bytes",
			expected: 512 * B,
		},
		{
			name:     "decimal uppercase",
			value:    "1.5 MB",
			expected: 1500 * KB,
		},
		{
			name:     "decimal lowercase",
			value:    "1.5 mb",
			expected: 1500 * KB,
		},
		{
			name:     "binary uppercase",
			value:    "1.5 MiB",
			expected: 1536 * KiB,
		},
		{
			name:     "binary lowercase",
			value:    "1.5 mib",
			expected: 1536 * KiB,
		},
		{
			name:     "extra spaces",
			value:    "  1.5   MiB  ",
			expected: 1536 * KiB,
		},
		{
			name:     "fractional bytes are truncated",
			value:    "1.5 B",
			expected: 1 * B,
		},
		{
			name:     "maximum int64 bytes",
			value:    "9223372036854775807 B",
			expected: Size(9223372036854775807),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actualValue, err := ParseSize(test.value)
			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}

			if actualValue != test.expected {
				t.Fatalf("expected %d, got %d", test.expected, actualValue)
			}
		})
	}
}

func TestParseSizeInvalidValues(t *testing.T) {
	tests := []struct {
		name  string
		value string
	}{
		{
			name:  "empty value",
			value: "",
		},
		{
			name:  "invalid number",
			value: "abc MiB",
		},
		{
			name:  "unsupported mixed decimal bit-like unit",
			value: "10 Mb",
		},
		{
			name:  "unsupported mixed binary bit-like unit",
			value: "10 Mib",
		},
		{
			name:  "unsupported lowercase bit unit",
			value: "10 b",
		},
		{
			name:  "unsupported unit",
			value: "10 XB",
		},
		{
			name:  "overflow bytes",
			value: "9223372036854775808 B",
		},
		{
			name:  "overflow after unit multiplication",
			value: "9223372036854775807 KiB",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := ParseSize(test.value)
			if err == nil {
				t.Fatalf("expected error, got nil")
			}
		})
	}
}

func TestMustParseSize(t *testing.T) {
	actualValue := MustParseSize("1.5 MiB")
	expectedValue := 1536 * KiB

	if actualValue != expectedValue {
		t.Fatalf("expected %d, got %d", expectedValue, actualValue)
	}
}

func TestMustParseSizeInvalidValue(t *testing.T) {
	defer func() {
		recoveredValue := recover()
		if recoveredValue == nil {
			t.Fatalf("expected panic, got nil")
		}
	}()

	_ = MustParseSize("10 Mb")
}
