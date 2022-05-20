package mtime

import "time"

//
type Aftertime struct {
	Timeparam
}

//
func NewAfterTime(refHourTarget time.Time) *Aftertime {
	return &Aftertime{Timeparam{hourTarget: refHourTarget, intervalHour: NOT_HOUR,
		intervalMinute: 0,
		intervalSecond: 0}}
}

//
func (ref *Aftertime) Calc() (bool, error) {

	var resultCalc bool
	resultProcess := ref.configIntervalParam()

	if resultProcess {
		duration, err := time.ParseDuration(ref.intervalParam)

		if err != nil {
			return false, err
		}

		timeDuration := ref.hourTarget.Add(duration)
		resultCalc = timeDuration.After(ref.hourIn)
	}

	return resultCalc, nil
}
