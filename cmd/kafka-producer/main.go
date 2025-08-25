package main

import "MarketPlace/kafka"

func main() {
  x := kafka.Ad{
	ID: "1",
	Title: "hi this  is test ",

  }
  kafka.ProduceAd(x)

}
