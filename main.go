package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gocolly/colly"
	gomail "gopkg.in/gomail.v2"
)

func main() {
	pollInterval := 15

	timerCh := time.Tick(time.Duration(pollInterval) * time.Second)

	for range timerCh {
		bikeInStock := checkStockOfBike()
		fmt.Println(bikeInStock)
		if bikeInStock {
			sendEmail()
		}
	}
}

func checkStockOfBike() bool {
	c := colly.NewCollector()
	bikeInStock := false
	c.OnHTML("a[data-size]", func(e *colly.HTMLElement) {
		if e.Attr("data-size") == "XS" {
			classes := strings.Split(e.Attr("class"), " ")
			for _, v := range classes {
				if v == "addToCart" {
					bikeInStock = true
				}
			}
		}
	})

	err := c.Visit("https://www.canyon.com/en-us/road/endurace/endurace-al-disc-7-0")
	if err != nil {
		fmt.Errorf("Error %s", err)
	}
	return bikeInStock
}

func sendEmail() {
	email := os.Getenv("SCRAPER_EMAIL_ADDRESS")
	password := os.Getenv("SCRAPER_PASSWORD")
	destination := os.Getenv("DESTINATION_ADDRESS")
	if email == "" || password == "" || destination == "" {
		log.Fatal("please set env variables")
	}
	m := gomail.NewMessage()
	m.SetHeader("From", email)
	m.SetHeader("To", destination)
	m.SetHeader("Subject", "BIKE IN STOCK")
	m.SetBody("text/html", "<b> GET THAT BIKE </b>")
	d := gomail.NewDialer("smtp.gmail.com", 587, email, password)

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}