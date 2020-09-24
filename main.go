// +build go1.13

package main

import (
	"fmt"
	"os"

	"code.cloudfoundry.org/cli/cf/cmd"
	"code.cloudfoundry.org/cli/command/common"
	"code.cloudfoundry.org/cli/util/command_parser"
	"code.cloudfoundry.org/cli/util/configv3"
	"code.cloudfoundry.org/cli/util/panichandler"
	plugin_util "code.cloudfoundry.org/cli/util/plugin"
	"code.cloudfoundry.org/cli/util/ui"
)

func main() {
	var exitCode int
	defer panichandler.HandlePanic()

	config, err := configv3.GetCFConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unexpected error: %s\n", err.Error())
		os.Exit(1)
	}

	commandUI, err := ui.NewUI(config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unexpected error: %s\n", err.Error())
		os.Exit(1)
	}

	p, err := command_parser.NewCommandParser()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unexpected error: %s\n", err.Error())
		os.Exit(1)
	}

	fmt.Printf("/main.go 1 commandUI: %T Args[1:]: %#v\n", *commandUI, os.Args[1:])
	exitCode, err = p.ParseCommandFromArgs(commandUI, os.Args[1:])
	fmt.Printf("/main.go 2 returned from ParseCommandFromArgs\n")
	if err == nil {
		fmt.Printf("/main.go 3 no errors --> exiting (by design)!\n")
		os.Exit(exitCode)
	}

	fmt.Printf("\n --- handling errors and error messaging ---\n")
	if unknownCommandError, ok := err.(command_parser.UnknownCommandError); ok {
		plugin, commandIsPlugin := plugin_util.IsPluginCommand(os.Args[1:])

		switch {
		case commandIsPlugin:
			err = plugin_util.RunPlugin(plugin)
			if err != nil {
				exitCode = 1
			}

		case common.ShouldFallbackToLegacy:
			cmd.Main(os.Getenv("CF_TRACE"), os.Args)
			//NOT REACHED, legacy main will exit the process

		default:
			unknownCommandError.Suggest(plugin_util.PluginCommandNames())
			fmt.Fprintf(os.Stderr, "%s\n", unknownCommandError.Error())
			os.Exit(1)
		}
	}

	os.Exit(exitCode)
}
