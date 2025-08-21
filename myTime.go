package main

import (
	"fmt"
	"time"
)

func myTime() {

	fmt.Println("welcome to time study of Go")

	// to get current time
	presentTime := time.Now()

	// if we use time.now() the output will be: 2025-08-17 16:03:08.625141839 +0530 IST m=+0.000122534
	fmt.Println(presentTime)

	// to get desired format output you have to give default example output: DD-MM-YYYY HH:MM:MS Day -> "01-02-2006 15:04:05 Monday" have to use this to get acurate output
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	// if want to create date from past or upcoming: createDate := time.Date(year: 2026, month: time.November, date: 21, hour: 07, minutes: 25, second: 00, nSecond: 00, zone: time.UTC)
	createDate := time.Date(2026, time.November, 21, 07, 25, 00, 00, time.UTC)
	fmt.Println(createDate)
	// pretty formating
	fmt.Println(createDate.Format("01-02-2006 15:04:05 Monday"))
}
