package pt3

import (
	"fmt"
	"os"
)

// Tune tunes the recording
func Tune(r *Record, device string) error {
	c := r.Channel

	// TODO driver frequency
	freqNo := c.setFreq
	slot := c.addFreq
	_ = freqNo
	_ = slot

	if device == "" {
		var devices []string
		switch c.channelType {
		case channelTypeSatellite:
			devices = bsDevices
		case channelTypeGround:
			devices = isdbtDevices
		}

		for _, dev := range devices {
			f, err := os.OpenFile(dev, os.O_RDONLY, os.ModePerm)
			if err == nil {
				r.Tuner = f
				break
			}
		}

		if r.Tuner == nil {
			return fmt.Errorf("Cannot tune to the specified channel")
		}
	} else {
		f, err := os.OpenFile(device, os.O_RDONLY, os.ModePerm)
		if err != nil {
			return fmt.Errorf("Cannot tune to the specified channel")
		}
		r.Tuner = f
	}

	// TODO power on LNB

	// TODO tune to specified channel

	return nil
}
