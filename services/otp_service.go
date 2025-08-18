package services

import (
	"errors"
	"log"
	"math/rand"
	"time"

	"MarketPlace/cache"
	"MarketPlace/logging"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

// تولید OTP شش رقمی
func GenerateOTP() string {
	otp := ""
	for i := 0; i < 6; i++ {
		otp += string('0' + rune(r.Intn(10)))
	}
	return otp
}

// ارسال OTP با محدودسازی درخواست‌ها
func SendOTP(phone string) error {
	logger := logging.GetLogger()

	// بررسی محدودیت زمانی
	if !cache.CanSendOTP(phone) {
		logger.Infow("OTP request blocked: please try later", "phone", phone)
		return errors.New("لطفاً کمی صبر کنید و دوباره تلاش کنید")
	}

	// بررسی تعداد درخواست‌ها در ۱۰ دقیقه اخیر
	if cache.OTPRequestCount(phone) >= 5 {
		logger.Infow("OTP request limit reached", "phone", phone)
		return errors.New("تعداد درخواست‌های مجاز در ۱۰ دقیقه گذشته تمام شده است")
	}

	otp := GenerateOTP()
	logger.Infow("Generated OTP", "phone", phone, "otp", otp)
	log.Println("ots is: ", otp)

	// ذخیره OTP در cache
	if err := cache.SetOTP(phone, otp); err != nil {
		logger.Errorw("Failed to save OTP in cache", "phone", phone, "error", err)
		return err
	}

	// ثبت زمان و افزایش شمارش درخواست
	cache.MarkOTPSent(phone)
	cache.IncrementOTPRequest(phone)

	logger.Infow("OTP sent successfully", "phone", phone, "otp", otp)
	return nil
}

// اعتبارسنجی OTP
func VerifyOTP(phone, otp string) error {
	logger := logging.GetLogger()

	storedOtp, err := cache.GetOTP(phone)
	if err != nil {
		logger.Errorw("OTP not found or expired", "phone", phone, "error", err)
		return errors.New("کد OTP منقضی شده یا موجود نیست")
	}

	if storedOtp != otp {
		logger.Infow("OTP mismatch", "phone", phone, "provided_otp", otp, "stored_otp", storedOtp)
		return errors.New("کد OTP نادرست است")
	}

	// حذف OTP پس از تایید موفق
	cache.DeleteOTP(phone)
	logger.Infow("OTP verified successfully", "phone", phone)
	return nil
}
