package timing

import (
	"testing"
	"time"
)

type timeDataType struct {
	stringFormat string
	timeFormat   time.Time
	timestamp    int64
}

var (
	location  *time.Location = GetLocation(445)
	timeCases []timeDataType = []timeDataType{{"2016/09/23 09:30",
		time.Date(2016, time.September, 23, 9, 30, 0, 0, location),
		1474610400,
	}, {"2016/08/02 11:40",
		time.Date(2016, time.August, 02, 11, 40, 0, 0, location),
		1470121800,
	}, {"2014/03/20 18:55",
		time.Date(2014, time.March, 20, 18, 55, 0, 0, location),
		1395329100,
	}, {"2016/11/10 12:20",
		time.Date(2016, time.November, 10, 12, 20, 0, 0, location),
		1478767800,
	}, {"2015/04/18 20:20",
		time.Date(2015, time.April, 18, 20, 20, 0, 0, location),
		1429372200,
	},
	}
)

func TestStringFormatToTime(t *testing.T) {
	for _, timeCase := range timeCases {
		converted := StringFormatToTime(timeCase.stringFormat)
		if !converted.Equal(
			timeCase.timeFormat) {

			t.Error("parsing (time to string)",
				timeCase.stringFormat,
				"---data time Object :", timeCase.timeFormat,
				"---converted to :", converted)
		}

	}
}

func TestTimeToStringFormat(t *testing.T) {
	for _, timeCase := range timeCases {
		converted := TimeToStringFormat(timeCase.timeFormat)
		if converted != timeCase.stringFormat {
			t.Error("parsing (time to string)",
				timeCase.stringFormat,
				"---time Object :", timeCase.timeFormat,
				"---converted to :", converted)
		}
	}
}

func TestStringFormatToTimestamp(t *testing.T) {
	for _, timeCase := range timeCases {
		converted := StringFormatToTimestamp(timeCase.stringFormat)
		if converted != timeCase.timestamp {
			t.Error("parsing (string to unix timestamp)",
				timeCase.stringFormat,
				"---correct timestamp :", timeCase.timestamp,
				"---converted to :", converted)
		}
	}

}
