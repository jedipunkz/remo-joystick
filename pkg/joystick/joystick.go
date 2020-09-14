package myjoystick

import (
	"gobot.io/x/gobot/platforms/joystick"
)

// Joystick is struct for Gobot Platform as Joystick
type Joystick struct {
	JoystickAdaptor *joystick.Adaptor
	Stick           *joystick.Driver
	Platform        string
}

// NewJoystick is constructor for Gobot Platform Joystick
func NewJoystick(platform string) *Joystick {
	js := new(Joystick)
	js.JoystickAdaptor = joystick.NewAdaptor()
	js.Stick = joystick.NewDriver(js.JoystickAdaptor, platform)
	return js
}
