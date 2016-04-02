package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

// Commands lists the available commands and help topics.
// The order here is the order in which they are printed by 'pt3 help'.
var commands = []*Command{
	cmdRec,
	cmdServer,
}

func main() {
	flag.Usage = usage
	flag.Parse()
	log.SetFlags(0)

	args := flag.Args()
	if len(args) < 1 {
		usage()
	}

	if args[0] == "help" {
		help(args[1:])
		return
	}

	for _, cmd := range commands {
		if cmd.Name() == args[0] && cmd.Runnable() {
			cmd.Flag.Usage = func() {
				cmd.Usage()
			}
			cmd.Flag.Parse(args[1:])
			args = cmd.Flag.Args()
			cmd.Run(cmd, args)
			exit()
			return
		}
	}

	fmt.Fprintf(os.Stderr, "pt3: unknown subcommand %q\nRun 'pt3 help' for usage.\n", args[0])
	setExitStatus(2)
	exit()
}
