package byteflow

import (
	"fmt"
	"math/big"
	"regexp"
	"strings"
)

const maxInt64 = int64(^uint64(0) >> 1)

var (
	maxInt64Value = big.NewRat(maxInt64, 1)
	sizePattern   = regexp.MustCompile(`^\s*([0-9]+(?:\.[0-9]+)?)\s*([A-Za-z]*)\s*$`)
)

// ParseSize parses a text value into a Size.
//
// Supported units are B, byte, bytes, KB, MB, GB, TB, KiB, MiB, GiB and TiB.
// Decimal units use base 1000, while binary units use base 1024.
// Lowercase decimal and binary unit names are accepted, for example kb, mb, kib and mib.
//
// If no unit is provided, the value is treated as bytes.
func ParseSize(value string) (Size, error) {
	numberValue, unitName, err := parseSizeParts(value)
	if err != nil {
		return 0, err
	}

	unitMultiplier, err := sizeUnitMultiplier(unitName)
	if err != nil {
		return 0, err
	}

	sizeValue := new(big.Rat).Mul(numberValue, big.NewRat(int64(unitMultiplier), 1))
	if sizeValue.Cmp(maxInt64Value) > 0 {
		return 0, fmt.Errorf("size value %q overflows Size", value)
	}

	truncatedValue := new(big.Int).Quo(sizeValue.Num(), sizeValue.Denom())

	return Size(truncatedValue.Int64()), nil
}

// MustParseSize parses a text value into a Size.
//
// It panics if the value cannot be parsed.
func MustParseSize(value string) Size {
	size, err := ParseSize(value)
	if err != nil {
		panic(err)
	}

	return size
}

func parseSizeParts(value string) (*big.Rat, string, error) {
	matches := sizePattern.FindStringSubmatch(value)
	if matches == nil {
		return nil, "", fmt.Errorf("invalid size value %q", value)
	}

	numberValue, ok := new(big.Rat).SetString(matches[1])
	if !ok {
		return nil, "", fmt.Errorf("invalid size number %q", matches[1])
	}

	unitName := strings.TrimSpace(matches[2])

	return numberValue, unitName, nil
}

func sizeUnitMultiplier(unitName string) (Size, error) {
	switch unitName {
	case "", "B", "byte", "bytes":
		return B, nil

	case "KB", "kb":
		return KB, nil
	case "MB", "mb":
		return MB, nil
	case "GB", "gb":
		return GB, nil
	case "TB", "tb":
		return TB, nil

	case "KiB", "kib":
		return KiB, nil
	case "MiB", "mib":
		return MiB, nil
	case "GiB", "gib":
		return GiB, nil
	case "TiB", "tib":
		return TiB, nil

	default:
		return 0, fmt.Errorf("unsupported size unit %q", unitName)
	}
}
