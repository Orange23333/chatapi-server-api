package main

import (
	"fmt"
	"time"
)

func test_main() {
	loc, _ := time.LoadLocation("UTC")
	time.ParseInLocation("2006-01-02 15:04:05", "2006-01-02 15:04:05", loc)

	fmt.Printf("")
}
