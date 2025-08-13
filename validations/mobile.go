package validations

import (
	"regexp"
	"strings"
)

// CheckIranianMobile شماره موبایل ایرانی رو چک می‌کنه
func CheckIranianMobile(phone string) bool {
	// حذف فاصله‌ها یا کاراکترهای اضافی احتمالی
	phone = strings.TrimSpace(phone)

	// بررسی اینکه فقط شامل عدد باشه
	if !regexp.MustCompile(`^\d+$`).MatchString(phone) {
		return false
	}

	// بررسی طول دقیق
	if len(phone) != 11 {
		return false
	}

	// بررسی الگوی شماره موبایل ایرانی
	re := regexp.MustCompile(`^09(1[0-9]|2[0-2]|3[0-9]|9[0-9])[0-9]{7}$`)
	return re.MatchString(phone)
}
