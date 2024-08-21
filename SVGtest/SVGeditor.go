package main

import (
	"fmt"

	s "github.com/timskillman/oksvg"
)

func main() {
	folder := "../testSVGs/"
	file := "beach"

	icon, errSvg := s.ReadIcon(folder + file + ".svg")
	if errSvg == nil {
		icon.SaveAsPngResized("beach.png", 2000, -1)
		icon.ViewBox.W *= icon.Transform.A
		icon.ViewBox.H *= icon.Transform.A
		icon.RotateOnPivot(1.57, icon.ViewBox.W/2, icon.ViewBox.H/2)
		err := icon.SaveAsPng("beachRot.png")
		if err != nil {
			fmt.Printf("error: %s", err)
		}
	} else {
		fmt.Println("error: %s", errSvg)
	}

}
