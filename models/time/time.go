package time

// import (
// 	"fmt"
// 	"strconv"
// 	"strings"
// 	"time"

// 	"github.com/feeeei/backpack-go/utils"
// )

// type Time time.Time

// func (t *Time) UnmarshalJSON(data []byte) error {
// 	str := strings.Trim(utils.BytesToString(data), "\"")

// 	ct, err := time.Parse(time.RFC3339, str)
// 	if err == nil {
// 		*t = Time(ct)
// 		return nil
// 	}

// 	ct, err = time.Parse(time.DateTime, str)
// 	if err == nil {
// 		*t = Time(ct)
// 		return nil
// 	}

// 	ct, err = time.Parse("2006-01-02T15:04:05", str)
// 	if err == nil {
// 		*t = Time(ct)
// 		return nil
// 	}

// 	timestamp, err := strconv.ParseInt(str, 10, 64)
// 	if err != nil {
// 		return fmt.Errorf("invalid time format: %s", str)
// 	}

// 	switch len(str) {
// 	case 13:
// 		*t = Time(time.Unix(0, timestamp*int64(time.Millisecond)))
// 	case 16:
// 		*t = Time(time.Unix(0, timestamp*int64(time.Microsecond)))
// 	default:
// 		return fmt.Errorf("invalid time format: %s", str)
// 	}

// 	return nil
// }
