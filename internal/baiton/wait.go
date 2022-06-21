package wait

import (
	"errors"
	"fmt"
	"os"
	"time"
)

func WaitFile(filepath string, delay_s, timeout_s uint) error {
	timeout := time.Now().Add(time.Duration(timeout_s) * time.Second)
	for {
		if _, err := os.Stat(filepath); errors.Is(err, os.ErrNotExist) {
			sleep := time.Duration(delay_s) * time.Second
			time.Sleep(sleep)
		} else {
			return nil
		}
		if time.Now().After(timeout) {
			return fmt.Errorf("file %v did not appear in %v seconds", filepath, timeout_s)
		}
	}
}
