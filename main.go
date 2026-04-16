package main

import (
	"fmt"
	"os"

	"gitea.ligthert.net/golang/caltxtnot/ctnArguments"
	"gitea.ligthert.net/golang/caltxtnot/ctnConfig"
	"gitea.ligthert.net/golang/caltxtnot/ctnDebug"
	"gitea.ligthert.net/golang/caltxtnot/ctnFetch"
	"gitea.ligthert.net/golang/caltxtnot/ctnNotify"
	"gitea.ligthert.net/golang/caltxtnot/ctnParse"
	ctnTOML "gitea.ligthert.net/golang/caltxtnot/ctnToml"
)

func main() {

	// Initialize a bunch of structs.
	config := ctnConfig.Init()
	ctnDebug := ctnDebug.CtnDebug{CtnConfig: &config}
	ctnArguments := ctnArguments.CtnArguments{CtnConfig: &config, CtnDebug: &ctnDebug}
	ctnFetch := ctnFetch.CtnFetch{CtnConfig: &config, CtnDebug: &ctnDebug}
	ctnParse := ctnParse.CtnParse{CtnConfig: &config, CtnDebug: &ctnDebug}
	ctnTOML := ctnTOML.CtnTOML{CtnConfig: &config, CtnDebug: &ctnDebug}
	ctnNotify := ctnNotify.CtnNotify{CtnConfig: &config, CtnDebug: &ctnDebug}

	// Parse the command-line arguments and validate input (a bit)
	ctnArguments.Parse()
	ctnArguments.Validate()

	// If --test is expressed, test the toml file and bugger off
	if config.Test == true {
		ctnTOML.ParseTOML(ctnFetch.Fetch(config.File))
		fmt.Println("Okay")
		os.Exit(0)
	}

	// Print some debug things
	ctnDebug.Debug("Calendar.txt Notifyer fully initialized")
	ctnDebug.Debug("Config file location: " + config.File)

	// Unmarshaling TOML config
	// In case of error: 💩➡️🛏️
	ctnTOML.ParseTOML(ctnFetch.Fetch(config.File))

	// Parse the Birthday file
	if config.Birthday_location != "" {
		ctnParse.ParseCalendar(ctnFetch.Fetch(config.Birthday_location), "birthday")
	}

	// Parse calender.txt file
	if config.Calendar_location != "" {
		ctnParse.ParseCalendar(ctnFetch.Fetch(config.Calendar_location), "calendar")
	}

	// Send stuff to receiver
	ctnNotify.Send()

}
