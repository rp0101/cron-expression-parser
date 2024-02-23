package util

import (
	"log"
	"strconv"
	"strings"
)

func StepValues(stepValueStr string, min, max int) string {
	var values []string
	step, err := strconv.Atoi(stepValueStr[2:])
	if err != nil {
		log.Println("error converting string to int, returning an empty string. Error: ", err.Error())
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
			log.Println("error converting string to int or entered value is not in specified range, " +
				"returning an empty string")
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
		log.Println("A range must have exactly 2 values, returning an empty string.")
		return ""
	}
	first, err := strconv.Atoi(range_parts[0])
	if err != nil || first < min || first > max {
		log.Println("error converting string to int or entered value is not in specified range, " +
			"returning an empty string")
		return ""
	}
	last, err := strconv.Atoi(range_parts[1])
	if err != nil || first < min || last > max || last < first {
		log.Println("error converting string to int or entered value is not in specified range, " +
			"returning an empty string")
		return ""
	}
	for i := first; i <= last; i++ {
		values = append(values, strconv.Itoa(i))
	}
	return strings.Join(values, " ")
}
