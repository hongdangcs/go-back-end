package bootstrap

import (
	"log"
	"time"
)

func GetTimeNow() string {

	currentTime := time.Now()
	loc := time.FixedZone("GMT+7", 7*60*60)
	currentTimeInGMT7 := currentTime.In(loc)
	timeString := currentTimeInGMT7.Format("15:04 02/01/2006")
	log.Print("now: ", timeString, " create a post")
	return timeString

}
