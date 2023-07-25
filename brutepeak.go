package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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
	experimental_points := parseFile(experimental_path, "\t")
	theoretical_points := parseFile(theoretical_path, " ")
	createTheoretical(theoretical_points, experimental_points)
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

func createTheoretical(theo_points, exp_points []point) []float64 {
	base := make([]float64, len(exp_points))
	for _, tp := range theo_points {
		for i, ep := range exp_points {
			gaussVal := getGaussianValue(ep, tp)
			base[i] = base[i] + gaussVal
		}

	}
	fmt.Println(base)
	return base
}
