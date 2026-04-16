package ctnNotify

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func (ctnNotify *CtnNotify) sendWebhook(url string, body io.Reader) {

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		os.Exit(1)
	}

	req.Header = map[string][]string{
		"Accept":       {"application/json"},
		"Content-Type": {"application/json"},
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	resp.Body.Close()

}
