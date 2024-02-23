package tests

import (
	"cron-expression-parser/service"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseCronService_Fail(t *testing.T) {
	response := map[string]string{
		"Command":      "/usr/bin/find",
		"Minute":       "0 15 30 45",
		"Hour":         "0",
		"Day of Month": "1 15",
		"Month":        "1 2 3 4 5 6 7 8 9 10 11 12",
		"Day of Week":  "1 2 3 4 5 6 7 8 9",
	}
	conArr := []string{"*/15", "0", "1,15", "*", "1-9", "/usr/bin/find"}
	service := service.NewCronService()
	service_resp := service.ParseCronService(conArr)
	assert.NotEqual(t, response, service_resp)
}

func TestParseCronService_Success(t *testing.T) {
	response := map[string]string{
		"Command":      "/usr/bin/find",
		"Minute":       "0 15 30 45",
		"Hour":         "0",
		"Day of Month": "1 15",
		"Month":        "1 2 3 4 5 6 7 8 9 10 11 12",
		"Day of Week":  "",
	}
	conArr := []string{"*/15", "0", "1,15", "*", "1-9", "/usr/bin/find"}
	service := service.NewCronService()
	service_resp := service.ParseCronService(conArr)
	assert.Equal(t, response, service_resp)
}
