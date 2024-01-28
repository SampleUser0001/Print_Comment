package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Comment struct {
	Time    string
	Content string
}

func printComments(comments []Comment) {
	startTime := time.Now()
	for _, comment := range comments {
		commentTime, err := convertToDuration(comment.Time)
		if err != nil {
			fmt.Println("Error parsing time:", err)
			return
		}
		elapsedTime := time.Since(startTime)
		diffTime := commentTime - elapsedTime
		if diffTime > 0 {
			time.Sleep(diffTime)
		}
		fmt.Println(comment.Time, comment.Content)
	}
}

func convertToDuration(timeString string) (time.Duration, error) {
	splited := strings.Split(timeString, ":")
	var hour, minute, second int = 0, 0, 0
	var err error
	if len(splited) == 2 {
		// 分秒
		hour = 0
		minute, err = strconv.Atoi(splited[0])
		if err != nil {
			return 0, fmt.Errorf("minute is not number")
		}
		second, err = strconv.Atoi(splited[1])
		if err != nil {
			return 0, fmt.Errorf("second is not number")
		}
	} else if len(splited) == 3 {
		// 時分秒
		hour, err = strconv.Atoi(splited[0])
		if err != nil {
			return 0, fmt.Errorf("hour is not number")
		}
		minute, err = strconv.Atoi(splited[1])
		if err != nil {
			return 0, fmt.Errorf("minute is not number")
		}
		second, err = strconv.Atoi(splited[2])
		if err != nil {
			return 0, fmt.Errorf("second is not number")
		}
	}

	duration := time.Duration(hour)*time.Hour + time.Duration(minute)*time.Minute + time.Duration(second)*time.Second
	return duration, nil
}

func main() {
	// Open the comments.tsv file
	file, err := os.Open("comments.tsv")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read the file as a CSV
	reader := csv.NewReader(file)
	reader.Comma = '\t'

	// Read all the records
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Create a slice to store the comments
	comments := make([]Comment, 0)

	// Iterate over the records and create Comment objects
	for _, record := range records {
		comment := Comment{
			Time:    record[0],
			Content: record[1],
		}
		comments = append(comments, comment)
	}

	printComments(comments)

}
