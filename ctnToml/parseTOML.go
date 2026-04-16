package ctnTOML

import (
	"log"

	"github.com/pelletier/go-toml/v2"
)

func (ctnTOML *CtnTOML) ParseTOML(tomlConfigFile []byte) {

	var TOMLConfig TOMLConfig

	ctnTOML.CtnDebug.Debug("Config contents: \n---\n" + string(tomlConfigFile) + "---")

	err := toml.Unmarshal(tomlConfigFile, &TOMLConfig)
	if err != nil {
		log.Fatalf("ERROR: toml.Unmarshal() == %v", err)
	}

	// Populate config when needed.

	// Standard stuff
	if TOMLConfig.Calendar.Calendar != "" {
		ctnTOML.CtnConfig.Calendar_location = TOMLConfig.Calendar.Calendar
	}

	if TOMLConfig.Calendar.Birthday != "" {
		ctnTOML.CtnConfig.Birthday_location = TOMLConfig.Calendar.Birthday
	}

	if TOMLConfig.Calendar.Format != "" {
		ctnTOML.CtnConfig.Format = TOMLConfig.Calendar.Format
	}

	if len(TOMLConfig.Calendar.Output) >= 1 {
		ctnTOML.CtnDebug.Debug("Found Output config.")
		ctnTOML.CtnConfig.Output = TOMLConfig.Calendar.Output
	}

	// Discord things
	if TOMLConfig.Discord.WebHookURL != "" {
		ctnTOML.CtnConfig.DiscordWebHookURL = TOMLConfig.Discord.WebHookURL
	}

}
