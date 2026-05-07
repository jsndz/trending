package scheduler

import "github.com/robfig/cron/v3"

func Schedule(scheduleFunc func()) error {
	c := cron.New()

	_, err := c.AddFunc("@every 1m", scheduleFunc)
	if err != nil {
		return err
	}
	c.Start()
	return nil
}
