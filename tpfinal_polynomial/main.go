package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	"math/rand"

	"github.com/guptarohit/asciigraph"
	"gonum.org/v1/plot/plotter"
)

type polynomial struct {
	coef   []int
	degree int
}

func (p *polynomial) evaluate(x float64) float64 {
	var r float64
	for i := 0; i <= p.degree; i++ {
		r += math.Pow(float64(x), float64(p.degree-i)) * float64(p.coef[i])
	}
	return r
}

func (p *polynomial) processPolynomialPart(coef int, degree int) {
	if p.degree < degree {
		var newCoef = make([]int, degree+1)

		newCoef[0] = coef

		for i := 1; i < len(newCoef); i++ {
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

func parsePolynomialPart(part string, isNegative bool) (coef int, degree int) {
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

	polynomialString := os.Args[1]

	log.Println("Polynomio ", polynomialString)

	var coef = make([]int, 0)
	p := polynomial{degree: 0, coef: coef}

	parts := strings.Split(polynomialString, "+")
	for _, part := range parts {
		if strings.Contains(part, "-") {
			parts2 := strings.Split(part, "-")
			for i, part2 := range parts2 {
				if i == 0 {
					p.processPolynomialPart(parsePolynomialPart(part2, false))
				} else {
					p.processPolynomialPart(parsePolynomialPart(part2, true))
				}
			}
		} else {
			p.processPolynomialPart(parsePolynomialPart(part, false))
		}
	}

	log.Printf("Degree: %d", p.degree)
	log.Printf("Coef: %v", p.coef)

	var ys []float64

	for x := -5; x < 5; x++ {
		y := p.evaluate(float64(x))
		ys = append(ys, y)
	}

	data := ys
	graph := asciigraph.Plot(data)

	fmt.Println(graph)
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
