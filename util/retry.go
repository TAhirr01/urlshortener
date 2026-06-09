package util

import (
	"fmt"
	"time"
)

func Retry(n int, sleepTime time.Duration, fn func() error) error {
	for i := 0; i < n; i++ {
		err := fn()
		if err == nil {
			return nil
		}
		time.Sleep(sleepTime)
	}
	return fmt.Errorf("all %d attempts failed", n)
}
