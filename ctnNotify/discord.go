package ctnNotify

import (
	"bytes"
)

// A place to chuck in all the Discord related functions.
// I am using webhooks, because a constant application is out of scope for the time-being and rocketscience in itself at the moment.
func (ctnNotify *CtnNotify) sendDiscord() {

	bodyTemplate := "{\"content\": \"TEMPLATE\"}"

	// Deal with Birthdays first
	birthdayMsg := ctnNotify.transformMarkdown(bodyTemplate, "birthday")
	birthdayJSON := []byte(birthdayMsg)
	birthdayBody := bytes.NewReader(birthdayJSON)
	ctnNotify.sendWebhook(ctnNotify.CtnConfig.DiscordWebHookURL, birthdayBody)

	// Deal with Calendar entries
	calendarMsg := ctnNotify.transformMarkdown(bodyTemplate, "calendar")
	calendarJSON := []byte(calendarMsg)
	calendarBody := bytes.NewReader(calendarJSON)
	ctnNotify.sendWebhook(ctnNotify.CtnConfig.DiscordWebHookURL, calendarBody)

}
