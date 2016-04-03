package main

import (
	"fmt"
	"os"

	"github.com/drillbits/pt3"
)

var cmdRec = &Command{
	Run:       runRec,
	UsageLine: "rec [-device device] channel rectime destfile",
	Short:     "record program",
	Long: `
Rec records...
	`,
}

var (
	device string
)

func init() {
	cmdRec.Flag.StringVar(&device, "device", "", "Specify devicefile to use")
}

func runRec(cmd *Command, args []string) {
	printf("pid = %d\n", os.Getpid())

	var err error

	r, err := pt3.NewRecord(args)
	if err != nil {
		fatalf("%s", err)
	}

	if !r.IsAvailable() {
		fatalf("Arguments are necessary!\nTry 'pt3 help rec' for more information.\n")
	}

	err = pt3.Tune(r, device)
	if err != nil {
		fatalf("Tune: %s", err)
	}

	// TODO b25

	// TODO splitter

	// TODO threading
}

func printf(format string, args ...interface{}) {
	os.Stderr.WriteString(fmt.Sprintf(format, args...))
}
