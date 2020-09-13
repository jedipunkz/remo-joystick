package main

import (
	"fmt"
	"log"
	"os"

	"remo-joystick/pkg/remo"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
	"github.com/tenntenn/natureremo"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/joystick"
	"golang.org/x/net/context"
)

const (
	confFile = ".remo-joystick"
	platform = "xbox360"
)

// Remo is struct for communicating to Natureremo API
type Remo struct {
	client *natureremo.Client
}

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

// NewRemo is contstructor for Nature Remo API
func NewRemo(token string) *Remo {
	api := new(Remo)
	api.client = natureremo.NewClient(token)
	return api
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

	cli := NewRemo(token).client
	ctx := context.Background()

	joystickAdaptor := joystick.NewAdaptor()
	stick := joystick.NewDriver(joystickAdaptor, platform)

	button := NewButtons()

	work := func() {
		stick.On(joystick.APress, func(data interface{}) {
			if err := remo.SendSignal(cli, ctx,
				button.AButtonAppliance, button.AButtonSignal); err != nil {
				log.Fatal(err)
			}
		})
		stick.On(joystick.BPress, func(data interface{}) {
			if err := remo.SendSignal(cli, ctx,
				button.AButtonAppliance, button.AButtonSignal); err != nil {
				log.Fatal(err)
			}
		})
		stick.On(joystick.XPress, func(data interface{}) {
			if err := remo.SendSignal(cli, ctx,
				button.XButtonAppliance, button.XButtonSignal); err != nil {
				log.Fatal(err)
			}
		})
		stick.On(joystick.YPress, func(data interface{}) {
			if err := remo.SendSignal(cli, ctx,
				button.YButtonAppliance, button.YButtonSignal); err != nil {
				log.Fatal(err)
			}
		})
	}

	robot := gobot.NewRobot("joystickBot",
		[]gobot.Connection{joystickAdaptor},
		[]gobot.Device{stick},
		work,
	)
	robot.Start()
}
