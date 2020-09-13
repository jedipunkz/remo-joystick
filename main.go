package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jedipunkz/remo-joystick/pkg/remo"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
	"github.com/tenntenn/natureremo"
	"golang.org/x/net/context"
)

func init() {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	viper.AddConfigPath(home)
	viper.SetConfigName(".remo-controller")
}

func main() {
	token := viper.GetString("token")
	aApl := viper.GetString("aApl")
	aSig := viper.GetString("aSig")

	cli := natureremo.NewClient(token)
	ctx := context.Background()

	a, err := remo.GetAppliance(ctx, cli, aApl)
	if err != nil {
		log.Fatal(err)
	}

	s := remo.GetSignal(a.Signals, aSig)
	if s == nil {
		log.Fatal("signal which you specified not found.")
	}

	if err := cli.SignalService.Send(ctx, s); err != nil {
		log.Fatal(err)
	}
}
