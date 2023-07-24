package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

const SIGMA = 5

func parseDatFile(path string) ([]float64, []float64, error) {
	xs := make([]float64, 0)
	ys := make([]float64, 0)
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		x, err := strconv.ParseFloat(line[0], 16)
		if err != nil {
			return nil, nil, err
		}
		xs = append(xs, x)
		y, err := strconv.ParseFloat(line[0], 16)
		if err != nil {
			return nil, nil, err
		}
		ys = append(ys, y)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return xs, ys, err
}

func getGaussianValue(x int, center float64) float64 {
	fx := float64(x)
	multiplier := 1 / (SIGMA * math.Sqrt(2*math.Pi))
	exponent := math.Exp(-math.Pow(fx-center, 2) / (2 * math.Pow(SIGMA, 2)))
	return multiplier * exponent
}

func createGuassian(xs, ys []float64, length int) []float64 {
	gaussFit := make([]float64, length)
	base := make([]float64, length)
	for _, x := range xs {
		for i := 0; i < length; i++ {
			gaussVal := getGaussianValue(i, x)
			base[i] = base[i] + gaussVal
		}

	}
	fmt.Println(base)
	return gaussFit
}

func Brutepeak(path string) {
	xs, ys, _ := parseDatFile(path)
	createGuassian(xs, ys, 3000)
}
