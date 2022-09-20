package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

// КЛИЕНТ ОТПРАВКИ POST и GET запросов

const (
	urlCrateEvent  = "http://127.0.0.1:8080/create_event"
	urlEventForDay = "http://127.0.0.1:8080/events_for_day?user_id=1&date=2022-09-15"
	reqBody2       = `{"Id": "1", "Date": "2022-09-15", "Description": "HPBD, Dmitriy" }`
)

func main() {
	httpClient := &http.Client{
		Timeout: 10 * time.Second,
	}

	//req, err := http.NewRequest(http.MethodGet, urlEventForDay, strings.NewReader(reqBody2))
	//if err != nil {
	//	log.Fatal(err)
	//}

	req, err := http.NewRequest(http.MethodPost, urlCrateEvent, strings.NewReader(reqBody2))
	if err != nil {
		log.Fatal(err)
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	fmt.Println(resp.StatusCode)
	body, err := io.ReadAll(resp.Body)
	fmt.Println("StatusCode:", resp.StatusCode)
	fmt.Println("HEADER:", resp.Header)
	fmt.Printf("BODY :\n %s\n", string(body))

}
