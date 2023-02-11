/*
	Calculate total download size for NYC taxi data for 2020

For each month, we have two files: green and yellow. For example:

	https://s3.amazonaws.com/nyc-tlc/trip+data/green_tripdata_2020-03.csv
	https://s3.amazonaws.com/nyc-tlc/trip+data/yellow_tripdata_2020-03.csv

Turn the below sequential algorithm to a concurrent one using goroutine per file.
*/
package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

type chanData struct {
	num int
	err error
}

var (
	urlTemplate = "https://s3.amazonaws.com/nyc-tlc/trip+data/%s_tripdata_2020-%02d.csv"
	colors      = []string{"green", "yellow"}
)

func downloadSize(url string, c chan<- chanData) {
	req, err := http.NewRequest(http.MethodHead, url, nil)
	if err != nil {
		c <- chanData{0, err}
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		c <- chanData{0, err}
		return
	}

	if resp.StatusCode != http.StatusOK {
		c <- chanData{0, fmt.Errorf(resp.Status)}
		return
	}

	num, err := strconv.Atoi(resp.Header.Get("Content-Length"))
	c <- chanData{num, err}
}

func main() {
	start := time.Now()
	size := 0
	ch := make(chan chanData)
	for month := 1; month <= 12; month++ {
		for _, color := range colors {
			url := fmt.Sprintf(urlTemplate, color, month)
			fmt.Println(url)
			go downloadSize(url, ch)
		}
	}
	for month := 1; month <= 12; month++ {
		for range colors {
			cd := <-ch
			if cd.err != nil {
				log.Fatal(cd.err)
			}
			size += cd.num
		}
	}

	duration := time.Since(start)
	fmt.Println(size, duration)
}
