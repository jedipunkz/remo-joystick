package remo

import (
	"context"
	"errors"
	"log"

	"github.com/tenntenn/natureremo"
)

func GetAppliance(ctx context.Context, cli *natureremo.Client, name string) (*natureremo.Appliance, error) {
	as, err := cli.ApplianceService.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	for _, a := range as {
		if a.Nickname == name {
			return a, nil
		}
	}

	return nil, errors.New("appliance not found")
}

func GetSignal(ss []*natureremo.Signal, name string) *natureremo.Signal {
	for _, s := range ss {
		if s.Name == name {
			return s
		}
	}
	return nil
}

func SendSignal(cli *natureremo.Client, ctx context.Context, apl, sig string) error {
	a, err := GetAppliance(ctx, cli, apl)
	if err != nil {
		log.Fatal(err)
		return err
	}

	s := GetSignal(a.Signals, sig)
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
