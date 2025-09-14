package services

import (
	"MarketPlace/cache"
	"MarketPlace/logging"
	"crypto/rand"
	"errors"
	"math/big"
)

func GenerateOTP() (string, error) {
	const digits = "0123456789"
	const length = 6
	otp := make([]byte, length)

	for i := 0; i < length; i++ {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(digits))))
		if err != nil {
			return "", err
		}
		otp[i] = digits[n.Int64()]
	}

	return string(otp), nil
}

func SendOTP(phone string) error {
	logger := logging.GetLogger()

	// بررسی بلاک شدن شماره
	if cache.IsPhoneBlocked(phone) {
		logger.Warnw("Phone temporarily blocked due to failed attempts", "phone", phone)
		return errors.New("شماره شما به دلیل تلاش‌های ناموفق موقتاً بلاک شده است")
	}

	// بررسی فاصله زمانی از آخرین ارسال
	if !cache.CanSendOTP(phone) {
		logger.Infow("OTP request blocked: please wait", "phone", phone)
		return errors.New("لطفاً حداقل ۳۰ ثانیه بین درخواست‌ها فاصله بگذارید")
	}

	// بررسی تعداد ارسال در یک ساعت
	if cache.OTPRequestCount(phone) >= 5 {
		logger.Infow("OTP request limit reached", "phone", phone)
		return errors.New("حداکثر تعداد ارسال OTP در یک ساعت انجام شده است")
	}

	// تولید OTP امن
	otp, err := GenerateOTP()
	if err != nil {
		logger.Errorw("Failed to generate OTP", "phone", phone, "error", err)
		return errors.New("خطا در تولید کد")
	}

	// ذخیره OTP در کش با TTL دو دقیقه
	if err := cache.SetOTP(phone, otp); err != nil {
		logger.Errorw("Failed to save OTP in cache", "phone", phone, "error", err)
		return errors.New("خطا در ذخیره‌سازی OTP")
	}

	// ثبت زمان ارسال و افزایش شمارش
	cache.MarkOTPSent(phone)         // ذخیره زمان آخرین ارسال
	cache.IncrementOTPRequest(phone) // افزایش شمارش در بازه یک‌ساعته

	// لاگ کردن OTP (فعلاً چون پیامک نداریم)
	logger.Infow("OTP generated and logged", "phone", phone, "otp", otp)

	return nil
}

func VerifyOTP(phone, otp string) error {
	logger := logging.GetLogger()

	// بررسی بلاک بودن شماره
	if cache.IsPhoneBlocked(phone) {
		logger.Warnw("Phone is blocked due to failed attempts", "phone", phone)
		return errors.New("شماره شما به دلیل تلاش‌های ناموفق موقتاً بلاک شده است")
	}

	// دریافت OTP ذخیره‌شده
	storedOtp, err := cache.GetOTP(phone)
	if err != nil {
		logger.Errorw("OTP not found or expired", "phone", phone, "error", err)
		return errors.New("کد OTP منقضی شده یا موجود نیست")
	}

	// مقایسه OTP
	if storedOtp != otp {
		logger.Infow("OTP mismatch", "phone", phone, "provided_otp", otp, "stored_otp", storedOtp)

		// افزایش شمارش تلاش‌های ناموفق
		cache.IncrementFailedAttempts(phone)

		// بررسی تعداد تلاش‌ها
		if cache.GetFailedAttempts(phone) >= 3 {
			cache.BlockPhone(phone)
			logger.Warnw("Phone blocked due to repeated OTP failures", "phone", phone)
			return errors.New("شماره شما به دلیل ۳ بار ورود اشتباه موقتاً بلاک شده است")
		}

		return errors.New("کد OTP نادرست است")
	}

	// موفقیت: حذف OTP و ریست تلاش‌ها
	cache.DeleteOTP(phone)
	cache.ResetFailedAttempts(phone)

	logger.Infow("OTP verified successfully", "phone", phone)
	return nil
}
