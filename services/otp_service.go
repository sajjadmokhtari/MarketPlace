package services

import (
    "errors"
    "log"
    "math/rand"
    "time"

    "MarketPlace/cache"
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
    // بررسی محدودیت زمانی
    if !cache.CanSendOTP(phone) {
        return errors.New("لطفاً کمی صبر کنید و دوباره تلاش کنید")
    }

    // بررسی تعداد درخواست‌ها در ۱۰ دقیقه اخیر
    if cache.OTPRequestCount(phone) >= 5 {
        return errors.New("تعداد درخواست‌های مجاز در ۱۰ دقیقه گذشته تمام شده است")
    }

    otp := GenerateOTP()
    log.Printf("SendOTP called for phone: %s with OTP: %s\n", phone, otp)

    err := cache.SetOTP(phone, otp)
    if err != nil {
        return err
    }

    // ثبت زمان و افزایش شمارش درخواست
    cache.MarkOTPSent(phone)
    cache.IncrementOTPRequest(phone)

    // اینجا می‌تونی کد ارسال SMS بزنی
    log.Printf("OTP برای شماره %s ارسال شد: %s\n", phone, otp)
    return nil
}

// اعتبارسنجی OTP
func VerifyOTP(phone, otp string) error {
    storedOtp, err := cache.GetOTP(phone)
    if err != nil {
        return errors.New("کد OTP منقضی شده یا موجود نیست")
    }

    if storedOtp != otp {
        return errors.New("کد OTP نادرست است")
    }

    cache.DeleteOTP(phone)
    return nil
}
