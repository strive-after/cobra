package timer

import (
	"time"
)

//对time的now方法进行封装，用于返回当前本地的时间time对象，此处封装主要是为了便于后续对time的进一步统一处理
func GetNowTime() time.Time {
	location, _ := time.LoadLocation("Asia/Shanghai")
	return  time.Now().In(location)
}

//对时间的推算
//ParseDuration 用于从字符串中解析duration 持续时间 ，有效单位ns us ms s m h 在add方法中我们可以传入druation，这样timer就可以得到持续之后的最终的时间
//因为我们不知道传入的值是什么所以没有办法直接用add来做我们需要先解析持续的时间，先用ParseDuration 来处理一下
/*
如果我们知道duration如下 那么我们可以直接用add来处理
GetNowTime().Add(time.Second * 60)
const (
	Nanosecond  Duration = 1
	Microsecond          = 1000 * Nanosecond
	Millisecond          = 1000 * Microsecond
	Second               = 1000 * Millisecond
	Minute               = 60 * Second
	Hour                 = 60 * Minute
)
 */
func GetCalculateTime(currentTimer time.Time,d string) (time.Time,error) {
	duration ,err := time.ParseDuration(d)
	if err != nil {
		return time.Time{}, err
	}
	return currentTimer.Add(duration),nil
}
