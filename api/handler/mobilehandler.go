package handler

import (
	"MarketPlace/validations"
	"encoding/json"
	"net/http"
)

func CheckPhoneHandler(w http.ResponseWriter, r *http.Request) {
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

	isValid := validations.CheckIranianMobile(req.Phone)
	if !isValid {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Valid: false, Message: "شماره موبایل معتبر نیست"})
		return
	}

	json.NewEncoder(w).Encode(Response{Valid: true})
}
