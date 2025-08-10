package validations

import (
	"regexp"
)

// CheckIranianMobile شماره موبایل ایرانی رو چک می‌کنه
func CheckIranianMobile(phone string) bool {
	re := regexp.MustCompile(`^09(1[0-9]|2[0-2]|3[0-9]|9[0-9])[0-9]{7}$`)
	return re.MatchString(phone)
}
