package main

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var statusCmd = cli.Command{
	Name:   "status",
	Action: statusAction,
}

func statusAction(ctx *cli.Context) error {

	colorRed := "\033[31m"
	colorGreen := "\033[32m"
	colorReset := "\033[0m"

	if isRunning, _ := nigiriState.GetBool("running"); !isRunning {
		fmt.Println("nigiri is" + string(colorRed) + " not running")
		return nil
	}

	isCIRunning, _ := nigiriState.GetBool("ci")
	ciColor := colorRed
	if isCIRunning {
		ciColor = colorGreen
	}

	isLiquidRunning, _ := nigiriState.GetBool("liquid")
	liquidColor := colorRed
	if isLiquidRunning {
		liquidColor = colorGreen
	}

	isLNRunning, _ := nigiriState.GetBool("ln")
	lnColor := colorRed
	if isLNRunning {
		lnColor = colorGreen
	}

	fmt.Println("nigiri is" + string(colorGreen) + " running")
	fmt.Println(string(colorReset), "CI"+string(ciColor), isCIRunning)
	fmt.Println(string(colorReset), "Liquid"+string(liquidColor), isLiquidRunning)
	fmt.Println(string(colorReset), "LN"+string(lnColor), isLNRunning)

	return nil
}
