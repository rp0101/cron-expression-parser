package service

import (
	"cron-expression-parser/engine"
)

type CronService struct {
	cronEngine *engine.CronEngine
}

func (cs *CronService) ParseCronService(cronArr []string) map[string]string {
	result := map[string]string{}
	result["Command"] = cronArr[5]
	result["Minute"] = cs.cronEngine.ParseCronEngine(cronArr[0], 0, 59)
	result["Hour"] = cs.cronEngine.ParseCronEngine(cronArr[1], 0, 23)
	result["Day of Month"] = cs.cronEngine.ParseCronEngine(cronArr[2], 1, 31)
	result["Month"] = cs.cronEngine.ParseCronEngine(cronArr[3], 1, 12)
	result["Day of Week"] = cs.cronEngine.ParseCronEngine(cronArr[4], 1, 7)
	return result
}

func NewCronService() *CronService {
	cronService := new(CronService)
	cronService.cronEngine = engine.NewCronEngine()
	return cronService
}
