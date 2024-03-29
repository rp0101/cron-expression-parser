package engine

import (
	"cron-expression-parser/util"
	_ "strconv"
	"strings"
	_ "strings"
)

type CronEngine struct {
}
type ICronEngine interface {
	ParseCronEngine(string, int, int) string
}

func (cr *CronEngine) ParseCronEngine(cronStr string, min int, max int) string {

	if strings.Contains(cronStr, "*/") {
		value := util.StepValues(cronStr, min, max)
		return value
	}
	if strings.Contains(cronStr, ",") {
		value := util.CommaSepValues(cronStr, min, max)
		return value
	}
	if cronStr == "*" {
		value := util.AllValues(cronStr, min, max)
		return value
	}
	if strings.Contains(cronStr, "-") {
		value := util.RangeValues(cronStr, min, max)
		return value
	}
	return cronStr
}

func NewCronEngine() *CronEngine {
	cronEngine := new(CronEngine)
	return cronEngine
}
