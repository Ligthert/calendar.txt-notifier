package ctnNotify

import (
	"strings"
)

func (ctnNotify *CtnNotify) transformMarkdown(message string, context string) string {

	var header string
	var msgBody string
	var template string

	switch context {
	case "birthday":
		header = "**Birthdays**"
		for _, v := range ctnNotify.CtnConfig.BirthdayEvents {
			msgBody = msgBody + "* " + v + "\\" + "n"
		}
	case "calendar":
		header = "**Events**"
		for _, v := range ctnNotify.CtnConfig.CalendarEvents {
			msgBody = msgBody + "* " + v + "\\" + "n"
		}
	}

	msgBody = strings.ReplaceAll(msgBody, "\n", "\\"+"n")
	template = header + "\\" + "n" + msgBody
	message = strings.ReplaceAll(message, "TEMPLATE", template)
	ctnNotify.CtnDebug.Debug(message)

	return message
}
