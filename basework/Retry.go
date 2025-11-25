package basework

import (
	"time"
)

func RetryWithBool(times int, delay time.Duration, function func() bool) bool {
	for i := 0; i < times; i++ {
		if function() {
			return true
		}

		if i < times-1 {
			time.Sleep(delay)
		}

	}

	return false
}

func RetryWithError(times int, delay time.Duration, function func() error) error {
	var err error

	for i := 0; i < times; i++ {
		err = function()

		if err == nil {
			return nil
		}

		if i < times-1 {
			time.Sleep(delay)
		}

	}
	return err
}
