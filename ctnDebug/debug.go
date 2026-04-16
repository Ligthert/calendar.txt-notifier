package ctnDebug

import "log"

func (ctnDebug *CtnDebug) Debug(debugline string) {

	if ctnDebug.CtnConfig.Debug == true {
		log.Println(debugline)
	}

}
