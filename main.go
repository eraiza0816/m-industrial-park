package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

func main() {
	args := os.Args[1:]
	var status string
	switch len(args) {
	case 0:
		status = "出勤"
	case 1:
		status = "退勤"
	default:
		status = "出勤"
	}

	now := time.Now()
	date := now.Format("2006-01-02")
	time := now.Format("15:04:05")

	record := []string{date, time, status}
	file, err := os.OpenFile("勤怠記録.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.Write(record)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}
