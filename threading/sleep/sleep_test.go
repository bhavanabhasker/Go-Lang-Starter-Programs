package sleep

import (
	"testing"
	"time"
)

func Test(t *testing.T) {
	const sec = 2
	start := time.Now()
	Sleep(sec)
	stop := time.Since(start).Seconds()
	if stop < sec || stop > sec*1.05 {
		t.Error("Incorrect sleep function")
	}

}
