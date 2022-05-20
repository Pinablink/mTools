package mtime

import (
	"time"
)

//
type Equaltime struct {
	Timeparam
}

//
func NewEqualTime(refHourTarget time.Time) *Equaltime {
	return &Equaltime{Timeparam{hourTarget: refHourTarget, intervalHour: NOT_HOUR,
		intervalMinute: 0,
		intervalSecond: 0}}
}

//
func (ref *Equaltime) Calc() (bool, error) {

	var resultCalc bool
	resultProcess := ref.configIntervalParam()

	if resultProcess {
		duration, err := time.ParseDuration(ref.intervalParam)

		if err != nil {
			return false, err
		}

		timeDuration := ref.hourTarget.Add(duration)
		resultCalc = timeDuration.Equal(ref.hourIn)
	}

	return resultCalc, nil
}
