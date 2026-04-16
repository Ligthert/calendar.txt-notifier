package ctnArguments

import (
	"os"
)

func (ctnArguments *CtnArguments) Validate() {

	if ctnArguments.isFlaggedPassed("config") == false {
		ctnArguments.CtnDebug.Debug("ERROR: '--config' is not set.")
		os.Exit(1)
	}

}
