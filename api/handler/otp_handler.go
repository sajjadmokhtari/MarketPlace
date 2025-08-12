package handler

import (
	"MarketPlace/services"
	"encoding/json"
	"log"
	"net/http"
)



// هندلر ارسال OTP
func SendOtpHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var req PhoneRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Println("Error decoding phone request:", err)
		writeJSON(w, http.StatusBadRequest, Response{Valid: false, Message: "درخواست نامعتبر"})
		return
	}

	log.Printf("SendOtpHandler received phone: %s", req.Phone)
	

	err = services.SendOTP(req.Phone)
	if err != nil {
		log.Println("Error sending OTP:", err)
		writeJSON(w, http.StatusInternalServerError, Response{Valid: false, Message: "خطا در ارسال OTP"})
		return
	}

	writeJSON(w, http.StatusOK, Response{Valid: true, Message: "کد OTP ارسال شد"})
}




func VerifyOtpHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
        return
    }

    var req OTPRequest
    err := json.NewDecoder(r.Body).Decode(&req)
    if err != nil {
        log.Println("Error decoding OTP request:", err)
        writeJSON(w, http.StatusBadRequest, Response{Valid: false, Message: "درخواست نامعتبر"})
        return
    }

    err = services.VerifyOTP(req.Phone, req.OTP)
    if err != nil {
        log.Println("OTP verification failed:", err)
        writeJSON(w, http.StatusBadRequest, Response{Valid: false, Message: err.Error()})
        return
    }

    //  ساخت جی دبلیو تی بعد از تایید او تی پی
    token, err := services.GenerateJWT(req.Phone, "user") // نقش رو فعلاً "user" می‌ذاریم
    if err != nil {
        log.Println("Error generating JWT:", err)
        writeJSON(w, http.StatusInternalServerError, Response{Valid: false, Message: "خطا در ساخت توکن"})
        return
    }

	log.Printf("Generated JWT for phone %s: %s", req.Phone, token)//دیدن توکن 

    // ✅ ذخیره توکن در کوکی
    http.SetCookie(w, &http.Cookie{
        Name:     "token",
        Value:    token,
        Path:     "/",
        HttpOnly: true,
        Secure:   false, // چون پروژه تمرینیه
        SameSite: http.SameSiteStrictMode,
    })

    // ✅ پاسخ موفقیت با پیام مناسب
    writeJSON(w, http.StatusOK, Response{Valid: true, Message: "شما با موفقیت وارد شدید، توکن ذخیره شد"})
}






















// تابع کمکی برای ارسال پاسخ JSON با وضعیت مشخص
func writeJSON(w http.ResponseWriter, status int, resp Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(resp)
}
