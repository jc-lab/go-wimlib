package common

import "flag"

type AppFlag struct {
	AppName string
	Command string
	Json    bool
}

func (a *AppFlag) InitFlags(flagSet *flag.FlagSet) {
	flagSet.BoolVar(&a.Json, "json", a.Json, "json output")
}
