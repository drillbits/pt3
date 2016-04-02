package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

var cmdRec = &Command{
	Run:       runRec,
	UsageLine: "rec channel rectime destfile",
	Short:     "record program",
	Long: `
Rec records...
	`,
}

func init() {
	// cmdRec.Flag.BoolVar(&flagA, "a", false, "")
}

func runRec(cmd *Command, args []string) {
	var err error

	var channel string
	if len(args) > 0 {
		channel = args[0]
	}

	var recTime time.Duration
	var recIndefinitely bool
	if len(args) > 1 {
		if args[1] == "-" {
			recIndefinitely = true
		} else {
			i, err := strconv.ParseInt(args[1], 10, 64)
			if err != nil {
				fatalf("rectime must be a number or `-`")
			}
			recTime = time.Duration(i) * time.Second
			if recTime < 0 {
				recIndefinitely = true
			}
		}
	}

	var dst io.Writer
	if len(args) > 2 {
		if args[2] == "-" {
			dst = os.Stdout
		} else {
			dst, err = os.OpenFile(args[2], os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
			if err != nil {
				fatalf("Cannot open output file: %s", err)
			}
		}
	}

	if channel == "" || (recTime < 0 && !recIndefinitely) || dst == nil {
		fatalf("Arguments are necessary!\nTry 'pt3 help rec' for more information.\n")
	}

	printf("pid = %d\n", os.Getpid())

	printf("channel = %s\n", channel)
	printf("recTime = %s\n", recTime)
	printf("recIndefinitely = %t\n", recIndefinitely)

	// TODO check tuner

	// TODO b25

	// TODO splitter

	// TODO threading
}

func printf(format string, args ...interface{}) {
	os.Stderr.WriteString(fmt.Sprintf(format, args...))
}
