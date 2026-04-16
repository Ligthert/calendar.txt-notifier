package ctnParse

import (
	"bufio"
	"fmt"
	"log"
	"strings"
	"time"
)

// It parses a calendar,
// or a birthday calendar
// as it has at the time of writing the same format (ish)
func (ctnParse *CtnParse) ParseCalendar(calendar []byte, context string) {

	// Variables
	var month int
	var day int

	month = int(time.Now().Month())
	day = time.Now().Day()

	// Generate string to compare with
	dateStr := fmt.Sprintf("%02d-%02d", month, day)
	ctnParse.CtnDebug.Debug("Current dateStr: " + dateStr)

	// Where the magic happens
	scanner := bufio.NewScanner(strings.NewReader(string(calendar)))
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) >= 5 && line[:5] == dateStr {
			// ctnParse.CtnConfig.Message = ctnParse.CtnConfig.Message + line + "\n"
			switch context {
			case "birthday":
				ctnParse.CtnConfig.BirthdayEvents = append(ctnParse.CtnConfig.BirthdayEvents, line)
			case "calendar":
				ctnParse.CtnConfig.CalendarEvents = append(ctnParse.CtnConfig.CalendarEvents, line)
			}
		}
	}

	// Throw an error, in case of an error.
	for err := scanner.Err(); err != nil; {
		log.Fatalln(err)
	}

}
