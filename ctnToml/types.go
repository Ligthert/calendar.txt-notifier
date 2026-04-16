package ctnTOML

import (
	"gitea.ligthert.net/golang/caltxtnot/ctnDebug"
	"gitea.ligthert.net/golang/caltxtnot/ctnTypes"
)

type CtnTOML struct {
	CtnConfig *ctnTypes.Config
	CtnDebug  *ctnDebug.CtnDebug
}

type TOMLConfig struct {
	Calendar struct {
		Calendar string
		Format   string
		Birthday string
		Output   []string
	}
	Discord struct {
		WebHookURL string
	}
}
