package main

import (
	"testing"
)

func TestCommandName(t *testing.T) {
	cmd := &Command{
		UsageLine: "test [-foo] [-bar]",
	}
	if cmd.Name() != "test" {
		t.Fatalf("expected: `test`, but got : `%s`", cmd.Name())
	}
}

func TestCommandRunnable(t *testing.T) {
	if (&Command{}).Runnable() {
		t.Fatalf("command without Run func is must be NOT runnable")
	}
	if !(&Command{Run: func(cmd *Command, args []string) {}}).Runnable() {
		t.Fatalf("command without Run func is must be runnable")
	}
}
