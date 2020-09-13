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
)

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
	aApl := viper.GetString("ButtonA.apl")
	aSig := viper.GetString("ButtonA.sig")
	bApl := viper.GetString("ButtonB.apl")
	bSig := viper.GetString("ButtonB.sig")

	cli := natureremo.NewClient(token)
	ctx := context.Background()

	joystickAdaptor := joystick.NewAdaptor()
	stick := joystick.NewDriver(joystickAdaptor, "xbox360")

	work := func() {
		stick.On(joystick.APress, func(data interface{}) {
			if err := remo.SendSignal(cli, ctx, aApl, aSig); err != nil {
				log.Fatal(err)
			}
		})
		stick.On(joystick.BPress, func(data interface{}) {
			if err := remo.SendSignal(cli, ctx, bApl, bSig); err != nil {
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
