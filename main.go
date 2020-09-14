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
)

type xbox360Buttons struct {
	a buttonActions
	b buttonActions
	x buttonActions
	y buttonActions
}

type dualshock4Buttons struct {
	circle   buttonActions
	triangle buttonActions
	square   buttonActions
	x        buttonActions
	up       buttonActions
	down     buttonActions
	right    buttonActions
	left     buttonActions
	r1       buttonActions
	r2       buttonActions
	l1       buttonActions
	l2       buttonActions
	option   buttonActions
	share    buttonActions
}

type buttonActions struct {
	press   aplSig
	release aplSig
}

type aplSig struct {
	appliance string
	signal    string
}

func newButtonsDualshock4() *dualshock4Buttons {
	buttons := new(dualshock4Buttons)
	buttons.circle.press.appliance = viper.GetString("dualshock4.CirclePress.apl")
	buttons.triangle.press.appliance = viper.GetString("dualshock4.TrianglePress.apl")
	buttons.square.press.appliance = viper.GetString("dualshock4.SquarePress.apl")
	buttons.x.press.appliance = viper.GetString("dualshock4.XPress.apl")
	buttons.up.press.appliance = viper.GetString("dualshock4.UpPress.apl")
	buttons.down.press.appliance = viper.GetString("dualshock4.DownPress.apl")
	buttons.right.press.appliance = viper.GetString("dualshock4.RightPress.apl")
	buttons.left.press.appliance = viper.GetString("dualshock4.LeftPress.apl")
	buttons.r1.press.appliance = viper.GetString("dualshock4.R1Press.apl")
	buttons.r2.press.appliance = viper.GetString("dualshock4.R1Press.apl")
	buttons.l1.press.appliance = viper.GetString("dualshock4.L1Press.apl")
	buttons.l2.press.appliance = viper.GetString("dualshock4.L1Press.apl")
	buttons.option.press.appliance = viper.GetString("dualshock4.OptionPress.apl")
	buttons.share.press.appliance = viper.GetString("dualshock4.SharePress.apl")
	buttons.circle.press.signal = viper.GetString("dualshock4.CirclePress.sig")
	buttons.triangle.press.signal = viper.GetString("dualshock4.TrianglePress.sig")
	buttons.square.press.signal = viper.GetString("dualshock4.SquarePress.sig")
	buttons.x.press.signal = viper.GetString("dualshock4.XPress.sig")
	buttons.up.press.signal = viper.GetString("dualshock4.UpPress.sig")
	buttons.down.press.signal = viper.GetString("dualshock4.DownPress.sig")
	buttons.right.press.signal = viper.GetString("dualshock4.RightPress.sig")
	buttons.left.press.signal = viper.GetString("dualshock4.LeftPress.sig")
	buttons.r1.press.signal = viper.GetString("dualshock4.R1Press.sig")
	buttons.r2.press.signal = viper.GetString("dualshock4.R1Press.sig")
	buttons.l1.press.signal = viper.GetString("dualshock4.L1Press.sig")
	buttons.l2.press.signal = viper.GetString("dualshock4.L1Press.sig")
	buttons.option.press.signal = viper.GetString("dualshock4.OptionPress.sig")
	buttons.share.press.signal = viper.GetString("dualshock4.SharePress.sig")
	return buttons
}

func newButtonsXbox360() *xbox360Buttons {
	buttons := new(xbox360Buttons)
	buttons.a.press.appliance = viper.GetString("xbox360.Apress.apl")
	buttons.a.press.signal = viper.GetString("xbox.Apress.sig")
	buttons.b.press.appliance = viper.GetString("xbox360.Bpress.apl")
	buttons.b.press.signal = viper.GetString("xbox.Xpress.sig")
	buttons.x.press.appliance = viper.GetString("xbox360.Xpress.apl")
	buttons.x.press.signal = viper.GetString("xbox.Cpress.sig")
	buttons.y.press.appliance = viper.GetString("xbox360.Ypress.apl")
	buttons.y.press.signal = viper.GetString("xbox.Ypress.sig")
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

	switch platform {
	case "dualshock4":
		button := newButtonsDualshock4()
		work := func() {
			j.Stick.On(joystick.CirclePress, func(data interface{}) {
				if err := r.SendSignalByAplSig(button.circle.press.appliance,
					button.circle.press.signal, ctx); err != nil {
					log.Fatal(err)
				}
			})

			j.Stick.On(joystick.TrianglePress, func(data interface{}) {
				if err := r.SendSignalByAplSig(button.circle.press.appliance,
					button.circle.press.signal, ctx); err != nil {
					log.Fatal(err)
				}
			})

			j.Stick.On(joystick.SquarePress, func(data interface{}) {
				if err := r.SendSignalByAplSig(button.circle.press.appliance,
					button.circle.press.signal, ctx); err != nil {
					log.Fatal(err)
				}
			})

			j.Stick.On(joystick.XPress, func(data interface{}) {
				if err := r.SendSignalByAplSig(button.circle.press.appliance,
					button.circle.press.signal, ctx); err != nil {
					log.Fatal(err)
				}
			})

			j.Stick.On(joystick.UpPress, func(data interface{}) {
				if err := r.SendSignalByAplSig(button.circle.press.appliance,
					button.circle.press.signal, ctx); err != nil {
					log.Fatal(err)
				}
			})

			j.Stick.On(joystick.DownPress, func(data interface{}) {
				if err := r.SendSignalByAplSig(button.circle.press.appliance,
					button.circle.press.signal, ctx); err != nil {
					log.Fatal(err)
				}
			})

			j.Stick.On(joystick.RightPress, func(data interface{}) {
				if err := r.SendSignalByAplSig(button.circle.press.appliance,
					button.circle.press.signal, ctx); err != nil {
					log.Fatal(err)
				}
			})

			j.Stick.On(joystick.LeftPress, func(data interface{}) {
				if err := r.SendSignalByAplSig(button.circle.press.appliance,
					button.circle.press.signal, ctx); err != nil {
					log.Fatal(err)
				}
			})

			j.Stick.On(joystick.R1Press, func(data interface{}) {
				if err := r.SendSignalByAplSig(button.circle.press.appliance,
					button.circle.press.signal, ctx); err != nil {
					log.Fatal(err)
				}
			})

			j.Stick.On(joystick.R2Press, func(data interface{}) {
				if err := r.SendSignalByAplSig(button.circle.press.appliance,
					button.circle.press.signal, ctx); err != nil {
					log.Fatal(err)
				}
			})

			j.Stick.On(joystick.L1Press, func(data interface{}) {
				if err := r.SendSignalByAplSig(button.circle.press.appliance,
					button.circle.press.signal, ctx); err != nil {
					log.Fatal(err)
				}
			})

			j.Stick.On(joystick.L2Press, func(data interface{}) {
				if err := r.SendSignalByAplSig(button.circle.press.appliance,
					button.circle.press.signal, ctx); err != nil {
					log.Fatal(err)
				}
			})

			j.Stick.On(joystick.OptionsPress, func(data interface{}) {
				if err := r.SendSignalByAplSig(button.circle.press.appliance,
					button.circle.press.signal, ctx); err != nil {
					log.Fatal(err)
				}
			})

			j.Stick.On(joystick.SharePress, func(data interface{}) {
				if err := r.SendSignalByAplSig(button.circle.press.appliance,
					button.circle.press.signal, ctx); err != nil {
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

	case "xbox360":
		button := newButtonsXbox360()
		work := func() {
			j.Stick.On(joystick.APress, func(data interface{}) {
				if err := r.SendSignalByAplSig(button.a.press.appliance,
					button.a.press.signal, ctx); err != nil {
					log.Fatal(err)
				}
			})

			j.Stick.On(joystick.BPress, func(data interface{}) {
				if err := r.SendSignalByAplSig(button.b.press.appliance,
					button.b.press.signal, ctx); err != nil {
					log.Fatal(err)
				}
			})

			j.Stick.On(joystick.XPress, func(data interface{}) {
				if err := r.SendSignalByAplSig(button.x.press.appliance,
					button.x.press.signal, ctx); err != nil {
					log.Fatal(err)
				}
			})

			j.Stick.On(joystick.YPress, func(data interface{}) {
				if err := r.SendSignalByAplSig(button.y.press.appliance,
					button.y.press.signal, ctx); err != nil {
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
}
