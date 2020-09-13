package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/jedipunkz/remo-joystick/pkg/remo"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
	"github.com/tenntenn/natureremo"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/joystick"
	"golang.org/x/net/context"
)

func init() {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	viper.AddConfigPath(home)
	viper.SetConfigName(".remo-joystick")

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

	// fmt.Println(bApl)
	// fmt.Println(bSig)
	// os.Exit(0)

	cli := natureremo.NewClient(token)
	ctx := context.Background()

	joystickAdaptor := joystick.NewAdaptor()
	// stick := joystick.NewDriver(joystickAdaptor, joystick.Xbox360)
	stick := joystick.NewDriver(joystickAdaptor, "xbox360")

	work := func() {
		stick.On(joystick.APress, func(data interface{}) {
			if err := remoSend(cli, ctx, aApl, aSig); err != nil {
				log.Fatal(err)
			}
		})
		stick.On(joystick.BPress, func(data interface{}) {
			if err := remoSend(cli, ctx, bApl, bSig); err != nil {
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

func remoSend(cli *natureremo.Client, ctx context.Context, apl, sig string) error {
	a, err := remo.GetAppliance(ctx, cli, apl)
	if err != nil {
		log.Fatal(err)
		return err
	}

	s := remo.GetSignal(a.Signals, sig)
	if s == nil {
		var errNotFound = errors.New("Signal Not Found")
		log.Fatal("signal which you specified not found.")
		return errNotFound
	}

	if err := cli.SignalService.Send(ctx, s); err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
