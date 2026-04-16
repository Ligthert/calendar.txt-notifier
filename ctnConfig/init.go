package ctnConfig

import (
	"gitea.ligthert.net/golang/caltxtnot/ctnTypes"
)

func Init() ctnTypes.Config {

	return ctnTypes.Config{
		Format: "MM-DD",
		Output: []string{"stdio"},
	}

}
