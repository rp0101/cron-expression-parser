package controller

import (
	"cron-expression-parser/service"
	"cron-expression-parser/validators"
	"strings"
)

type CronController struct {
	cronService *service.CronService
}

func (cc *CronController) ParseCronController(cronStr string) (map[string]string, error) {
	cronArr := strings.Fields(cronStr)
	err := validators.ValidateInput(cronArr)
	if err != nil {
		return nil, err
	}
	result := cc.cronService.ParseCronService(cronArr)
	return result, nil
}

func NewCronController() *CronController {
	mController := new(CronController)
	mController.cronService = service.NewCronService()
	return mController
}
