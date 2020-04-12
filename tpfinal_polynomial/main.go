package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/guptarohit/asciigraph"
)

type polynomial struct {
	coef   []int
	degree int
}

func (p *polynomial) roots() []float64 {
	var roots []float64
	var coefs []float64
	for i := 0; i < len(p.coef); i++ {
		coefs = append(coefs, float64(p.coef[i]))
	}
	return p.roots2(coefs, roots)
}

func (p *polynomial) roots2(coefs []float64, roots []float64) []float64 {
	if len(p.coef) > 3 {
		k := p.findK(coefs)
		if k != nil {
			roots = append(roots, *k)
			coefs = p.dividePolynomialByXMinusK(coefs, k)
			return p.roots2(coefs, roots)
		}
	} else if len(p.coef) == 3 {
		roots = append(roots, p.getQuadraticRoots(coefs[0], coefs[1], coefs[2])...)
	} else if len(coefs) == 2 {
		roots = append(roots, p.getQuadraticRoots(0, coefs[0], coefs[1])...)
	}
	sort.Float64s(roots)
	for i := 0; i < len(roots); i++ {
		if roots[i] == 0 {
			roots[i] = 0
		}
		for j := i + 1; j < len(roots); j++ {
			if roots[i] == roots[j] {
				roots = append(roots[:j], roots[j+1:]...)
				j--
			}
		}
	}
	return roots
}

func (p *polynomial) findK(coefs []float64) *float64 {
	if len(coefs) > 2 {
		constant := coefs[len(coefs)-1]
		if constant == 0 {
			a := float64(0)
			return &a
		}
		var trialKValues []float64

		for i := 1; i <= int(math.Abs(constant)); i++ {
			if int(math.Abs(constant))%i == 0 {
				trialKValues = append(trialKValues, float64(i))
				trialKValues = append(trialKValues, float64(i*-1))
			}
		}

		for _, k := range trialKValues {
			sumOfTerms := float64(0)
			for i := 0; i < len(coefs); i++ {
				sumOfTerms += coefs[i] * math.Pow(k, float64(len(coefs)-1-i))
			}
			if sumOfTerms == 0 {
				return &k
			}
		}
	}
	return nil
}

func (p *polynomial) dividePolynomialByXMinusK(coefs []float64, k *float64) []float64 {
	if k != nil && coefs != nil && len(coefs) > 2 {
		var newCoefficients []float64
		newCoefficients = append(newCoefficients, coefs[0])

		for i := 1; i < len(coefs)-1; i++ {
			newCoefficients = append(newCoefficients, newCoefficients[i-1]**k+coefs[i])
		}

		lastIndex := len(coefs) - 1

		if newCoefficients[lastIndex-1]**k+coefs[lastIndex] != 0 {
			return nil
		}

		return newCoefficients
	}
	return nil
}

func (p *polynomial) getQuadraticRoots(a float64, b float64, c float64) []float64 {
	var roots []float64

	if a != 0 {
		nRoots := 0
		discriminant := math.Pow(b, 2) - 4*a*c
		if discriminant > 0 {
			nRoots = 2
		} else if discriminant == 0 {
			nRoots = 1
		}

		if nRoots == 1 {
			roots = append(roots, b*-1/(2*a))
		} else if nRoots == 2 {
			roots = append(roots, ((b*-1 + math.Sqrt(discriminant)) / (2 * a)))
			roots = append(roots, ((b*-1 - math.Sqrt(discriminant)) / (2 * a)))
		}
	} else {
		roots = append(roots, (c*-1)/b)
	}
	return roots
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

	if len(os.Args) != 2 {
		log.Fatalf("Invalid parameters. Enter a polynomial function. Ex: x^2+2x-8")
	}
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

	r := p.roots()
	log.Printf("Real Roots: %v", r)

	var ys []float64
	var start, end float64

	if len(r) == 0 {
		start = -5
		end = 5
	} else {
		start = r[0] - 1
		end = r[len(r)-1] + 1
	}

	for x := start; x <= end; x++ {
		y := p.evaluate(float64(x))
		ys = append(ys, y)
	}

	data := ys
	graph := asciigraph.Plot(data)

	fmt.Println(graph)

}
