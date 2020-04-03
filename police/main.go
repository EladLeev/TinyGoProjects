package main

import (
	"machine"
	"time"
)

func main() {
	// The blue and the red leds are connected to PIN 12,13
	led, led2 := machine.Pin(13), machine.Pin(12)
	led.Configure(machine.PinConfig{1})
	led2.Configure(machine.PinConfig{1})

	// Led switch
	ledSwitch := true
	ledSwitch2 := false

	for {
		led.Set(ledSwitch)
		led2.Set(ledSwitch2)
		ledSwitch, ledSwitch2 = !ledSwitch, !ledSwitch2
		time.Sleep(250 * time.Millisecond)
	}
}
