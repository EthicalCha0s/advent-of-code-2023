package main

type Game struct {
	MaxRed   int32
	MaxGreen int32
	MaxBlue  int32
}

type ColourInput struct {
	R int32
	G int32
	B int32
}

func (g *Game) PassResults(colours *ColourInput) {
	if g.MaxRed > colours.R {
		g.MaxRed = colours.R
	}

	if g.MaxGreen > colours.G {
		g.MaxGreen = colours.G
	}

	if g.MaxBlue > colours.B {
		g.MaxBlue = colours.B
	}
}

func (g Game) GetMinimumCubes() int32 {
	sum := g.MaxRed + g.MaxGreen + g.MaxBlue
	return sum
}
