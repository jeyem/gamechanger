package timing

import (
	"log"
	"time"
)

var (
	TimeZoneLocation *time.Location
)

const (
	lessThanDayFormat string = "20060102"
	moreThanDayFormat string = "200601"
	datetimeFormat           = "2006/01/02 15:04"
)

func StringFormatToTime(input string) time.Time {
	t, err := time.ParseInLocation(datetimeFormat, input,
		TimeZoneLocation)
	if err != nil {
		log.Println(err)
	}
	return t
}

func CurrentStringTime() string {
	return TimeToStringFormat(GetCurrentTime())
}

func TimestamptToString(t int64) string {
	return TimeToStringFormat(time.Unix(t/1000, 0))
}

func TimeToStringFormat(input time.Time) string {
	//	input = input.In(TimeZoneLocation)
	return input.In(TimeZoneLocation).Format(datetimeFormat)
}

func TimeToStringFormatUTC(input time.Time) string {
	return input.UTC().Format(datetimeFormat)
}

func StringFormatToTimestamp(input string) int64 {
	datetime := StringFormatToTime(input)
	return datetime.Unix()
}

func TimeToTimestamp(input time.Time) int64 {
	if input.Before(time.Now()) {
		return input.Unix()
	}
	return time.Now().Unix()
}
func TimeToTimestampForCassandra(input time.Time) int64 {
	return input.Unix() * 1000
}

func GetLocation(timeZoneCode int) *time.Location {
	location, err := time.LoadLocation(TimeZoneData[timeZoneCode][1].(string))
	if err != nil {
		location = time.Now().Location()
	}
	TimeZoneLocation = location
	return location
}

func GetCurrentTime() time.Time {
	return time.Now().In(TimeZoneLocation)
}

func FixChartTime(startDate, endDate time.Time) (sDate, eDate time.Time) {
	if (startDate.Equal(time.Time{})) && (endDate.Equal(time.Time{})) {
		endDate = time.Now()
		startDate = endDate.AddDate(0, 0, -7)
	} else if (startDate.Equal(time.Time{})) && (!endDate.Equal(time.Time{})) {
		startDate = endDate.AddDate(0, 0, -1)
	} else if (!startDate.Equal(time.Time{})) && (endDate.Equal(time.Time{})) {
		endDate = time.Now()
	}
	return startDate, endDate
}

// func ConvertFromMinuteTable(day, quarterOfDay int) int64 {
// 	createTime, err := time.Parse(
// 		lessThanDayFormat, strconv.Itoa(day))
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	minutes := (quarterOfDay * 15) * int(time.Minute)
// 	createTime = createTime.Add(time.Duration(minutes))
// 	return TimeToTimestampForCassandra(createTime)
// }

// func ConvertFromHourTable(day, hourOfDay int) int64 {
// 	createTime, err := time.Parse(
// 		lessThanDayFormat, strconv.Itoa(day))
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	hours := hourOfDay * int(time.Hour)
// 	createTime = createTime.Add(time.Duration(hours))
// 	return TimeToTimestampForCassandra(createTime)
// }

// func ConvertFromHoursTable(day, fourHoursOfDay int) int64 {
// 	createTime, err := time.Parse(
// 		lessThanDayFormat, strconv.Itoa(day))
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	hours := (fourHoursOfDay * 4) * int(time.Hour)
// 	createTime = createTime.Add(time.Duration(hours))
// 	return TimeToTimestampForCassandra(createTime)
// }

// func ConvertFromMonthTable(month, dayOfMonth int) int64 {
// 	createTime, err := time.Parse(
// 		moreThanDayFormat, strconv.Itoa(month))
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	days := dayOfMonth - 1
// 	createTime = createTime.AddDate(0, 0, days)
// 	return TimeToTimestampForCassandra(createTime)
// }

func CreateDayRange() string {
	t := time.Now()
	daysRange := t.Format("20060102") + ","
	loopRange := 186
	for i := 1; i < loopRange; i++ {
		daysRange += t.AddDate(0, 0, -i).Format("20060102")
		if i < loopRange-1 {
			daysRange += ","
		}
	}
	return daysRange
}
