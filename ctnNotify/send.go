package ctnNotify

import "fmt"

func (ctnNotify *CtnNotify) Send() {

	// itterate through config.Notify
	for _, val := range ctnNotify.CtnConfig.Output {

		if val == "stdio" {
			ctnNotify.CtnDebug.Debug("Printing to stdio")
			if len(ctnNotify.CtnConfig.BirthdayEvents) >= 1 {
				fmt.Println("Birthdays:")
				for _, v := range ctnNotify.CtnConfig.BirthdayEvents {
					fmt.Println("- " + v)
				}
			}
			if len(ctnNotify.CtnConfig.CalendarEvents) >= 1 {
				fmt.Println("Events:")
				for _, v := range ctnNotify.CtnConfig.CalendarEvents {
					fmt.Println("- " + v)
				}
			}

		}

		if val == "discord" {
			ctnNotify.sendDiscord()
		}

	}

}
