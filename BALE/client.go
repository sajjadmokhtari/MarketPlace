package bale

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

const (
	baleToken  = "159398792:Fn5NASXzEpPsx4AtjkCPSwiAXIRkIZNLgbpzR9WV"
	baleChatID = "1248533088"
)

// این تابع OTP رو با پیام مناسب برای خودت ارسال می‌کنه
func SendOTPTOBALE(otp string) error {
	message := fmt.Sprintf("سلام، OTP شما هست: %s", otp)
	return sendToBaleSimple(baleToken, baleChatID, message)
}

// تابع اصلی ارسال پیام به بله
func sendToBaleSimple(token, chatID, text string) error {
	baseURL := fmt.Sprintf("https://tapi.bale.ai/bot%s/sendMessage", token)
	escapedText := url.QueryEscape(text)
	fullURL := fmt.Sprintf("%s?chat_id=%s&text=%s", baseURL, chatID, escapedText)

	resp, err := http.Get(fullURL)
	if err != nil {
		return fmt.Errorf("خطا در ارسال درخواست: %v", err)
	}
	defer resp.Body.Close()

	respBody, readErr := io.ReadAll(resp.Body)
	if readErr != nil {
		return fmt.Errorf("خطا در خواندن پاسخ: %v", readErr)
	}

	log.Printf("📡 وضعیت پاسخ: %d\n", resp.StatusCode)
	log.Println("📨 پاسخ API بله:", string(respBody))

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("خطای HTTP: %d", resp.StatusCode)
	}

	return nil
}
