package main

type Game struct {
	ID       int
	MaxRed   int
	MaxGreen int
	MaxBlue  int
}

type ColourInput struct {
	R int
	G int
	B int
}

func (g *Game) PassResults(colours *ColourInput) {
	if g.MaxRed < colours.R {
		g.MaxRed = colours.R
	}

	if g.MaxGreen < colours.G {
		g.MaxGreen = colours.G
	}

	if g.MaxBlue < colours.B {
		g.MaxBlue = colours.B
	}
}

func (g Game) IsPossibleWithCubes(cubes *ColourInput) bool {
	if g.MaxRed > cubes.R {
		return false
	}

	if g.MaxGreen > cubes.G {
		return false
	}

	if g.MaxBlue > cubes.B {
		return false
	}

	return true
}

func (g Game) GetPowerSet() int {
	return g.MaxRed + g.MaxGreen + g.MaxBlue
}
