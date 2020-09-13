package remo

import (
	"context"
	"errors"

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
