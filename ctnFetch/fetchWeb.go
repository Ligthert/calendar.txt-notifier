package ctnFetch

import (
	"io"
	"log"
	"net/http"
)

func (ctnFetch *CtnFetch) fetchWeb(location string) []byte {

	req, err := http.NewRequest("GET", location, nil)
	if err != nil {
		log.Fatalln(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return data

}
