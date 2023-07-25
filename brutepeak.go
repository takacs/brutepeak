package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const SIGMA = 2

type point struct {
	x         float64
	intensity float64
}

func BrutePeak(experimental_path, theoretical_path string) {
	experimental_points := normalize(parseFile(experimental_path, "\t"))
	theoretical_points_base := normalize(parseFile(theoretical_path, " "))
	minLSQ := float64(1000)
	var minPoints []point
	for i := 0; i <= 10000; i++ {
		theoretical_points := make([]point, len(theoretical_points_base))
		copy(theoretical_points, theoretical_points_base)
		theoretical_points = randomizeXPositions(theoretical_points)
		theoretical_points = createTheoretical(theoretical_points, experimental_points)
		lsq := leastSquare(experimental_points, theoretical_points)
		if lsq < minLSQ {
			minLSQ = lsq
			minPoints = theoretical_points
		}
		// fmt.Printf("Current iteration LSQ: %v | Minimum LSQ: %v\n", lsq, minLSQ)
	}
	for _, point := range minPoints {
		fmt.Printf("%v,%v\n", point.x, point.intensity)
	}
}

func parseFile(path, sep string) []point {
	points := make([]point, 0)
	pattern := regexp.MustCompile(`\s+`)

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		linestring := pattern.ReplaceAllString(scanner.Text(), " ")
		line := strings.Split(linestring, " ")
		x, _ := strconv.ParseFloat(line[0], 64)
		intensity, _ := strconv.ParseFloat(strings.ReplaceAll(line[1], "\\", ""), 64)
		points = append(points, point{x: x, intensity: intensity})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return points
}

func getGaussianValue(ep, tp point) float64 {
	multiplier := tp.intensity / (SIGMA * math.Sqrt(2*math.Pi))
	exponent := math.Exp(-math.Pow(ep.x-tp.x, 2) / (2 * math.Pow(SIGMA, 2)))
	return multiplier * exponent
}

func createTheoretical(theo_points, exp_points []point) []point {
	base := make([]point, len(exp_points))
	for _, tp := range theo_points {
		for i, ep := range exp_points {
			gaussVal := getGaussianValue(ep, tp)
			base[i] = point{x: ep.x, intensity: base[i].intensity + gaussVal}
		}

	}
	return base
}

func leastSquare(from, to []point) float64 {
	var lsf float64
	for i := range from {
		distance := math.Pow(from[i].intensity-to[i].intensity, 2)
		lsf = lsf + distance
	}
	return lsf
}

func normalize(spectra []point) []point {
	var max float64
	for i := range spectra {
		if spectra[i].intensity > max {
			max = spectra[i].intensity
		}
	}
	for i := range spectra {
		spectra[i].intensity = spectra[i].intensity / max
	}
	return spectra
}

func randomizeXPositions(spectra []point) []point {
	for i := range spectra {
		spectra[i].x = spectra[i].x + (-5 + 10*rand.Float64())
	}
	return spectra
}
