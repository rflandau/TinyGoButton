/*
This package represents a single button that communicates via the boolean state of the given pin.
*/
package button

import (
	"machine"
	"time"
)

type Device struct {
	pin            machine.Pin
	statePrior     bool          // was the button being held the last time we checked?
	stateLastCheck time.Time     // the last time we checked if the button was being held
	stateDuration  time.Duration // the total duration of the current, contiguous state of the button (held or not held)
}

// Sets pins for A and B as inputs
func New(pin machine.Pin) *Device {
	btns := &Device{
		pin:            pin,
		stateLastCheck: time.Now(),
		stateDuration:  0}
	// buttons
	btns.pin.Configure(machine.PinConfig{Mode: machine.PinInput})

	return btns
}

func (dev *Device) Pressed() bool {
	return dev.pin.Get()
}

// Updates the statistics on A's held duration by checking the current state of A.
// The more often this subroutine is called, the more accurate the time statistics are.
// Returns whether A is being held and how long A has been held for.
// This subroutine can be used in place of .Get().
// It is a more costly function, but, as stated above, is more accurate the more rapidly it is called.
func (dev *Device) Held() (bool, time.Duration) {
	cur := dev.pin.Get()
	if cur == dev.statePrior { // if the button is in the same state as it previously was, accumulate
		dev.stateDuration += time.Since(dev.stateLastCheck)
	} else { // if the button's state differs, clear the accumulation
		dev.stateDuration = 0
	}
	dev.statePrior = cur

	// update the last time we checked A being held
	dev.stateLastCheck = time.Now()
	return dev.statePrior, dev.stateDuration
}
