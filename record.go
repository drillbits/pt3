package pt3

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

// Record represents the state of an recording.
type Record struct {
	Tuner       io.Reader
	Destination io.Writer
	LNB         int
	MsqID       int
	StartTime   time.Time
	RecSec      time.Duration
	Channel     *isdbtChannel
	Indefinite  bool
}

// NewRecord returns a new recording state.
func NewRecord(args []string) (*Record, error) {
	r := &Record{}

	// channel
	if len(args) > 0 {
		c, ok := isdbtConvertTable[args[0]]
		if !ok {
			return nil, fmt.Errorf("invalid channel: %s", args[0])
		}
		r.Channel = c
	}

	// rectime
	if len(args) > 1 {
		err := r.SetTime(args[1])
		if err != nil {
			return nil, err
		}
	}

	// destfile
	if len(args) > 2 {
		err := r.SetDestination(args[2])
		if err != nil {
			return nil, err
		}
	}

	return r, nil
}

// SetTime sets recording duration.
func (r *Record) SetTime(s string) error {
	if s == "-" {
		r.Indefinite = true
		return nil
	}

	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return fmt.Errorf("rectime must be a number or `-`")
	}

	r.RecSec = time.Duration(i) * time.Second
	if r.RecSec < 0 {
		r.Indefinite = true
	}

	return nil
}

// SetDestination sets the recording destination.
func (r *Record) SetDestination(s string) error {
	if s == "-" {
		r.Destination = os.Stdout
		return nil
	}

	dst, err := os.OpenFile(s, os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return err
	}
	r.Destination = dst

	return nil
}

// IsAvailable reports whether the recording is available.
func (r *Record) IsAvailable() bool {
	return r.Channel != nil && (r.RecSec > 0 || r.Indefinite) && r.Destination != nil
}
