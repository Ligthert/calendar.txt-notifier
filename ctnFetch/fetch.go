package ctnFetch

import (
	"log"
	"net/url"
	"strings"
)

func (ctnFetch *CtnFetch) Fetch(location string) []byte {

	var fetchWeb bool
	var fetchFile bool
	var data []byte

	if strings.Contains(location, "://") {

		u, err := url.ParseRequestURI(location)
		if err != nil {
			log.Fatalln("Error parsing file location: " + location)
		}

		switch u.Scheme {
		case "file":
			fetchFile = true
		case "http":
			fetchWeb = true
		case "https":
			fetchWeb = true
		}

	} else {
		fetchFile = true
	}

	if fetchWeb == true {
		data = ctnFetch.fetchWeb(location)
	}

	if fetchFile == true {
		location = ctnFetch.cleanLocation(location)
		data = ctnFetch.fetchFile(location)
	}

	return data

}
