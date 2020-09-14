package mygobot

import (
	"gobot.io/x/gobot/platforms/joystick"
)

// Joystick is struct for Gobot Platform as Joystick
type Joystick struct {
	JoystickAdaptor *joystick.Adaptor
	Stick           *joystick.Driver
	Platform        string
}

// NewGobot is constructor for Gobot
func NewGobot(platform string) *Joystick {
	gobot := new(Joystick)
	gobot.JoystickAdaptor = joystick.NewAdaptor()
	gobot.Stick = joystick.NewDriver(gobot.JoystickAdaptor, platform)
	return gobot
}
