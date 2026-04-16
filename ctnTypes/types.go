package ctnTypes

type Config struct {
	// Paramters
	File  string
	Debug bool
	Test  bool

	// Config file things and its derivatives
	Calendar_location string
	Calendar          []byte
	CalendarEvents    []string
	Format            string
	Birthday_location string
	Birthday          []byte
	BirthdayEvents    []string
	Output            []string

	// Discord values
	DiscordWebHookURL string
}
