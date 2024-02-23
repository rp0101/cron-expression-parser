package respository

import (
	"cron-expression-parser/util"
	_ "strconv"
	"strings"
	_ "strings"
)

type CronRepository struct {
}
type ICronRepository interface {
	ParseCronRepository(string, int, int) string
}

func (cr *CronRepository) ParseCronRepository(cronStr string, min int, max int) string {

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

func NewCronRepository() *CronRepository {
	cronRepository := new(CronRepository)
	return cronRepository
}
