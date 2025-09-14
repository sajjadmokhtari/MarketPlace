package utils

import (
	"github.com/google/uuid"
)


func GenerateJTI() string {
	return uuid.New().String()
}
// func NowPlusMinutes(minutes int) time.Time {
//     return time.Now().Add(time.Duration(minutes) * time.Minute)
// }
