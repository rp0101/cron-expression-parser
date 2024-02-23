package validators

import "errors"

func ValidateInput(cronArr []string) error {
	if len(cronArr) < 6 {
		return errors.New("invalid cron string, length should be equal to 6")
	}
	return nil
}
