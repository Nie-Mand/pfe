package utils

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"sync"
)

const (
	URL = "https://api.courier.com/send"
)


var waiter sync.WaitGroup

func Prepare(threads int) {
	waiter.Add(threads)
}

func Send(destination Destination) {
	defer waiter.Done()
	token := os.Getenv("COURIER_TOKEN")
	template := os.Getenv("TEMPLATE_ID")
	brand := os.Getenv("BRAND_ID")

	if token == "" {
		log.Fatal("COURIER_TOKEN is not set")
	}

	if template == "" {
		log.Fatal("TEMPLATE_ID is not set")
	}

	if brand == "" {
		log.Fatal("BRAND_ID is not set")
	}


	data := map[string]interface{}{
		"message": map[string]interface{}{
			"to": map[string]interface{}{
				"email": destination.Email,
			},
			"template": template,
			"brand_id": brand,
			"data": map[string]interface{}{
				"company": destination.Company,
				"position": destination.Position,
			},
		},
	}

	jsonStr, err := json.Marshal(data)

	if err != nil {
		log.Fatal(err)
	}

	body := bytes.NewBuffer(jsonStr)

	client := &http.Client{}
	req, err := http.NewRequest("POST", URL, body)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer " + token)

	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	log.Println("Sent email to " + destination.Company)
}


func Wait() {
	waiter.Wait()
}

