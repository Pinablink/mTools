package mtime

import (
	"time"
)

// Relaciona o hourTarget com o hourIn
type Previustime struct {
	Timeparam
}

//
func NewPreviustime(refHourTarget time.Time) *Previustime {
	return &Previustime{Timeparam{hourTarget: refHourTarget, intervalHour: NOT_HOUR,
		intervalMinute: 0,
		intervalSecond: 0}}
}

//
func (ref *Previustime) Calc() (bool, error) {

	var resultCalc bool
	resultProcess := ref.configIntervalParam()

	if resultProcess {
		duration, err := time.ParseDuration(ref.intervalParam)

		if err != nil {
			return false, err
		}

		timeDuration := ref.hourTarget.Add(duration)
		resultCalc = timeDuration.Before(ref.hourIn)
	}

	return resultCalc, nil
}
