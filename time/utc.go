package time

import (
	"fmt"
	"time"
)

func GetUTCDate() string {
	now := time.Now()
	return fmt.Sprintf("%d-%02d-%02d", now.UTC().Year(), now.UTC().Month(), now.UTC().Day())
}

func GetUTCDateTime() string {
	now := time.Now()
	return fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", now.UTC().Year(), now.UTC().Month(), now.UTC().Day(),
		now.UTC().Hour(), now.UTC().Minute(), now.UTC().Second())
}
