package main

import (
	"log"
	"math"
	"strconv"
	"strings"

	"math/rand"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

type polynomial struct {
	coef   []int
	degree int
}

//http://fooplot.com/?lang=es#W3sidHlwZSI6MCwiZXEiOiJ4XjIrMngrOC0zeF4yIiwiY29sb3IiOiIjMDAwMDAwIn0seyJ0eXBlIjoxMDAwLCJ3aW5kb3ciOlsiLTkuMDkzNzQ5OTk5OTk5OTk2IiwiNy4xNTYyNDk5OTk5OTk5OTY0IiwiLTEuMjE4NzUiLCI4Ljc4MTI0OTk5OTk5OTk5NiJdfV0-
func (p *polynomial) evaluate(x float64) float64 {
	var r float64
	for i := 0; i <= p.degree; i++ {
		r += math.Pow(float64(x), float64(p.degree-i)) * float64(p.coef[i])
	}
	return r
}

func (p *polynomial) processPolynomialPart(coef int, degree int) {
	if p.degree < degree {
		log.Printf("Coef: %v", p.coef)

		var newCoef = make([]int, degree+1)

		newCoef[0] = coef

		for i := 1; i < len(newCoef); i++ {
			//p.degree=2 degree=10
			if i < (degree-p.degree) || p.degree == 0 {
				newCoef[i] = 0
			} else {
				newCoef[i] = p.coef[i-(degree-p.degree)]
			}
		}

		p.degree = degree
		p.coef = newCoef
	} else {
		p.coef[len(p.coef)-degree-1] = p.coef[len(p.coef)-degree-1] + coef
	}
}

func parsePolynomialPart(part string, isNegative bool) (coef int, degree int) { //  3x^2 | x^2 | 2x | x  | 3, true | false
	//coef, degree int
	if isNegative {
		coef = -1
	} else {
		coef = 1
	}
	if !strings.Contains(part, "x") {
		degree = 0
		i, _ := strconv.Atoi(part)
		coef = coef * i
	} else if !strings.Contains(part, "^") {
		degree = 1
		part = strings.ReplaceAll(part, "x", "")
		if len(part) > 0 {
			i, _ := strconv.Atoi(part)
			coef = coef * i
		}
	} else {
		parts := strings.Split(part, "^")
		i, _ := strconv.Atoi(parts[1])
		degree = i
		part = strings.ReplaceAll(parts[0], "x", "")
		if len(part) > 0 {
			i, _ := strconv.Atoi(part)
			coef = coef * i
		}
	}
	return coef, degree
}

func main() {

	// polynomialString := os.Args[1]

	// log.Println("Polynomio ", polynomialString)

	// var coef = make([]int, 0)
	// p := polynomial{degree: 0, coef: coef}

	// parts := strings.Split(polynomialString, "+")
	// for _, part := range parts {
	// 	if strings.Contains(part, "-") {
	// 		parts2 := strings.Split(part, "-")
	// 		for i, part2 := range parts2 {
	// 			if i == 0 {
	// 				p.processPolynomialPart(parsePolynomialPart(part2, false))
	// 			} else {
	// 				p.processPolynomialPart(parsePolynomialPart(part2, true))
	// 			}
	// 		}
	// 	} else {
	// 		p.processPolynomialPart(parsePolynomialPart(part, false))
	// 	}
	// }
	//
	// log.Println("Degree: %d", p.degree)
	// log.Println("Coef: %v", p.coef)
	//log.Println("Con 0: ", p.evaluate(float64(0)))
	//log.Println("Con 1: ", p.evaluate(float64(1)))
	//log.Println("Con 2: ", p.evaluate(float64(2)))

	// for x := -100; x < 100; x++ {
	// 	y := p.evaluate(float64(x))
	// 	log.Printf("x:%d y:%f", x, y)
	// }

	rand.Seed(int64(0))

	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	p.Title.Text = "Plotutil example"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	err = plotutil.AddLinePoints(p,
		"First", randomPoints(15),
		"Second", randomPoints(15),
		"Third", randomPoints(15))
	if err != nil {
		panic(err)
	}

	// Save the plot to a PNG file.
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "points.png"); err != nil {
		panic(err)
	}
}

// randomPoints returns some random x, y points.
func randomPoints(n int) plotter.XYs {
	pts := make(plotter.XYs, n)
	for i := range pts {
		if i == 0 {
			pts[i].X = rand.Float64()
		} else {
			pts[i].X = pts[i-1].X + rand.Float64()
		}
		pts[i].Y = pts[i].X + 10*rand.Float64()
	}
	return pts
}
