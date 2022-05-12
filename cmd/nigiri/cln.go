package main

import (
	"errors"
	"os"
	"os/exec"

	"github.com/urfave/cli/v2"
)

var cln = cli.Command{
	Name:   "cln",
	Usage:  "invoke Core Lightning command line",
	Action: clnAction,
	Flags: []cli.Flag{
		&noTTYFlag,
	},
}

func clnAction(ctx *cli.Context) error {

	if isRunning, _ := nigiriState.GetBool("running"); !isRunning {
		return errors.New("nigiri is not running")
	}

	network, err := nigiriState.GetString("network")
	if err != nil {
		return err
	}

	isCi, err := nigiriState.GetBool("ci")
	if err != nil {
		return err
	}

	isNoTTY := ctx.Bool("no-tty")
	ttyOption := "-it"
	if isCi || isNoTTY {
		ttyOption = "-i"
	}
	rpcArgs := []string{"exec", ttyOption, "cln", "lightning-cli", "--network=" + network}
	cmdArgs := append(rpcArgs, ctx.Args().Slice()...)
	bashCmd := exec.Command("docker", cmdArgs...)
	bashCmd.Stdin = os.Stdin
	bashCmd.Stdout = os.Stdout
	bashCmd.Stderr = os.Stderr

	if err := bashCmd.Run(); err != nil {
		return err
	}

	return nil
}
