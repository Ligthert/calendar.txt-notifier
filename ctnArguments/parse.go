package ctnArguments

import (
	"flag"
)

func (ctnArguments *CtnArguments) Parse() {

	var argConfigFile string
	var argDebug bool
	var argTest bool

	flag.StringVar(&argConfigFile, "config", "calendar.txt", "File or URL location of where to find (the raw) calendar.txt")
	flag.BoolVar(&argDebug, "debug", false, "Print debug messages for development purposes [true,false] (default:false)")
	flag.BoolVar(&argTest, "test", false, "Test the config.toml file [true,false] (default:false)")

	flag.Parse()

	ctnArguments.CtnConfig.File = argConfigFile
	ctnArguments.CtnConfig.Debug = argDebug
	ctnArguments.CtnConfig.Test = argTest
}
