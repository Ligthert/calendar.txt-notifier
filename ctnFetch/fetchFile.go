package ctnFetch

import (
	"log"
	"os"
)

func (ctnFetch *CtnFetch) fetchFile(location string) []byte {

	ctnFetch.CtnDebug.Debug("Fetching file: " + location)

	data, err := os.ReadFile(location)
	if err != nil {
		log.Fatalln(err)
	}

	return data
}
