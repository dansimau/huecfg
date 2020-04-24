package cmd

import (
	"fmt"
	"os"

	"github.com/dansimau/huecfg/pkg/hue"
	flags "github.com/jessevdk/go-flags"
	"github.com/mitchellh/go-homedir"
)

const configFilePath = "~/.huecfg"

// Cmd is the top level command line options.
// TODO: fix export (make unexported)
type Cmd struct {
	Host     string `short:"a" long:"host" description:"host address for Hue Bridge"`
	Username string `short:"u" long:"username" description:"username from Hue Bridge registration"`

	Verbose []bool `short:"v" description:"Increase verbosity"`
	Debug   bool   `short:"d" description:"Enable debug output"`
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

func (c *Cmd) getHueAPI() *hue.API {
	h := &hue.API{
		Host:     c.Host,
		Username: c.Username,
	}

	if len(cmd.Verbose) > 0 {
		h.Debug = true
	}

	return h
}

func (c *Cmd) getHue() *hue.Hue {
	h := hue.NewConn(c.Host, c.Username)

	if len(cmd.Verbose) > 0 {
		h.API.Debug = true
	}

	return h
}
