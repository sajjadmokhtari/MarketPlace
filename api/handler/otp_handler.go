package handler

import (
	"MarketPlace/services"
	"encoding/json"
	"log"
	"net/http"
)

func SendOtpHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var req PhoneRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Valid: false, Message: "درخواست نامعتبر"})
		return
	}

	log.Printf("SendOtpHandler received phone: %s", req.Phone) // لاگ اضافه شده

	err = services.SendOTP(req.Phone)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Response{Valid: false, Message: "خطا در ارسال OTP"})
		return
	}

	json.NewEncoder(w).Encode(Response{Valid: true, Message: "کد OTP ارسال شد"})
}

func VerifyOtpHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var req OTPRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Valid: false, Message: "درخواست نامعتبر"})
		return
	}

	err = services.VerifyOTP(req.Phone, req.OTP)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Valid: false, Message: err.Error()})
		return
	}

	json.NewEncoder(w).Encode(Response{Valid: true, Message: "شما با موفقیت وارد شدید"})
}
