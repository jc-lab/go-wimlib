package main

import "C"
import (
	"flag"
	"fmt"
	"github.com/jc-lab/go-wimlib/binding"
	"github.com/jc-lab/go-wimlib/command/apply"
	"github.com/jc-lab/go-wimlib/command/common"
	"github.com/jc-lab/go-wimlib/command/info"
	"github.com/pkg/errors"
	"log"
	"os"
)

type CommandHandler func(appFlag *common.AppFlag, args []string) (interface{}, error)

var commands = map[string]CommandHandler{
	"apply": apply.Main,
	"info":  info.Main,
}

func main() {
	var appFlag common.AppFlag

	appName := os.Args[0]
	appFlag.InitFlags(flag.CommandLine)
	flag.Parse()

	args := flag.Args()

	if len(args) < 1 {
		fmt.Printf("usage: %s <command>\n", appName)
		fmt.Printf("\t%s apply arguments...\n", appName)
		os.Exit(2)
	}

	err := binding.GlobalInit(0)
	if err != nil {
		err = errors.Wrap(err, "wimlib global init failed")
		if appFlag.Json {
			common.WriteFinishError(err)
		}
		log.Fatalln(err)
	}
	defer binding.GlobalCleanup()

	command := args[0]
	commandHandler, ok := commands[command]
	if !ok {
		err = errors.Errorf("unknown command: %s", command)
		if appFlag.Json {
			common.WriteFinishError(err)
		}
		log.Fatalln(err)
	}

	data, err := commandHandler(&appFlag, args[1:])
	if err != nil {
		if appFlag.Json {
			common.WriteFinishError(err)
		}
		log.Fatalln(err)
	} else {
		if appFlag.Json {
			common.WriteFinishSuccess(data)
		}
	}
}
