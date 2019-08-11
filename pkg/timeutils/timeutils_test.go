package timeutils

import (
	"fmt"
	"math"
	"testing"
	"time"
)

func TestNowInSecond(t *testing.T) {
	timeNowInSeconds := time.Now().Unix()
	ts := NowInSecond()
	if math.Abs(float64(ts-int64(timeNowInSeconds))) > 1 {
		t.Errorf("timestamp now in second are not equavent. timeNowInSeconds=%d, ts=%d", timeNowInSeconds, ts)
	}
}

func TestNowMonth(t *testing.T) {
	month := NowMonth()
	fmt.Printf("month: %d", month)

}
