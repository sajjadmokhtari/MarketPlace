package kafka

import (
	"MarketPlace/logging"
	"context"
	"encoding/json"
	"fmt"

	"github.com/segmentio/kafka-go"
)

// Ad ساختار آگهی
type Ad struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

// ProduceAd پیام یک آگهی جدید به Kafka می‌فرسته
func ProduceAd(ad Ad) {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "new-ads",
	})
	defer writer.Close()

	msg, err := json.Marshal(ad) // تبدیل به JSON
	if err != nil {
		logging.GetLogger().Errorw("❌ Error marshaling ad:", err)
		fmt.Println("❌ Error marshaling ad:", err)
		return
	}

	err = writer.WriteMessages(context.Background(), // ارسال پیام
		kafka.Message{
			Key:   []byte(ad.ID),
			Value: msg,
		},
	)
	if err != nil {
		logging.GetLogger().Errorw("❌ Could not write message:", err)
		fmt.Println("❌ Could not write message:", err)
		return
	}

	logging.GetLogger().Infow("✅ Message sent successfully", "message", string(msg))
	fmt.Println("✅ Message sent successfully:", string(msg)) // چاپ مستقیم
}
