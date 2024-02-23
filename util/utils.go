package util

import (
	"strconv"
	"strings"
)

func StepValues(stepValueStr string, min, max int) string {
	var values []string
	step, err := strconv.Atoi(stepValueStr[2:])
	if err != nil {
		return ""
	}
	for i := min; i <= max; i += step {
		values = append(values, strconv.Itoa(i))
	}
	return strings.Join(values, " ")
}

func AllValues(allValueStr string, min, max int) string {
	var values []string
	for i := min; i <= max; i++ {
		values = append(values, strconv.Itoa(i))
	}
	return strings.Join(values, " ")
}
func CommaSepValues(commaSepValueStr string, min, max int) string {
	var values []string
	parts := strings.Split(commaSepValueStr, ",")
	for _, part := range parts {
		value, err := strconv.Atoi(part)
		if err != nil || value < min || value > max {
			return ""
		}
		values = append(values, strconv.Itoa(value))
	}
	return strings.Join(values, " ")
}

func RangeValues(rangeValueStr string, min, max int) string {
	var values []string
	range_parts := strings.Split(rangeValueStr, "-")
	if len(range_parts) != 2 {
		return ""
	}
	first, err := strconv.Atoi(range_parts[0])
	if err != nil || first < min || first > max {
		return ""
	}
	last, err := strconv.Atoi(range_parts[1])
	if err != nil || last < min || last > max || last < first {
		return ""
	}
	for i := first; i <= last; i++ {
		values = append(values, strconv.Itoa(i))
	}
	return strings.Join(values, " ")
}
