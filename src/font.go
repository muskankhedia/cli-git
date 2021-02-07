package main

import (
	"fmt"
	"github.com/mbndr/figlet4go"
)

func projectText() {

	ascii := figlet4go.NewAsciiRender()

	options := figlet4go.NewRenderOptions()
	options.FontName = "block"
	options.FontColor = []figlet4go.Color{
		figlet4go.ColorGreen,
		figlet4go.ColorYellow,
		figlet4go.ColorCyan,
	}

	renderStr, _ := ascii.RenderOpts("CLI Git", options)
	fmt.Print(renderStr)

}
