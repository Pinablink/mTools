package mtime

import (
	"fmt"
	"strconv"
	"time"
)

type HOUR string

const (
	NOT_HOUR HOUR = ""
	HOUR_1   HOUR = "1h"
	HOUR_2   HOUR = "2h"
	HOUR_3   HOUR = "3h"
	HOUR_4   HOUR = "4h"
	HOUR_5   HOUR = "5h"
	HOUR_6   HOUR = "6h"
	HOUR_7   HOUR = "7h"
	HOUR_8   HOUR = "8h"
	HOUR_9   HOUR = "9h"
	HOUR_10  HOUR = "10h"
	HOUR_11  HOUR = "11h"
	HOUR_12  HOUR = "12h"
	HOUR_13  HOUR = "13h"
	HOUR_14  HOUR = "14h"
	HOUR_15  HOUR = "15h"
	HOUR_16  HOUR = "16h"
	HOUR_17  HOUR = "17h"
	HOUR_18  HOUR = "18h"
	HOUR_19  HOUR = "19h"
	HOUR_20  HOUR = "20h"
	HOUR_21  HOUR = "21h"
	HOUR_22  HOUR = "22h"
	HOUR_23  HOUR = "23h"
	HOUR_24  HOUR = "24h"
)

//
type Timeparam struct {
	hourTarget     time.Time
	hourIn         time.Time
	intervalHour   HOUR
	intervalMinute int
	intervalSecond int
	intervalParam  string
	HProcessamento string
}

//
func (ref *Timeparam) DefHourIn(refHourIn time.Time) {
	ref.hourIn = refHourIn
}

//
func (ref *Timeparam) DefHourInterval(refHourInterval HOUR) {
	ref.intervalHour = refHourInterval
}

//
func (ref *Timeparam) DefMinuteInterval(refMinuteInterval int) {

	if refMinuteInterval >= 0 && refMinuteInterval <= 59 {
		ref.intervalMinute = refMinuteInterval
	}

}

//
func (ref *Timeparam) DefSecondInterval(refSecondInterval int) {
	ref.intervalSecond = refSecondInterval
}

//
func (ref *Timeparam) configIntervalParam() bool {
	ref.intervalParam = ref.facMessage()
	//fmt.Println(ref.intervalParam)
	return true
}

//
func (ref *Timeparam) facMessage() string {

	var val []int = make([]int, 0)
	val = minutesVal(val, ref.intervalMinute)
	val = secondsVal(val, ref.intervalSecond)

	return messageParam(val, ref.intervalHour)
}

//
func minutesVal(ref []int, refmin int) []int {

	if refmin > 0 && refmin <= 59 {
		ref = append(ref, refmin)
	} else {
		ref = append(ref, 0)
	}

	return ref
}

//
func secondsVal(ref []int, refsec int) []int {

	if refsec > 0 && refsec <= 59 {
		ref = append(ref, refsec)
	} else {
		ref = append(ref, 0)
	}

	return ref
}

//
func messageParam(values []int, confHMS HOUR) string {

	var conf string

	if confHMS == NOT_HOUR {
		conf = "%s%s%s%s"
		conf = fmt.Sprintf(conf, strconv.Itoa(values[0]), "m", strconv.Itoa(values[1]), "s")
	} else {
		conf = "%s%s%s%s%s"
		conf = fmt.Sprintf(conf, string(confHMS), strconv.Itoa(values[0]), "m", strconv.Itoa(values[1]), "s")
	}

	return conf
}
