package ctnFetch

import (
	"path/filepath"
	"strings"
)

func (ctnFetch *CtnFetch) cleanLocation(location string) string {

	// Strip the protocol identifier, the hackey lazy way
	if strings.Contains(location, "file://") {
		location = location[7:]
	}

	return filepath.Clean(location)

}
