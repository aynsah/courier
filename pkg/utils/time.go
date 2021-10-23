package utils

import (
	"strconv"
	"time"
)

func GetTime() (year string, month string, day string, hour string, min string, sec string) {
	iyear, tmonth, iday := time.Now().Date()
	ihour, imin, isec := time.Now().Clock()
	imonth := int(tmonth)

	year, month, day = strconv.Itoa(iyear), strconv.Itoa(imonth), strconv.Itoa(iday)
	hour, min, sec = strconv.Itoa(ihour), strconv.Itoa(imin), strconv.Itoa(isec)

	if len(day) == 1 {
		day = "0" + day
	}

	if len(month) == 1 {
		month = "0" + month
	}

	return
}
