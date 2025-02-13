package datetime

import (
	"time"
)

func GetTimeNow() string {
	return time.Now().String()
}