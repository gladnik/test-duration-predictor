package main

import (
	"strings"
	"strconv"
	"os"
	"bufio"
	"fmt"
)

var (
	reader = bufio.NewReader(os.Stdin)
)

func ReadFromStdin() string {
	result, _ := reader.ReadString('\n')
	witl := result[:len(result)-1]
	return witl
}

func StringArrAsFloats64(strArr []string) []float64 {
	var floats []float64
	for _, i := range strArr {
		j, err := strconv.ParseFloat(i, 64)
		if err != nil {
			panic(err)
		}
		floats = append(floats, j)
	}
	return floats
}

func GetTimings() []float64 {
	var input = ReadFromStdin()
	return StringArrAsFloats64(strings.Fields(input))
}

func PredictNextByWindow(timings []float64, windowSize int) float64 {
	if len(timings) < windowSize {
		return timings[len(timings)-1]
	}
	var total float64 = 0
	for _, value := range timings[len(timings)-windowSize:] {
		total += value
	}
	return total / float64(windowSize)
}

func main() {
	var timings = GetTimings()
	fmt.Println(PredictNextByWindow(timings, 5))
}
