package main

import (
	"os"

	svg "github.com/ajstarks/svgo"
)

func main() {
	width := 2000
	height := 850
	canvas := svg.New(os.Stdout)
	canvas.Start(width, height)
	canvas.Gstyle("opacity:1.0;stroke-width:5")

	// border
	canvas.Roundrect(0, 0, 2000, 850, 280, 280, "stroke:black;fill:white")

	// referee crease
	canvas.Path("M1000,850 L900,850 A100,100 1 0,1 1100,850 z", "stroke:red;fill:white")

	// center faceoff w/center line
	canvas.Circle(1000, 425, 150, "stroke:blue;fill:white")
	canvas.Line(1000, 0, 1000, 850, "stroke:red;stroke-dasharray:5")
	canvas.Circle(1000, 425, 10, "stroke:blue;fill:blue")

	drawNeutralZoneDots(canvas, 800, 205)
	drawNeutralZoneDots(canvas, 800, 645)
	drawNeutralZoneDots(canvas, 1200, 205)
	drawNeutralZoneDots(canvas, 1200, 645)

	// attacking/defensive faceoff circles
	drawAttackingZoneFaceoff(canvas, 310, 205)
	drawAttackingZoneFaceoff(canvas, 310, 645)
	drawAttackingZoneFaceoff(canvas, 1690, 205)
	drawAttackingZoneFaceoff(canvas, 1690, 645)

	// goal lines
	drawGoalLine(canvas, 110)
	drawGoalLine(canvas, 1890)

	// goal crease
	canvas.Path("M110,425 L110,365 A60,60 1 0,1 110,485 z", "stroke:red;fill:blue")
	canvas.Path("M1890,425 L1890,485 A60,60 1 0,1 1890,365 z", "stroke:red;fill:blue")

	// defensive/attacking zone lines
	canvas.Line(750, 0, 750, 850, "stroke:blue")
	canvas.Line(1250, 0, 1250, 850, "stroke:blue")
	canvas.Gend()
	canvas.End()
}

func drawAttackingZoneFaceoff(canvas *svg.SVG, x, y int) {
	canvas.Gstyle("stroke:red")
	// circles
	canvas.Circle(x, y, 150, "fill:white")
	canvas.Circle(x, y, 10, "fill:red")
	// |
	canvas.Line(x-20, y-10, x-20, y-40)
	canvas.Line(x+20, y-10, x+20, y-40)
	canvas.Line(x-20, y+10, x-20, y+40)
	canvas.Line(x+20, y+10, x+20, y+40)
	// -
	canvas.Line(x-20, y-10, x-60, y-10)
	canvas.Line(x+20, y-10, x+60, y-10)
	canvas.Line(x-20, y+10, x-60, y+10)
	canvas.Line(x+20, y+10, x+60, y+10)
	// hash marks
	canvas.Line(x-20, y+150, x-20, y+170)
	canvas.Line(x+20, y+150, x+20, y+170)
	canvas.Line(x-20, y-150, x-20, y-170)
	canvas.Line(x+20, y-150, x+20, y-170)
	canvas.Gend()
}

func drawNeutralZoneDots(canvas *svg.SVG, x, y int) {
	canvas.Circle(x, y, 10, "stroke:red;fill:red")
}

func drawGoalLine(canvas *svg.SVG, x int) {
	canvas.Line(x, 57, x, 793, "stroke:red")
}
