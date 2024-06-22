package main

import (
	"fmt"
	"math/rand"
	"time"
)

var MAX_CHICKEN_PRICE float32 = 5
var MAX_TOFU_PRICE float32 = 5

func main() {
	chickenCh := make(chan string)
	tofuCh := make(chan string)

	websites := []string{"walmart.com", "costco.com", "wholefoods.com"}
	for i := range websites {
		go checkChickenPrices(websites[i], chickenCh)
		go checkTofuPrice(websites[i], tofuCh)
	}
	sendMessage(chickenCh, tofuCh)
}

func checkChickenPrices(web string, chickenCh chan string) {
	for {
		time.Sleep(time.Second * 1)
		chickenPrice := rand.Float32() * 20
		if chickenPrice <= MAX_CHICKEN_PRICE {
			chickenCh <- web
			break
		}
	}
}

func checkTofuPrice(web string, tofuCh chan string) {
	for {
		time.Sleep(time.Second * 1)
		chickenPrice := rand.Float32() * 20
		if chickenPrice <= MAX_TOFU_PRICE {
			tofuCh <- web
			break
		}
	}
}
func sendMessage(chickenCh, tofuCh chan string) {
	select {
	case website := <-chickenCh:
		fmt.Printf("\nText sent: Found a deal on chicken at %v.", website)
	case website := <-tofuCh:
		fmt.Printf("\nEmail sent: Found a deal on chicken at %v.", website)
	}
}
