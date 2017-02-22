package controls

import (
	"fmt"
	"time"

	"github.com/kidoman/embd"
)

// NewStartLED ... gets a new start led pin
func NewStartLED() (embd.DigitalPin, error) {

	pin, err := embd.NewDigitalPin(startLEDPinName)
	if err != nil {
		return nil, fmt.Errorf("Error encountered while trying to create start LED pin: %v", err)
	}
	pin.SetDirection(embd.Out)

	return pin, nil
}

// BlinkStartLED ...
func BlinkStartLED(pin embd.DigitalPin, killChan chan int) error {
	defer startLEDOff(pin)

	val, err := pin.Read()
	if err != nil {
		return err
	}

	togglePin(val, pin)

	t := time.NewTicker(500 * time.Millisecond)

	for {
		select {
		case <-t.C:
			val, err := pin.Read()
			if err != nil {
				return err
			}

			togglePin(val, pin)
		case <-killChan:
			return nil
		}
	}
}

func togglePin(state int, pin embd.DigitalPin) {
	if state == embd.Low {
		pin.Write(embd.High)
		return
	}

	pin.Write(embd.Low)
}

func startLEDOff(pin embd.DigitalPin) error {
	return pin.Write(embd.Low)
}

func startLEDOn(pin embd.DigitalPin) error {
	return pin.Write(embd.High)
}