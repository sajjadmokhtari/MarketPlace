package kafka

import (
	"MarketPlace/logging"
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

// کانکت شدن به کافکا و خواندن پیام
func ConsumeAds() {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "new-ads",
		GroupID: "ads-consumer-group",
	})
	defer reader.Close()

	logging.GetLogger().Infow("👂 Waiting for messages...")
	fmt.Println("👂 Consumer started. Waiting for messages...") // چاپ مستقیم در ترمینال

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			logging.GetLogger().Errorw("❌ Could not read message:", err)
			fmt.Println("❌ Error reading message:", err)
			continue
		}

		logging.GetLogger().Infow("📩 Received", "message", string(msg.Value))
		fmt.Println("📩 Received message:", string(msg.Value)) // چاپ مستقیم
	}
}
