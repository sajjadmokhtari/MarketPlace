package handler

type PhoneRequest struct {
	Phone string `json:"phone"`
}

type OTPRequest struct {
	Phone string `json:"phone"`
	OTP   string `json:"otp"`
}

type Response struct {
	Valid   bool   `json:"valid"`
	Message string `json:"message,omitempty"`
}
