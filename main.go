package main

import (
	"fmt"
	"log"
	"os"

	mygobot "remo-joystick/pkg/gobot"
	myremo "remo-joystick/pkg/remo"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/joystick"
	"golang.org/x/net/context"
)

const (
	confFile = ".remo-joystick"
	platform = "xbox360"
)

// Buttons is struct controller pad buttons
type Buttons struct {
	AButtonAppliance string
	AButtonSignal    string
	BButtonAppliance string
	BButtonSignal    string
	XButtonAppliance string
	XButtonSignal    string
	YButtonAppliance string
	YButtonSignal    string
}

// NewButtons is contstructor for Controller Pad's Buttons
func NewButtons() *Buttons {
	buttons := new(Buttons)
	buttons.AButtonAppliance = viper.GetString("AButton.apl")
	buttons.AButtonSignal = viper.GetString("AButton.sig")
	buttons.BButtonAppliance = viper.GetString("BButton.apl")
	buttons.BButtonSignal = viper.GetString("BButton.sig")
	buttons.XButtonAppliance = viper.GetString("XButton.apl")
	buttons.XButtonSignal = viper.GetString("XButton.sig")
	buttons.YButtonAppliance = viper.GetString("YButton.apl")
	buttons.YButtonSignal = viper.GetString("YButton.sig")
	return buttons
}

func init() {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	viper.AddConfigPath(home)
	viper.SetConfigName(confFile)

	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file: ", viper.ConfigFileUsed())
	}
}

func main() {
	token := viper.GetString("token")
	platform := viper.GetString("platform")

	r := myremo.NewRemo(token)
	ctx := context.Background()

	j := mygobot.NewGobot(platform)

	button := NewButtons()

	work := func() {
		j.Stick.On(joystick.APress, func(data interface{}) {
			if err := r.GetAppliance(ctx, button.AButtonAppliance); err != nil {
				log.Fatal(err)
			}
			if err := r.GetSignal(r.Appliance.Signals, button.AButtonSignal); err != nil {
				log.Fatal(err)
			}
			if err := r.SendSignal(ctx); err != nil {
				fmt.Println(err)
				log.Fatal(err)
			}
		})
		j.Stick.On(joystick.BPress, func(data interface{}) {
			if err := r.GetAppliance(ctx, button.BButtonAppliance); err != nil {
				log.Fatal(err)
			}
			if err := r.GetSignal(r.Appliance.Signals, button.BButtonSignal); err != nil {
				log.Fatal(err)
			}
			if err := r.SendSignal(ctx); err != nil {
				fmt.Println(err)
				log.Fatal(err)
			}
		})
		j.Stick.On(joystick.XPress, func(data interface{}) {
			if err := r.GetAppliance(ctx, button.XButtonAppliance); err != nil {
				log.Fatal(err)
			}
			if err := r.GetSignal(r.Appliance.Signals, button.XButtonSignal); err != nil {
				log.Fatal(err)
			}
			if err := r.SendSignal(ctx); err != nil {
				fmt.Println(err)
				log.Fatal(err)
			}
		})
		j.Stick.On(joystick.YPress, func(data interface{}) {
			if err := r.GetAppliance(ctx, button.YButtonAppliance); err != nil {
				log.Fatal(err)
			}
			if err := r.GetSignal(r.Appliance.Signals, button.YButtonSignal); err != nil {
				log.Fatal(err)
			}
			if err := r.SendSignal(ctx); err != nil {
				fmt.Println(err)
				log.Fatal(err)
			}
		})
	}

	robot := gobot.NewRobot("joystickBot",
		[]gobot.Connection{j.JoystickAdaptor},
		[]gobot.Device{j.Stick},
		work,
	)
	robot.Start()
}
