# remo-joystick

## Description

No need to talk to Alexa. Remo-Joystick make you enable to controll natureremo devices with Joystick such as Xbox360, Dualshock4 Controll Pads.

## Pre-Requirement(s)

- [SDL](https://www.libsdl.org/)

## Installation

```bash
$ go get github.com/jedipunkz/remo-joustick.git
```
## Setup $HOME/.remo-joystick.yaml file as follow

```yaml
token: <YOUR_NATUREREMO_TOKEN_HERE>
platform: xbox360 # or dualshock3, dualshock4
AButton:
  apl: <YOUR_APPLIANCE_NAME>
  sig: <YOUR_SIGNAL_NAME>
BButton:
  apl: <YOUR_APPLIANCE_NAME>
  sig: <YOUR_SIGNAL_NAME>
XButton:
  apl: <YOUR_APPLIANCE_NAME>
  sig: <YOUR_SIGNAL_NAME>
YButton:
  apl: <YOUR_APPLIANCE_NAME>
  sig: <YOUR_SIGNAL_NAME>
```

## Boot Process

```bash
$ remo-joystick
```

## TODO: Dockerize

when i can,i will.
