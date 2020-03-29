package cmd

import (
	"fmt"
	"os"

	flags "github.com/jessevdk/go-flags"
	"github.com/mitchellh/go-homedir"
)

const configFilePath = "~/.huecfg"

// Cmd is the top level command line options.
type Cmd struct {
	Verbose []bool `short:"v" description:"Increase verbosity"`
}

var (
	cmd    = &Cmd{}
	parser = flags.NewParser(cmd, flags.HelpFlag)
)

// Run executes the program with the specified arguments and returns the code
// the process should exit with.
func Run(args []string) (exitCode int) {
	expandedConfigFilePath, errDetectHomedir := homedir.Expand(configFilePath)
	if errDetectHomedir != nil {
		fmt.Fprintf(os.Stderr, "WARNING: Cannot detect home dir: %s\n", errDetectHomedir)
	}

	if expandedConfigFilePath != "" {
		if err := flags.IniParse(expandedConfigFilePath, cmd); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}

	_, err := parser.ParseArgs(args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
		return 1
	}

	return 0
}
