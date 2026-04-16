package ctnArguments

import "flag"

func (ctnArguments *CtnArguments) isFlaggedPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}
