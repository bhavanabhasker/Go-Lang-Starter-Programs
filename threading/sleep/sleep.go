package sleep

import "time"

func Sleep(x int) {
	<-time.After(time.Second * time.Duration(x))
}
