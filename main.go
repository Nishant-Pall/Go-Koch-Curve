package main

import (
	"math"

	"github.com/Nishant-Pall/go-koch-curve/binary"
	"github.com/fogleman/gg"
)

func main() {

	brep := binary.NewBinaryRep(65535)

	bitSumArr := brep.GetBitSumArray()

	const S = 20240
	dc := gg.NewContext(S, S)
	dc.SetRGB(0, 0, 0)
	dc.Clear()
	dc.Translate(S/2, S/2)
	dc.Scale(40, 40)
	dc.SetLineWidth(2)
	dc.SetRGBA(0, 1, 0, 0.4)

	var x, y float64
	x, y = 0, 200

	dx := 1 / (math.Sqrt(2) * 3)
	dy := 1 / (math.Sqrt(2) * 3)

	dc.MoveTo(x, y)
	rad := 60 * math.Pi / 180
	for _, bit := range bitSumArr {
		if bit == 0 {
			x += dx
			y += dy
			dc.LineTo(x, y)
		} else {
			oldDx := dx
			oldDy := dy
			dx = oldDx*math.Cos(rad) - oldDy*math.Sin(rad)
			dy = oldDx*math.Sin(rad) + oldDy*math.Cos(rad)
		}
	}
	dc.Stroke()

	dc.SavePNG("out.png")
}
