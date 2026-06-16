package byteflow

import (
	"fmt"
	"math/big"
	"regexp"
	"strings"
)

var ratePattern = regexp.MustCompile(`^\s*([0-9]+(?:\.[0-9]+)?)\s*([A-Za-z]+)\s*(?:(/\s*s)|(ps))\s*$`)

// ParseRate parses a text value into a Rate.
//
// Supported slash-based units are B/s, byte/s, bytes/s, KB/s, MB/s, GB/s, TB/s,
// KiB/s, MiB/s, GiB/s and TiB/s.
// Supported ps-based units are Bps, KBps, MBps, GBps, TBps,
// KiBps, MiBps, GiBps and TiBps.
// Decimal units use base 1000, while binary units use base 1024.
// Lowercase slash-based decimal and binary unit names are accepted,
// for example kb/s, mb/s, kib/s and mib/s.
func ParseRate(value string) (Rate, error) {
	numberValue, unitName, usesCompactSuffix, err := parseRateParts(value)
	if err != nil {
		return 0, err
	}

	if usesCompactSuffix && unitName != "B" && strings.ToLower(unitName) == unitName {
		return 0, fmt.Errorf("unsupported compact lowercase rate unit %q", unitName+"ps")
	}

	unitMultiplier, err := rateUnitMultiplier(unitName)
	if err != nil {
		return 0, err
	}

	rateValue := new(big.Rat).Mul(numberValue, big.NewRat(int64(unitMultiplier), 1))
	if rateValue.Cmp(maxInt64Value) > 0 {
		return 0, fmt.Errorf("rate value %q overflows Rate", value)
	}

	truncatedValue := new(big.Int).Quo(rateValue.Num(), rateValue.Denom())

	return Rate(truncatedValue.Int64()), nil
}

// MustParseRate parses a text value into a Rate.
//
// It panics if the value cannot be parsed.
func MustParseRate(value string) Rate {
	rate, err := ParseRate(value)
	if err != nil {
		panic(err)
	}

	return rate
}

func parseRateParts(value string) (*big.Rat, string, bool, error) {
	matches := ratePattern.FindStringSubmatch(value)
	if matches == nil {
		return nil, "", false, fmt.Errorf("invalid rate value %q", value)
	}

	numberValue, ok := new(big.Rat).SetString(matches[1])
	if !ok {
		return nil, "", false, fmt.Errorf("invalid rate number %q", matches[1])
	}

	unitName := strings.TrimSpace(matches[2])
	usesCompactSuffix := matches[4] == "ps"

	return numberValue, unitName, usesCompactSuffix, nil
}

func rateUnitMultiplier(unitName string) (Rate, error) {
	sizeMultiplier, err := sizeUnitMultiplier(unitName)
	if err != nil {
		return 0, fmt.Errorf("unsupported rate unit %q", unitName)
	}

	return Rate(sizeMultiplier), nil
}
